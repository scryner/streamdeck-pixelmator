package pixelmator

import "sync"

type Pixelmator struct {
	values map[ColorAdjustment]ColorAdjustmentValue

	quit chan bool
	lock sync.RWMutex
}

func NewPixelmator() *Pixelmator {
	// start sync

	return &Pixelmator{
		values: make(map[ColorAdjustment]ColorAdjustmentValue),
	}
}

func (px *Pixelmator) Stop() {
	px.quit <- true
}
