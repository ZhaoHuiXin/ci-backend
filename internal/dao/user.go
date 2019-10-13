package dao

import "context"

type UserDao interface {
	Hello(ctx context.Context)
}
