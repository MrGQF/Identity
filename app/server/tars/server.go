package tars

import "identify/app/service"

var (
	svc *service.Service
)

func New(s *service.Service) (imp *validateImpl, cf func(), err error) {
	svc = s

	imp = new(validateImpl)

	cf = func() {}
	return imp, cf, nil
}
