package domain

import "context"

type Consumer interface {
	Consume(ctx context.Context, handle func(key, value []byte) error) error
	Close() error
}
