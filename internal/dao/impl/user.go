package daoImpl

import (
	"ci-backend/internal/utils"
	"context"
)

type UserDaoImpl struct {
	wand *utils.Wand
}

var DefaulUserDaoImpl *UserDaoImpl

func NewUserImpl(wand *utils.Wand) {
	DefaulUserDaoImpl.wand = wand
}

func (u *UserDaoImpl) Hello(ctx context.Context) {

}
