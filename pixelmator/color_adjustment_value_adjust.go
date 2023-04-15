package pixelmator

import (
	"fmt"
	"github.com/scryner/streamdeck-pixelmator/applescript"
)

func (r *rangeValue) adjust(v any) error {
	delta, ok := v.(int)
	if !ok {
		return fmt.Errorf("parameter is must be integer")
	}

	current := r.Value
	tobe := current + delta

	switch {
	case delta == 0:
		// never happened
		return nil
	case delta > 0 && current == r.MaxOfRange:
		// do nothing
		return nil
	case delta > 0 && current < r.MaxOfRange && tobe > r.MaxOfRange:
		// set to max value or range
		tobe = r.MaxOfRange
	case delta < 0 && current == r.MinOfRange:
		// do nothing
		return nil
	case delta < 0 && current > r.MinOfRange && tobe < r.MinOfRange:
		// set to min value or range
		tobe = r.MinOfRange
	}

	term := r.adj.getTerm().osascriptTerm
	query := fmt.Sprintf(adjustQueryFormat, fmt.Sprintf("\t\t\tset its %s to %d", term, tobe))

	// run query
	if _, err := applescript.Run(query); err != nil {
		return err
	}

	r.Value = tobe
	return nil
}

func (r *rangeGroup) adjust(v any) error {
	applied, ok := v.(bool)
	if !ok {
		return fmt.Errorf("parameter is must be boolean")
	}

	term := r.adj.getTerm().osascriptTerm
	query := fmt.Sprintf(adjustQueryFormat, fmt.Sprintf("\t\t\tset %v to its %s", applied, term))

	// run query
	if _, err := applescript.Run(query); err != nil {
		return err
	}

	r.IsApplied = applied
	return nil
}

func (r *rangeGroup) adjustChild(child ColorAdjustment, delta int) error {
	adj, ok := r.Group[child]
	if !ok {
		return fmt.Errorf("%s is not children of %s", child.getTerm().osascriptTerm, r.adj.getTerm().osascriptTerm)
	}

	return adj.adjust(delta)
}

const adjustQueryFormat = `if application "Pixelmator Pro" is running then
	tell application "Pixelmator Pro"
		tell the color adjustments of the current layer of the front document
%s
		end tell
	end tell
end if`
