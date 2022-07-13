package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	infile := os.Args[1]

	buf, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	filescanner := bufio.NewScanner(buf)
	scanner := bufio.NewScanner(os.Stdin)
	replacer := strings.NewReplacer(
		"a", "?l",
		"b", "?l",
		"c", "?l",
		"d", "?l",
		"e", "?l",
		"f", "?l",
		"g", "?l",
		"h", "?l",
		"i", "?l",
		"j", "?l",
		"k", "?l",
		"l", "?l",
		"m", "?l",
		"n", "?l",
		"o", "?l",
		"p", "?l",
		"q", "?l",
		"r", "?l",
		"s", "?l",
		"t", "?l",
		"u", "?l",
		"v", "?l",
		"w", "?l",
		"x", "?l",
		"y", "?l",
		"z", "?l",
		"A", "?u",
		"B", "?u",
		"C", "?u",
		"D", "?u",
		"E", "?u",
		"F", "?u",
		"G", "?u",
		"H", "?u",
		"I", "?u",
		"J", "?u",
		"K", "?u",
		"L", "?u",
		"M", "?u",
		"N", "?u",
		"O", "?u",
		"P", "?u",
		"Q", "?u",
		"R", "?u",
		"S", "?u",
		"T", "?u",
		"U", "?u",
		"V", "?u",
		"W", "?u",
		"X", "?u",
		"Y", "?u",
		"Z", "?u",
		"0", "?d",
		"1", "?d",
		"2", "?d",
		"3", "?d",
		"4", "?d",
		"5", "?d",
		"6", "?d",
		"7", "?d",
		"8", "?d",
		"9", "?d",
		" ", "?s",
		"!", "?s",
		"\"", "?s",
		"#", "?s",
		"$", "?s",
		"%", "?s",
		"&", "?s",
		"\\", "?s",
		"(", "?s",
		")", "?s",
		"*", "?s",
		"+", "?s",
		",", "?s",
		"-", "?s",
		".", "?s",
		"'", "?s",
		"/", "?s",
		":", "?s",
		";", "?s",
		"<", "?s",
		"=", "?s",
		">", "?s",
		"?", "?s",
		"@", "?s",
		"[", "?s",
		"]", "?s",
		"^", "?s",
		"_", "?s",
		"`", "?s",
		"{", "?s",
		"|", "?s",
		"}", "?s",
		"~", "?s",
	)
	var tokens []string

	for filescanner.Scan() {
		tokens = append(tokens, filescanner.Text())

		if err := filescanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	for scanner.Scan() {
		mask := replacer.Replace(scanner.Text())

		for _, value := range tokens {
			tokenmask := replacer.Replace(value)
			if strings.Contains(mask, tokenmask) {

				// format mask token chars into sub chars
				newword := strings.Replace(mask, tokenmask, value, -1)
				newword = strings.Replace(newword, "?u", "?", -1)
				newword = strings.Replace(newword, "?l", "?", -1)
				newword = strings.Replace(newword, "?d", "?", -1)
				newword = strings.Replace(newword, "?s", "?", -1)

				// loop over the string and finish the sub
				for i, c := range scanner.Text() {
					for x, y := range newword {
						if string(y) == "?" && x == i {
							newword = replaceAtIndex(newword, rune(c), i)
							break
						}
					}

				}
				fmt.Println(newword)
			}

			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}

	}

}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
