package main

import (
	"github.com/kyokomi/emoji/v2"
	"os"
	"testing"
)

func TestCreateTextWithEmoji(t *testing.T) {
	msg := createTextWithEmoji()
	rightText := emoji.Sprint("Hello :world_map:!")
	if msg != rightText {
		t.Fatalf(`CreateTextWithEmoji() = %s, but it must be equal to %s`, msg, rightText)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
