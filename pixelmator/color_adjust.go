package pixelmator

import (
	"fmt"
	"github.com/scryner/streamdeck-pixelmator/applescript"
)

func (r *RangeValue) Adjust(delta int) error {
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
	query := fmt.Sprintf(adjustQueryFormat, fmt.Sprintf("\t\t\tset %d to its %s", tobe, term))

	// run query
	if _, err := applescript.Run(query); err != nil {
		return err
	}

	r.Value = tobe
	return nil
}

func (rg *RangeGroup) Apply(applied bool) error {
	term := rg.adj.getTerm().osascriptTerm
	query := fmt.Sprintf(adjustQueryFormat, fmt.Sprintf("\t\t\tset %v to its %s", applied, term))

	// run query
	if _, err := applescript.Run(query); err != nil {
		return err
	}

	rg.IsApplied = applied
	return nil
}

func (rg *RangeGroup) AdjustChild(child ColorAdjustment, delta int) error {
	adj, ok := rg.Group[child]
	if !ok {
		return fmt.Errorf("%s is not children of %s", child.getTerm().osascriptTerm, rg.adj.getTerm().osascriptTerm)
	}

	return adj.Adjust(delta)
}

const adjustQueryFormat = `if application "Pixelmator Pro" is running then
	tell application "Pixelmator Pro"
		tell the color adjustments of the current layer of the front document
%s
		end tell
	end tell
end if`
