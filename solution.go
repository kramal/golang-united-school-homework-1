package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func createTextWithEmoji() string {
	return emoji.Sprint("Hello :world_map:!")
}

func main() {
	worldMessage := createTextWithEmoji()
	fmt.Println(worldMessage)
}
