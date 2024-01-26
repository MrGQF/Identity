package param

type GetUserInfoByTokenReq struct {
	Token string `form:"token" validate:"required"`
}
