package notifier

import (
	"log"
)

// MultiNotifierNotifier is a wrapper for multiple notifiers.
type MultiNotifier struct {
	Notifiers []Notifier
}

func NewMultiNotifier() *MultiNotifier {
	return &MultiNotifier{
		Notifiers: []Notifier{},
	}
}

func (m *MultiNotifier) Register(n Notifier) {
	m.Notifiers = append(m.Notifiers, n)
}

// Notify will make best effort to notify each Notifier, but will only log errors rather than return
// them.
func (m *MultiNotifier) Notify(b []byte) error {
	for _, n := range m.Notifiers {
		if err := n.Notify(b); err != nil {
			log.Printf("ERROR Notifier = %v payload = %s : err = %v", n, b, err)
		}
	}

	return nil
}
