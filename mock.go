package notifier

import (
	"context"
	"strings"
)

// MockNotifier is a simple mock Notifier for test usage.
type MockNotifier struct {
	Notifications []string
}

// NewMockNotifier returns a new MockNotifier.
func NewMockNotifier() *MockNotifier {
	return &MockNotifier{}
}

// Notify stores a notification for reading later.
func (s *MockNotifier) Notify(ctx context.Context, msg string) error {
	msg = strings.Replace(msg, "\n", " ", -1)

	s.Notifications = append(s.Notifications, msg)

	return nil
}
