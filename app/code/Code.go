package code

type ECode struct {
	Code int
	Msg  string
}

var (
	/**公共错误定义*/
	SuccessCode = ECode{Code: 200, Msg: "success"}
	AuthFailed  = ECode{Code: 401, Msg: "Unauthorized Login!"}
	SystemError = ECode{Code: 400, Msg: "System error! Please try again later!"}
)
