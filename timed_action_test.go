package timedaction

import (
	"testing"
	"time"
)

func TestSetTimedAction(t *testing.T) {
	ta := NewTimedActionStore()

	// Make sure timed-actions get triggered

	triggered := false
	if err := ta.Set("test", &TimedAction{
		Action: func() { triggered = true },
		Timer:  time.NewTimer(50 * time.Millisecond),
	}); err != nil {
		t.Fatal(err)
	}

	// wait for timed-action to trigger
	time.Sleep(100 * time.Millisecond)

	if !triggered {
		t.Fatal("failed to trigger timed-action")
	}
}

func TestCancelTimedAction(t *testing.T) {
	ta := NewTimedActionStore()

	// Make sure you can cancel timed-actions

	triggered := false
	ta.Set("test", &TimedAction{
		Action: func() { triggered = true },
		Timer:  time.NewTimer(50 * time.Millisecond),
	})

	if err := ta.Cancel("test"); err != nil {
		t.Fatal(err)
	}

	// wait for timed-action to trigger
	time.Sleep(100 * time.Millisecond)

	if triggered {
		t.Fatal("failed to cancel timed-action")
	}
}
