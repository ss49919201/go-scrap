package main

import (
	"fmt"
	"regexp"
)

func convertColons(s string, exceptions []string) string {
	compileStr := ` :(`
	for i, v := range exceptions {
		compileStr += v
		if i < len(exceptions)-1 {
			compileStr += `|`
		}
	}
	compileStr += `) `
	re := regexp.MustCompile(compileStr)
	fmt.Println(re.FindAllString(s, -1))
	return re.ReplaceAllString(s, "${1}${2}${3}")
}

func main() {
	input := "hello :abc :world :der :example"
	output := convertColons(input, []string{"abc", "def"})
	fmt.Println(output) // Output: hello::abc:world::der:example
}
