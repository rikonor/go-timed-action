Timed Actions
---

### Usage

```go
tas := timedaction.NewTimedActionStore()

ta := &timedaction.TimedAction{
  Action: func() { fmt.Println("action") },
  Timer: time.NewTimer(10 * time.Second),
}

// Set a timed-action to be triggered
tas.Set("xyz", ta)

// Can cancel the timed-action
tas.Cancel("xyz")
```
