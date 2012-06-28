package slugify

import(
	"regexp"
	"fmt"
	"strings"
)

func S(str string) string {
	r, _ := regexp.Compile("^ +| +$")						// Trim
	str = string(r.ReplaceAll([]byte(str), []byte("")))
	str = strings.ToLower(str)
	
	from := "aáäâeéëeiíiîoóöőôuúüűunç·/_,:;"
	to   := "aaaaeeeeiiiiooooouuuuunc------"
	i := 0
	for _, v := range from {
		str = strings.Replace(str, string(v), string(to[i]), -1)	// Remove accents
		i++
	}
	
	r, _ = regexp.Compile("[^a-z0-9 -]")
	str = string(r.ReplaceAll([]byte(str), []byte("")))		// Remove invalid chars
	r, _ = regexp.Compile(" +") 
	str = string(r.ReplaceAll([]byte(str), []byte("-")))	// Collapse whitespace and replace by -
	r, _ = regexp.Compile("-+")
	str = string(r.ReplaceAll([]byte(str), []byte("-")))	// Collapse dashes
	
	return str
}
