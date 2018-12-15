package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {

	expect := "/root\n/root/usr\n/root/usr/B\n/root/A\n"

	rootDir := NewDirectory("root")

	usrDir := NewDirectory("usr")
	fileA := NewFile("A")

	rootDir.Add(usrDir)
	rootDir.Add(fileA)

	fileB := NewFile("B")

	usrDir.Add(fileB)

	result := rootDir.Print("")

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
