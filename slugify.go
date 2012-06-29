package slugify

import(
	"regexp"
	"strings"
)

func S(str string) string {
	str = strings.ToLower(strings.Trim(str, " "))						// Trim than Lower-Case
	
	from := "aáäâeéëeiíiîoóöőôuúüűunç·/_,:;"
	to   := "aaaaeeeeiiiiooooouuuuunc------"
	i := 0
	for _, v := range from {
		str = strings.Replace(str, string(v), string(to[i]), -1)		// Remove accents
		i++
	}
	
	str = regexp.MustCompile("[^a-z0-9-]").ReplaceAllString(str, "-")	// Remove invalid chars (spaces are invalid)
	str = regexp.MustCompile("-+").ReplaceAllString(str, "-")			// Collapse dashes
	
	return str
}
