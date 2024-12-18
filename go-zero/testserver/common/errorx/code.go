package errorx

const Ok = 200

var (
	Success          = NewError(200, "success")
	DBError          = NewError(10000, "db is error")
	AlreadyRegister  = NewError(10100, "user already register")
	NameOrPwdError   = NewError(10101, "username or password error")
	TokenError       = NewError(10102, "token error")
	TokenExpired     = NewError(10103, "token expired")
	TokenInvalid     = NewError(10104, "token invalid")
	UserNotExist     = NewError(10105, "user not exist")
	UserAlreadyExist = NewError(10106, "user already exist")
)
