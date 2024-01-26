package tars

import (
	"context"
)

type validateImpl struct {
}

func (validate *validateImpl) CheckAlive(tarsCtx context.Context, req string, res *string) (ret int32, err error) {
	*res = req
	return 0, nil
}
