package pixelmator

import "testing"

func TestSync(t *testing.T) {
	m := make(map[ColorAdjustment]ColorAdjustmentValue)
	m[Exposure] = newRangeValue(Exposure)
	m[Saturation] = newRangeValue(Saturation)
	if vig, err := newRangeGroup(Vignette); err != nil {
		t.Errorf("failed to make range group: %v", err)
		t.FailNow()
	} else {
		m[Vignette] = vig
	}

	// sync
	err := syncColorAdjustments(m)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for k, v := range m {
		t.Logf("%s: %v", k.getTerm().osascriptTerm, v)
	}
}
