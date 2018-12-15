package state

import (
	"testing"
)

func TestState(t *testing.T) {

	expect := "Vrrr... Brrr... Vrrr..." +
		"Vrrr... Brrr... Vrrr..." +
		"Белые розы, Белые розы. Беззащитны шипы..."

	mobile := NewMobileAlert()

	result := mobile.Alert()
	result += mobile.Alert()

	mobile.SetState(&MobileAlertSong{})

	result += mobile.Alert()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
