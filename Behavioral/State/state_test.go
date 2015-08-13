package state

import (
	"testing"
)

func TestState(t *testing.T) {

	expect := "Vrrr... Brrr... Vrrr..." +
			  "Vrrr... Brrr... Vrrr..." +
			  "Белые розы, Белые розы. Беззащитны шипы..."
	
	modile := NewModileAlert()
	
	result := modile.Alert()
	result += modile.Alert()
	
	modile.SetState(&MobileAlertSong{})
	
	result += modile.Alert()
	
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}