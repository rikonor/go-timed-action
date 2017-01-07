package timedaction

import (
	"fmt"
	"time"
)

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

func NewTimedActionStore() TimedActionStore {
	return &timedActionStore{store: make(map[string]*TimedAction)}
}

func (tas *timedActionStore) get(id string) (*TimedAction, error) {
	ta, ok := tas.store[id]
	if !ok {
		return nil, fmt.Errorf("no timed-action %s", id)
	}

	return ta, nil
}

func (tas *timedActionStore) Set(id string, ta *TimedAction) error {
	x, _ := tas.get(id)
	if x != nil {
		return fmt.Errorf("timed-action %s already exists", id)
	}

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

func (tas *timedActionStore) Cancel(id string) error {
	ta, err := tas.get(id)
	if err != nil {
		return err
	}

	ta.cancelled = true
	ta.Timer.Reset(0)

	return nil
}
