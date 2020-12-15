package global

var AppSetting = struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}{
	DefaultPageSize: 10,
	MaxPageSize:     100,
	LogSavePath:     "storage/logs",
	LogFileName:     "app",
	LogFileExt:      ".log",
}
