package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

// Level 用來定義 Logger 的分級
type Level int8

// Fields ...
type Fields map[string]interface{}

// 根據 Level 建立 enum
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

// Logger 會紀錄 logger 的內容
type Logger struct {
	// NOTICE: 為什麼 newLogger 要是 pointer
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

// NewLogger 用來產生 Logger 物件
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// WithFields 可以用來為 Logger 物件添加說明欄位
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()

	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

// WithContext 可以用 Logger 添加 context
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx

	return ll
}

// WithCaller 會回傳呼叫該方法的檔案、行數、program counter
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()

	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return ll
}

// WithCallersFrames 會回傳完整的堆疊資訊
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}

	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}

	ll := l.clone()
	ll.callers = callers
	return ll
}

// JSONFormat 會將 Logger 中原本的 fields 和額外添加的屬性（level, message, time, caller）
// 合併後回傳成一個 map
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	// 根據原本 fields 的數量，再加 4（level, time, message, caller）
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["caller"] = l.callers

	// 如果原本有透過 WithFields 添加 field 的話，則把這些 field 放到 data 中
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			// 如果 data 中原本沒有該 key 則把它放進去
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

// Output 會將 Logger 物件以 JSON 格式 Print 出來
func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

// Info 會將參數以 message 放入 Logger 後，以 LevelInfo 的等級輸出
func (l *Logger) Info(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

// Infof 可以把 message 的內容加以 format 後輸出
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

// Fatal 可以把 message 的內容放入 Logger 後輸出，以 LevelFatal 的等級輸出
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

// Fatalf 可以把 message 的內容加以 format 後輸出，以 LevelFatal 的等級輸出
func (l *Logger) Fatalf(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}
