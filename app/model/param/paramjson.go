package param

import (
	"identify/app/model/dto"
)

type RefreshTokenReq struct {
	RefreshToken string `json:"refreshtoken" validate:"required"`
}

type RefreshTokenRes struct {
	TokenInfo dto.TokenInfo
}

type LogoutReq struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refreshtoken" validate:"required"`
}

type GetPassportReq struct {
	Token      string `json:"token" validate:"required"`
	QsId       string `json:"qsId"`
	Product    string `json:"product" `
	Version    string `json:"version" `
	Imei       string `json:"imei" `
	Sdsn       string `json:"sdsn"`
	Securities string `json:"securities" `
	Nohqlist   string `json:"nohqlist" `
	Newwgflag  string `json:"newwgflag" `
	Ip         string `json:"ip"`
}

type GetCookieReq struct {
	Token     string `json:"token" validate:"required"`
	SignValid string `json:"signvalid" validate:"required"` // cookie过期时间, 格式：yyyymmdd
}

type GetSessionIdReq struct {
	Token string `json:"token" validate:"required"`
}

type Res struct {
	Data interface{}
}
