package timedaction

import "time"

// TimedAction is an action that will occur based on a timer
type TimedAction struct {
	Action    func()
	Timer     *time.Timer
	cancelled bool
}

type TimedActionStore interface {
	Set(id string, ta *TimedAction) error
	Cancel(id string) error
}

type timedActionStore struct {
	store map[string]*TimedAction
}

// NewTimedActionStore creates a new in-memory TimedActionStore
func NewTimedActionStore() TimedActionStore {
	return &timedActionStore{store: make(map[string]*TimedAction)}
}

func (tas *timedActionStore) get(id string) (*TimedAction, bool, error) {
	ta, ok := tas.store[id]
	if !ok {
		return nil, false, nil
	}

	return ta, true, nil
}

// Set creates or overwrites a TimedAction
func (tas *timedActionStore) Set(id string, ta *TimedAction) error {
	tas.store[id] = ta

	go func() {
		<-ta.Timer.C
		if !ta.cancelled {
			ta.Action()
		}
		delete(tas.store, id)
	}()

	return nil
}

// Cancel cancels a TimedAction if it exists
func (tas *timedActionStore) Cancel(id string) error {
	ta, ok, err := tas.get(id)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	ta.cancelled = true
	ta.Timer.Reset(0)

	return nil
}
