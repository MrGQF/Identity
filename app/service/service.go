package service

import (
	"gitee.com/QunXiongZhuLu/kratos/pkg/conf/paladin"
	bm "gitee.com/QunXiongZhuLu/kratos/pkg/net/http/blademaster"
	"identify/app/dao"
)

type Service struct {
	ac     *paladin.Map
	job    *paladin.Map
	dao    dao.DaoImpl
	client *bm.Client
}

func NewService(d dao.DaoImpl, c *bm.Client) (s *Service, cf func(), err error) {
	s = &Service{
		ac:     &paladin.TOML{},
		job:    &paladin.TOML{},
		dao:    d,
		client: c,
	}
	cf = func() {}
	err = paladin.Watch("application.yaml", s.ac)
	return
}
