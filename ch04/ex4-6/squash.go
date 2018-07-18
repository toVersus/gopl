// Package squash squashes Unicode spaces into ASCII space.
package squash

import (
	"unicode"
	"unicode/utf8"
)

func squash(b []byte) []byte {
	var res []byte
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if unicode.IsSpace(r) {
			r = rune(32)
		}
		res = append(res, []byte(string(r))...)
		b = b[size:]
	}
	return res
}
