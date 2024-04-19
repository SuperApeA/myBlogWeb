package error

const (
	LoginOut = "用户未登录或登录已过期！"
)

var clientError, serverError map[string]bool

func init() {
	clientError, serverError = make(map[string]bool), make(map[string]bool)
	// 用户侧问题
	clientError[LoginOut] = true

	// 服务侧问题
	//serverError[] = true
}

func GetClientError() map[string]bool {
	return clientError
}

func GetServerError() map[string]bool {
	return serverError
}
