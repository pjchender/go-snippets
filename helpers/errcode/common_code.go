package errcode

// 定義所有的 ErrCode
var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "internal server error")
	InvalidParams             = NewError(10000001, "invalid parameters")
	NotFound                  = NewError(10000002, "not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "驗證失敗，找不到對應的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "驗證失敗，Token 錯誤")
	UnauthorizedTokenTimeout  = NewError(10000005, "驗證失敗，Token 逾時")
	UnauthorizedTokenGenerate = NewError(10000006, "驗證失敗，Token 產生失敗")
	TooManyRequest            = NewError(10000007, "請求過多")
)
