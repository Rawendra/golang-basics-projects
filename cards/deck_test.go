package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeckFunction(t *testing.T) {
	d := newDeck()
	fmt.Println(len(d))
	if len(d) != 16 {
		t.Errorf("Expecting the length to 16, but got %v", len(d))
	}
}

func TestSaveToFileAndnewDeckFromFile(t *testing.T) {
	filename := "testingFile"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)
	rd := newDeckFromFile(filename)

	if len(rd) != 16 {
		t.Errorf("Expecting the length to 16, but got %v", len(d))
	}

}
