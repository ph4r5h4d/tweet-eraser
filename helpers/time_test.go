package helpers

import (
	"testing"
	"time"
)

const TEN_DAYS_IN_HOURS = 240

func TestIsDeletable(t *testing.T) {
	t1 := time.Now().Add(time.Duration(-TEN_DAYS_IN_HOURS) * time.Hour)
	if !IsDeletable(t1, 7) {
		t.Fail()
	}

	if IsDeletable(t1, 11) {
		t.Fail()
	}

	if !IsDeletable(t1, 10) {
		t.Fail()
	}
}
