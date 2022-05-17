package exchange

import (
	"context"
)

type Exchange interface {
	Receive(ctx context.Context)
	Disconnect()
}
