#!/usr/bin/env python3

import random
import sys

UNICODE_BLOCKS = {
	# Basic Multilingual Plane (0000-ffff)
	# https://en.wikipedia.org/wiki/Plane_(Unicode)#Basic_Multilingual_Plane
	"hebrew":         (0x0590, 0x05ff),
	"currency":       (0x20a0, 0x20cf),
	"letterlike":     (0x2100, 0x214f),
	"misc_technical": (0x2300, 0x23ff),
	"geometric":      (0x25a0, 0x25ff),
	"misc_symbols":   (0x2600, 0x26ff),
	"dingbats":       (0x2700, 0x27bf),
	# Supplementary Multilingual Plane (10000-1ffff)
	# https://en.wikipedia.org/wiki/Plane_(Unicode)#Supplementary_Multilingual_Plane
	"aegean_nums":        (0x10100, 0x1013f),
	"ancient_greek_nums": (0x10140, 0x1018f),
	"phaistos_disc":      (0x101d0, 0x101ff),
	"math_alnum":         (0x1d400, 0x1d7ff),
	"emoji":              (0x1f300, 0x1f5ff),
	"mahjong":            (0x1f000, 0x1f02f),
	"dominos":            (0x1f030, 0x1f09f),
	"playing_cards":      (0x1f0a0, 0x1f0ff),
}


def randchr(block):
    """Returns a random character from the unicode block."""
    assert isinstance(block, tuple)
    assert len(block) == 2

    return chr(random.choice(range(*block)) + 1)


def gen_randchrs(n, block_list):
    """Generates n random chars from list of unicode blocks."""
    assert isinstance(n, int), "n must be an integer"
    assert isinstance(block_list, (tuple, list)), "block_list must be a tuple or a list"
    assert block_list, "block list must not be empty"
    
    for _ in range(n):
        b = random.choice(block_list)
        yield randchr(b)


if __name__ == '__main__':
    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument('-n', type=int, default=80,
                        help='Number of characters to print')
    parser.add_argument('-b', metavar="BLOCKNAME", nargs='+', default=['geometric'], help='block to use')
    parser.add_argument('-l', action="store_true", help='list all blocks')
    parser.add_argument('-d', action="store_true", help='dump all blocks')
    conf = parser.parse_args()

    if conf.d:
        for b in UNICODE_BLOCKS.values():
            for c in range(*b):
                print(chr(c), end=' ')
        print()
        sys.exit(0)

    if conf.l:
        for k,v in UNICODE_BLOCKS.items():
            print("{:5x} {:5x}  {:20s}".format(v[0],v[1],k))
        sys.exit(0)

    for c in gen_randchrs(conf.n, tuple(UNICODE_BLOCKS.values())):
        print(c, end=' ')
    print()
