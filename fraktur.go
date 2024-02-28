package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var reverseFlag = flag.Bool("r", false, "reverse the text transformation")

var frakturMap = map[string]string{
	"a": "𝔞",
	"b": "𝔟",
	"c": "𝔠",
	"d": "𝔡",
	"e": "𝔢",
	"f": "𝔣",
	"g": "𝔤",
	"h": "𝔥",
	"i": "𝔦",
	"j": "𝔧",
	"k": "𝔨",
	"l": "𝔩",
	"m": "𝔪",
	"n": "𝔫",
	"o": "𝔬",
	"p": "𝔭",
	"q": "𝔮",
	"r": "𝔯",
	"s": "𝔰",
	"t": "𝔱",
	"u": "𝔲",
	"v": "𝔳",
	"w": "𝔴",
	"x": "𝔵",
	"y": "𝔶",
	"z": "𝔷",
	"ä": "𝔞" + "̈",
	"ö": "𝔬" + "̈",
	"ü": "𝔲" + "̈",
	"ß": "𝔰" + "̸",
	"A": "𝔄",
	"B": "𝔅",
	"C": "ℭ",
	"D": "𝔇",
	"E": "𝔈",
	"F": "𝔉",
	"G": "𝔊",
	"H": "ℌ",
	"I": "ℑ",
	"J": "𝔍",
	"K": "𝔎",
	"L": "𝔏",
	"M": "𝔐",
	"N": "𝔑",
	"O": "𝔒",
	"P": "𝔓",
	"Q": "𝔔",
	"R": "ℜ",
	"S": "𝔖",
	"T": "𝔗",
	"U": "𝔘",
	"V": "𝔙",
	"W": "𝔚",
	"X": "𝔛",
	"Y": "𝔜",
	"Z": "ℨ",
	"Ä": "𝔄" + "̈",
	"Ö": "𝔒" + "̈",
	"Ü": "𝔘" + "̈",
	"ẞ": "𝔖" + "̸",
	"0": "𝟘",
	"1": "𝟙",
	"2": "𝟚",
	"3": "𝟛",
	"4": "𝟜",
	"5": "𝟝",
	"6": "𝟞",
	"7": "𝟟",
	"8": "𝟠",
	"9": "𝟡",
}

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *reverseFlag {
			fmt.Println(reverseTransform(line))
		} else {
			fmt.Println(transform(line))
		}
	}
}

func transform(input string) string {
	output := input
	for key, val := range frakturMap {
		output = strings.ReplaceAll(output, key, val)
	}
	return output
}

func reverseTransform(input string) string {
	output := input
	for key, val := range frakturMap {
		output = strings.ReplaceAll(output, val, key)
	}
	return output
}

