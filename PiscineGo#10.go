package z01dakar

import (
	"errors"
	"os"
	"unicode/utf8"
)

// PrintRune prints a single rune (Unicode code point) and returns any error if the encoding or the writing fails.
func PrintRune(r rune) error {
	l := utf8.RuneLen(r)
	if l == -1 {
		return errors.New("ERROR: invalid value to encode in UTF-8")
	}
	p := make([]byte, l)
	utf8.EncodeRune(p, r)
	_, err := os.Stdout.Write(p)
	return err
}

// PrintStr prints a string et and returns any error if the writing fails.
func PrintStr(str string) error {
	var err error
	for _, r := range str {
		err = PrintRune(r)
		if err != nil {
			break
		}
	}
	return err
}
