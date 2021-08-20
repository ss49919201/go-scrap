package pattern

import (
	"regexp"
)

func ConvertSpace(src, rpl string) string {
	rp := regexp.MustCompile(`\n *`)
	return rp.ReplaceAllString(src, rpl)
}
