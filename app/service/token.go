package service

import (
	"encoding/base64"
	"encoding/json"
	h "gitee.com/QunXiongZhuLu/KongMing/http"
	"gitee.com/QunXiongZhuLu/KongMing/tars-protocol/account"
	it "gitee.com/QunXiongZhuLu/KongMing/tars-protocol/identify"
	bm "gitee.com/QunXiongZhuLu/kratos/pkg/net/http/blademaster"
	"identify/app/model/dto"
	"identify/app/model/param"
	"strconv"
)

func (s *Service) ValidateToken(token string, common *h.Common) (cert *it.CertificationInfo, extern *dto.TokenExtern, err error) {
	var (
		externByte []byte
	)

	// 校验Token
	cert, err = s.TarsCheckToken(common.AppId, token, common.Version)
	if err != nil {
		return nil, nil, err
	}

	// 解析Token返回值
	externByte, err = base64.StdEncoding.DecodeString(cert.Extern)
	if err != nil {
		return
	}
	extern = new(dto.TokenExtern)
	err = json.Unmarshal(externByte, extern)
	if err != nil {
		return
	}

	return cert, extern, nil
}

func (s *Service) GetUserInfoByToken(req *param.GetUserInfoByTokenReq, common *h.Common) (info *dto.UserInfo, err error) {
	var (
		extern *dto.TokenExtern
	)

	// 校验Token
	_, extern, err = s.ValidateToken(req.Token, common)
	if err != nil {
		return
	}

	info = &dto.UserInfo{
		UserId:  extern.UserId,
		Account: extern.NickName,
	}
	return
}

func (s *Service) RefreshToken(ctx *bm.Context, req *param.RefreshTokenReq) (res *param.RefreshTokenRes, err error) {
	var (
		tokenInfo *it.TokenInfo
		common    *h.Common
	)

	common = h.GetCommon(ctx)

	tokenInfo, err = s.TarsRefreshToken(common.AppId, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	res = &param.RefreshTokenRes{
		TokenInfo: struct {
			Token           string
			ExpiresIn       string
			RefreshToken    string
			RefreshExpireIn string
		}{Token: tokenInfo.AccessToken, ExpiresIn: tokenInfo.Access_expire, RefreshToken: tokenInfo.RefreshToken, RefreshExpireIn: tokenInfo.Refresh_expire},
	}

	return res, nil
}

func (s *Service) Logout(ctx *bm.Context, req *param.LogoutReq) (err error) {
	var (
		common *h.Common
	)

	common = h.GetCommon(ctx)

	// 校验Token
	_, _, err = s.ValidateToken(req.Token, common)
	if err != nil {
		return
	}

	// 清除登录状态
	err = s.TarsClearToken(common.AppId, req.RefreshToken)
	if err != nil {
		return
	}

	return
}

func (s *Service) GetPCPassport(ctx *bm.Context, req *param.GetPassportReq) (passport string, err error) {
	var (
		extern *dto.TokenExtern
		common *h.Common
		info   = new(account.GetPassportRes)
	)

	// 签名校验
	common = h.GetCommon(ctx)
	if common.AppSecret, err = s.TarsGetAppSecret(common.AppId, common.AppKey); err != nil {
		return
	}
	/*err = validate.ValidateSign(*req, *common)
	if err != nil {
		return "", err
	}*/

	// 校验Token
	_, extern, err = s.ValidateToken(req.Token, common)
	if err != nil {
		return "", err
	}

	authPassportReq := &account.PCPassport{
		UserId:     strconv.Itoa(extern.UserId),
		QsId:       req.QsId,
		Product:    req.Product,
		Version:    req.Version,
		Imei:       req.Imei,
		Sdsn:       req.Sdsn,
		Securities: req.Securities,
		Nohqlist:   req.Nohqlist,
		Newwgflag:  req.Newwgflag,
		Ip:         ctx.RemoteIP(),
	}

	err = s.TarsGetPCPassport(authPassportReq, info)
	if err != nil {
		return "", err
	}

	return info.Passport, nil
}

func (s *Service) GetCookie(ctx *bm.Context, req *param.GetCookieReq, res *account.GetCookieRes) (err error) {
	var (
		common *h.Common
		extern *dto.TokenExtern
	)

	// 校验Token
	common = h.GetCommon(ctx)
	_, extern, err = s.ValidateToken(req.Token, common)
	if err != nil {
		return
	}

	if err = s.TarsGetCookie(int32(extern.UserId), req.SignValid, res); err != nil {
		return
	}

	return
}
