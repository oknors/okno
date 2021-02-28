package utl

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

var skip = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk,
	unicode.Lm,
}

var safe = []*unicode.RangeTable{
	unicode.Letter,
	unicode.Number,
}

// MakeSlug slugifies a string. The result will only contain lowercase letters,
// digits and dashes. It will not begin or end with a dash, and it
// will not contain runs of multiple dashes.
//
// It is NOT forced into being ASCII, but may contain any Unicode
// characters, with the above restrictions.
func MakeSlug(text string) string {
	buf := make([]rune, 0, len(text))
	dash := false
	for _, r := range norm.NFKD.String(text) {
		switch {
		case unicode.IsOneOf(safe, r):
			buf = append(buf, unicode.ToLower(r))
			dash = true
		case unicode.IsOneOf(skip, r):
		case dash:
			buf = append(buf, '-')
			dash = false
		}
	}
	if i := len(buf) - 1; i >= 0 && buf[i] == '-' {
		buf = buf[:i]
	}
	return string(buf)
}

// // MakeSlugf slugifies a formated string
// func MakeSlugf(format string, a ...interface{}) string {
// 	return MakeSlug(fmt.Sprintf(format, a...))
// }
