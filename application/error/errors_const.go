package error

const (
	Success = 0
	Fail    = 1
)

// 常见的网络错误
// const (
// 	ErrTemporaryRedirect   ErrorCode = 307
// 	ErrUnauthorized        ErrorCode = 401
// 	ErrForbidden           ErrorCode = 403
// 	ErrNotFound            ErrorCode = 404
// 	ErrTooManyRequests     ErrorCode = 429
// 	ErrInternalServerError ErrorCode = 500
// 	ErrBadGateway          ErrorCode = 502
// 	ErrServiceUnavaliable  ErrorCode = 503
// 	ErrGatewayTimeout      ErrorCode = 504
// )

// 基础错误
const (
	ErrAuthFailed ErrorCode = iota + 10000
	ErrParameterValidate
	ErrDBConnect
	ErrDBErr
	ErrDBQuery
	ErrDBDelete
	ErrDBOperate
	// app 错误类型
	ErrAppCreated
	ErrAppUpdated
	ErrAppDeleted
	ErrAppNotFound
	// user 类型错误
	ErrUserCreated
	ErrUserUpdated
	ErrUserDeleted
	ErrUserNotFound
)

const (
	UndefinedError ErrorCode = 99999
)

var CCErrors = map[ErrorCode]string{
	Success: "Success",
	Fail:    "Fail",
	// ErrTemporaryRedirect:   "Temporary Redirect",
	// ErrUnauthorized:        "Unauthorized",
	// ErrForbidden:           "Forbidden",
	// ErrNotFound:            "Not Found",
	// ErrTooManyRequests:     "Too Many Requests",
	// ErrInternalServerError: "Internal Server Error",
	// ErrBadGateway:          "Bad Gateway",
	// ErrServiceUnavaliable:  "Service Unavaliable",
	// ErrGatewayTimeout:      "Gateway Timeout",
	// DB
	ErrAuthFailed:        "鉴权失败",
	ErrParameterValidate: "参数校验错误",
	ErrDBConnect:         "DB 连接失败",
	ErrDBErr:             "DB 操作失败",
	ErrDBQuery:           "查询失败",
	ErrDBDelete:          "删除失败",
	ErrDBOperate:         "DB 操作失败",
	UndefinedError:       "未知错误",
	// app
	ErrAppCreated:  "App创建失败",
	ErrAppUpdated:  "App更新失败",
	ErrAppDeleted:  "App删除失败",
	ErrAppNotFound: "App不存在",
	// user
	ErrUserCreated:  "用户创建失败",
	ErrUserUpdated:  "用户更新失败",
	ErrUserDeleted:  "用户删除失败",
	ErrUserNotFound: "用户不存在",
}

func GetErrorMsg(code ErrorCode) string {
	msg, ok := CCErrors[code]
	if !ok {
		return CCErrors[UndefinedError]
	}

	return msg
}
