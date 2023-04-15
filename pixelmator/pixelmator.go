package pixelmator

import (
	"fmt"
	"sync"
)

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

func (px *Pixelmator) adjust(adj ColorAdjustment, v any) error {
	px.lock.Lock()
	defer px.lock.Unlock()

	if adjustable, ok := px.values[adj]; !ok {
		return fmt.Errorf("invalid data: adjustment %s is not registered", adj.getTerm().osascriptTerm)
	} else {
		return adjustable.adjust(v)
	}
}

func (px *Pixelmator) registerAdjustment(adj ColorAdjustment) error {
	px.lock.Lock()
	defer px.lock.Unlock()

	px.values[adj] = adj.newValue()
	return syncColorAdjustments(px.values)
}

func (px *Pixelmator) deregisterAdjustment(adj ColorAdjustment) error {
	px.lock.Lock()
	defer px.lock.Unlock()

	delete(px.values, adj)
	return syncColorAdjustments(px.values)
}
