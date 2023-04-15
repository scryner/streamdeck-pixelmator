package pixelmator

import "testing"

func TestParseAdjustment(t *testing.T) {
	ev := "com.scryner.pixelmator.adjust.1"
	adj, err := parseAdjustment(ev)
	if err != nil {
		t.Errorf("failed to parse adjustment: %v", err)
		t.FailNow()
	}

	t.Log("adj:", adj.getTerm())
}
