package dto

type TokenExtern struct {
	UserId   int
	NickName string
}

type UserInfo struct {
	UserId  int
	Account string
}

type TokenInfo struct {
	Token           string // 当 token 超时后，可以调用token刷新接口进行刷新，每次刷新生成新的token
	ExpiresIn       string // Token过期时间，秒数
	RefreshToken    string // 当 refreshtoken 失效的后，需要用户重新登录
	RefreshExpireIn string // refreshToken过期小时数，秒数
}
