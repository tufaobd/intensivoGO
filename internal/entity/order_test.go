package entity

import "testing"

func TestIfGetAnyError(t *testing.T) {
	order := Order{}
	if order.Validate() == nil {
		t.Error("Precisa de ter ID")
	}
}
