package pixelmator

import (
	"fmt"
	"math"
	"strconv"
)

type ColorAdjustmentValue interface {
	getTerms() []term
	setValues(map[string]string) error
	adjust(v any) error
}

type rangeInterface interface{}

type rangeValue struct {
	adj        ColorAdjustment
	Value      int
	MinOfRange int
	MaxOfRange int
	step       int
	rangeInterface
}

func (r *rangeValue) getTerms() []term {
	return []term{r.adj.getTerm()}
}

func (r *rangeValue) setValues(m map[string]string) error {
	if v, ok := m[r.adj.getTerm().osascriptVariable]; ok {
		return parseInt(v, &r.Value)
	}

	return nil
}

func newRangeValue(adj ColorAdjustment) *rangeValue {
	min, max := adj.getRange()
	totalStep := math.Abs(float64(min) + float64(max))
	step := int(totalStep) / 200

	return &rangeValue{
		adj:        adj,
		Value:      0,
		MinOfRange: min,
		MaxOfRange: max,
		step:       step,
	}
}

type rangeGroup struct {
	adj       ColorAdjustment
	IsApplied bool
	Group     map[ColorAdjustment]*rangeValue
}

func (r *rangeGroup) getTerms() []term {
	terms := []term{r.adj.getTerm()}

	for _, child := range r.Group {
		terms = append(terms, child.getTerms()...)
	}

	return terms
}

func (r *rangeGroup) setValues(m map[string]string) error {
	if s, ok := m[r.adj.getTerm().osascriptVariable]; ok {
		if err := parseBool(s, &r.IsApplied); err != nil {
			return err
		}
	}

	var adjs []ColorAdjustment
	for k := range r.Group {
		adjs = append(adjs, k)
	}

	for _, adj := range adjs {
		if s, ok := m[adj.getTerm().osascriptVariable]; ok {
			if err := parseInt(s, &r.Group[adj].Value); err != nil {
				return err
			}
		}
	}

	return nil
}

func newRangeGroup(adj ColorAdjustment) (*rangeGroup, error) {
	children := adj.getChildren()
	if len(children) < 1 {
		return nil, fmt.Errorf("color adjustment '%s' is not range group", adj.getTerm().osascriptTerm)
	}

	m := make(map[ColorAdjustment]*rangeValue)
	for _, child := range children {
		m[child] = newRangeValue(child)
	}

	return &rangeGroup{
		adj:   adj,
		Group: m,
	}, nil
}

func parseInt(s string, i *int) error {
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}

	*i = int(i64)
	return nil
}

func parseBool(s string, b *bool) error {
	bb, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}

	*b = bb
	return nil
}
