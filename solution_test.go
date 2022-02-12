package main

import (
	"github.com/kyokomi/emoji/v2"
	"testing"
)

func TestGetMessage(t *testing.T) {
	msg := GetMessage()
	rightText := emoji.Sprint("Hello :world_map:!")
	if msg != rightText {
		t.Fatalf(`GetMessage() = %s, but it must be equal to %s`, msg, rightText)
	}
}