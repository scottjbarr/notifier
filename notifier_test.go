package notifier

import (
	"context"
	"reflect"
	"testing"
)

func TestNotifier(t *testing.T) {
	m := NewMockNotifier()

	ctx := context.Background()

	if err := m.Notify(ctx, "hello"); err != nil {
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
