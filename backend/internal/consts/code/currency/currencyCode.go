package currency

const (
	SUCCESS      = 0   // 成功
	ERROR        = 7   // 通用错误
	UNAUTHORIZED = 401 // 未授权
	FORBIDDEN    = 403 // 禁止访问
	NOT_FOUND    = 404 // 资源不存在
	SERVER_ERROR = 500 // 服务器错误
	// ... 可以添加更多错误码
)
