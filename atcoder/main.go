package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	N := sc.Text()
	ss := strings.Split(N, " ")
	fmt.Println(ss)
}
