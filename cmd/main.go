package main

import (
	"flag"

	"github.com/daniel-felipe/hn"
)

func main() {
    total := flag.Int("t", 3, "The amount of news to display")
    flag.Parse()

    hn.ShowLatests(*total) 
}
