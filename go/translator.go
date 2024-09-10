package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Braille to English and English to Braille mapping
var brailleToEnglish = map[string]string{
	"O.....": "a", "O.O...": "b", "OO....": "c", "OO.O..": "d", "O..O..": "e",
	"OOO...": "f", "OOOO..": "g", "O.OO..": "h", ".OO...": "i", ".OOO..": "j",
	"O...O.": "k", "O.O.O.": "l", "OO..O.": "m", "OO.OO.": "n", "O..OO.": "o",
	"OOO.O.": "p", "OOOOO.": "q", "O.OOO.": "r", ".OO.O.": "s", ".OOOO.": "t",
	"O...OO": "u", "O.O.OO": "v", ".OOO.O": "w", "OO..OO": "x", "OO.OOO": "y", "O..OOO": "z",
}

var englishToBraille = map[string]string{
	"a": "O.....", "b": "O.O...", "c": "OO....", "d": "OO.O..", "e": "O..O..",
	"f": "OOO...", "g": "OOOO..", "h": "O.OO..", "i": ".OO...", "j": ".OOO..",
	"k": "O...O.", "l": "O.O.O.", "m": "OO..O.", "n": "OO.OO.", "o": "O..OO.",
	"p": "OOO.O.", "q": "OOOOO.", "r": "O.OOO.", "s": ".OO.O.", "t": ".OOOO.",
	"u": "O...OO", "v": "O.O.OO", "w": ".OOO.O", "x": "OO..OO", "y": "OO.OOO", "z": "O..OOO",
}

var brailleToNumber = map[string]string{
	"O.....": "1", "O.O...": "2", "OO....": "3", "OO.O..": "4", "O..O..": "5",
	"OOO...": "6", "OOOO..": "7", "O.OO..": "8", ".OO...": "9", ".OOO..": "0",
}

var numberToBraille = map[string]string{
	"1": "O.....", "2": "O.O...", "3": "OO....", "4": "OO.O..", "5": "O..O..",
	"6": "OOO...", "7": "OOOO..", "8": "O.OO..", "9": ".OO...", "0": ".OOO..",
}

func brailleToText(text string) string {
	charSlice := []string{}
	for i := 0; i < len(text)-5; i++ {
		newString := text[i : i+6]
		charSlice = append(charSlice, newString)
		i += 5
	}
	isNumber := false
	isCapital := false
	convertedSlice := []string{}
	for _, v := range charSlice {
		_, existsV := brailleToEnglish[v]
		if existsV && isNumber {
			convertedSlice = append(convertedSlice, brailleToNumber[v])
		} else if existsV && isCapital {
			capitalV := strings.ToUpper(brailleToEnglish[v])
			convertedSlice = append(convertedSlice, capitalV)
			isCapital = false
		} else if existsV {
			convertedSlice = append(convertedSlice, brailleToEnglish[v])
		} else if v == ".O.OOO" {
			isNumber = true
		} else if v == ".....O" {
			isCapital = true
		} else if v == "......" {
			convertedSlice = append(convertedSlice, " ")
			isNumber = false
		}
	}
	convertedString := strings.Join(convertedSlice, "")

	return convertedString
}

func textToBraille(text string) string {
	// text = strings.ToLower(text)
	chars := strings.Split(text, "")

	brailleText := []string{}
	firstNumber := false
	for _, v := range chars {
		if v == " " {
			firstNumber = false
			brailleText = append(brailleText, "......")
		} else if _, err := strconv.Atoi(v); err == nil && !firstNumber {
			firstNumber = true
			brailleText = append(brailleText, ".O.OOO")
			brailleText = append(brailleText, numberToBraille[v])
		} else if _, err := strconv.Atoi(v); err == nil && firstNumber {
			// brailleText = append(brailleText, ".O.OOO")
			brailleText = append(brailleText, numberToBraille[v])
		} else if strings.ToUpper(v) == v {
			// fmt.Println(v)
			brailleText = append(brailleText, ".....O")
			brailleText = append(brailleText, englishToBraille[strings.ToLower(v)])
		} else {
			brailleText = append(brailleText, englishToBraille[v])
		}

	}

	convertedString := strings.Join(brailleText, "")

	return convertedString
}

func main() {
	input := strings.Join(os.Args[1:], " ")
	if strings.Contains(input, ".") {
		fmt.Println(brailleToText(input))
	} else {
		fmt.Println(textToBraille(input))
	}
}
