package notifier

// Notifier is the interface for sending notifications.
type Notifier interface {
	Notify([]byte) error
}
