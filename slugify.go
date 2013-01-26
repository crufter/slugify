package slugify

import (
	"regexp"
	"strings"
)

var replacements = map[rune]string{
	'á':  "a",
	'à':  "a",
	'ä':  "a",
	'â':  "a",
	'é':  "e",
	'è':  "e",
	'ë':  "e",
	'ê':  "e",
	'í':  "i",
	'ì':  "i",
	'i':  "i",
	'î':  "i",
	'ó':  "o",
	'ò':  "o",
	'ö':  "o",
	'ő':  "o",
	'ô':  "o",
	'ú':  "u",
	'ù':  "u",
	'ü':  "u",
	'ű':  "u",
	'ç':  "c",
	'·':  "-",
	'–':  "-",
	'–':  "-", // – (U+2013) != -
	'/':  "-",
	'_':  "-",
	',':  "-",
	':':  "-",
	';':  "-",
	'\'': "", // collapses "adam's" to "adams"
}

func S(str string) string {
	str = strings.ToLower(strings.TrimSpace(str)) // Trim all whitespace, then Lower-Case
	newstr := ""

	// Iterate over string, for each character check if there is a replacement, if so, append it to newstr, if not, append original value to newstr
	for _, char := range str {
		if replacement, ok := replacements[char]; ok {
			newstr += replacement
		} else {
			newstr += string(char)
		}
	}

	newstr = regexp.MustCompile("[^a-z0-9-]").ReplaceAllString(newstr, "-") // Remove invalid chars (spaces are invalid)
	newstr = regexp.MustCompile("-+").ReplaceAllString(newstr, "-")         // Collapse dashes

	return newstr
}
