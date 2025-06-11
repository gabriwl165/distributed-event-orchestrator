package domain

import "context"

type Publisher interface {
	Publish(ctx context.Context, topic string, key, value []byte) error
}
