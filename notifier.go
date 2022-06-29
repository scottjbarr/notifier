package notifier

import (
	"context"
)

// Notifier is the interface for sending notifications.
type Notifier interface {
	Notify(context.Context, []byte) error
}
