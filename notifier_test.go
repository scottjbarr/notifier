package notifier

import (
	"reflect"
	"testing"
)

func TestNotifier(t *testing.T) {
	m := NewMockNotifier()

	if err := m.Notify("hello"); err != nil {
		t.Fatal(err)
	}

	want := []string{
		"hello",
	}

	got := m.Notifications

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}
