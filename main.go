package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	worldMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(worldMessage)
}
