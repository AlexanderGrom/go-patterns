package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {

	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Error("Objects are not equal!\n")
	}
}
