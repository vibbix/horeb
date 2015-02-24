package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var blockFlag UnicodeBlock

func init() {
	flag.Var(&blockFlag, "r", "codepoint range")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dumpBlocks := flag.Bool("d", false, "print all Blocks")
	listBlocks := flag.Bool("l", false, "list all Blocks")
	nchars := flag.Int("n", 300, "number of characters to print")
	block := flag.String("b", "geometric", "unicode block by name")
	flag.Parse()

	output := make(chan rune)

	if *dumpBlocks {
		for name, block := range Blocks {
			fmt.Printf("%s: %U, %U\n", name, block.start, block.end)
			block.GenRunes(output)
		}
		for r := range output {
			fmt.Printf("%c ", r)
		}
		os.Exit(0)
	}

	if *listBlocks {
		dumpAll(output)
		os.Exit(0)
	}

	var b *UnicodeBlock
	if blockFlag.end > 0 {
		b.start = blockFlag.start
		b.end = blockFlag.end
		//b.PrintRandom(*nchars)
		os.Exit(0)
	}

	b, valid := Blocks[*block]
	// If invalid block passed in, use a default instead
	if !valid {
		b = Blocks["geometric"]
	}
	b.GenRandomRunes(output, *nchars)
}
