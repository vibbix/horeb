package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type UnicodeBlock struct {
	start, end rune
}

type BlockMap map[string]*UnicodeBlock

// For info about fonts supporting specific unicode blocks, see for example:
// http://www.fileformat.info/info/unicode/block/index.htm
var Blocks = &BlockMap{

	// Basic Multilingual Plane (0000-ffff)
	// https://en.wikipedia.org/wiki/Plane_(Unicode)#Basic_Multilingual_Plane
	"hebrew":         &UnicodeBlock{0x0590, 0x05ff},
	"currency":       &UnicodeBlock{0x20a0, 0x20cf},
	"letterlike":     &UnicodeBlock{0x2100, 0x214f},
	"misc_technical": &UnicodeBlock{0x2300, 0x23ff},
	"geometric":      &UnicodeBlock{0x25a0, 0x25ff},
	"misc_symbols":   &UnicodeBlock{0x2600, 0x26ff},
	"dingbats":       &UnicodeBlock{0x2700, 0x27bf},

	// Supplementary Multilingual Plane (10000-1ffff)
	// https://en.wikipedia.org/wiki/Plane_(Unicode)#Supplementary_Multilingual_Plane
	"aegean_nums":        &UnicodeBlock{0x10100, 0x1013f},
	"ancient_greek_nums": &UnicodeBlock{0x10140, 0x1018f},
	"phaistos_disc":      &UnicodeBlock{0x101d0, 0x101ff},
	"math_alnum":         &UnicodeBlock{0x1d400, 0x1d7ff},
	"emoji":              &UnicodeBlock{0x1f300, 0x1f5ff},
	"mahjong":            &UnicodeBlock{0x1f000, 0x1f02f},
	"dominos":            &UnicodeBlock{0x1f030, 0x1f09f},
	"playing_cards":      &UnicodeBlock{0x1f0a0, 0x1f0ff},
}

// Returns a *UnicodeBlock at random.
func (bm *BlockMap) RandomBlock() *UnicodeBlock {
	for k, _ := range bm {
		return bm[k]
	}
	return nil
}

// START implement sort.Interface --------------------------------
func (s ByUnicodeBlock) Len() int {
	return len(s)
}

func (s ByUnicodeBlock) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByUnicodeBlock) Less(i, j int) bool {
	return int(s[i].start) < int(s[j].start)
}

// END implement sort.Interface ----------------------------

// List blockmap contents.
func (bm *BlockMap) List() {
	var blocks []*UnicodeBlock
	for _, b := range bm {
		blocks = append(blocks, b)
	}
	sort.Sort(ByUnicodeBlock(blocks))
	for _, b := range blocks {
		fmt.Printf("%5x %5x\n", b.start, b.end)
	}

	//for _, n := range names {
	//	block := Blocks[n]
	//	fmt.Printf("%5x %5x  %s\n", block.start, block.end, n)
	//	if all {
	//		block.Print()
	//		fmt.Println()
	//	}
	//}
}

// Returns a single rune at random from UnicodeBlock.
func (b *UnicodeBlock) RandomRune() rune {
	return rune(rand.Intn(int(b.end-b.start)) + int(b.start) + 1)
}

// Print all printable runes in UnicodeBlock.
func (b *UnicodeBlock) Print() {
	for i := b.start; i <= b.end; i++ {
		if strconv.IsPrint(i) {
			fmt.Printf("%c ", i)
		}
	}
	fmt.Println()
}

// Print n random runes from UnicodeBlock.
func (b *UnicodeBlock) PrintRandom(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", b.RandomRune())
	}
	fmt.Println()
}
