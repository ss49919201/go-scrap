package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 1000000

var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), "")
	for _, v := range lines {
		slice = append(slice, toInt(v))
	}
	return slice
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type mem map[int]int

func main() {
	S := readIntSlice()

	// スプリットかどうか
	// ピン1が倒れているかどうか
	if S[0] == 0 {
		if S[6] == 1 {
			if S[1] == 1 || S[7] == 1 {
				if S[3] == 0 {
					fmt.Println("Yes")
					return
				}
			}
			if S[4] == 1 {
				if S[3] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[1] == 0 && S[7] == 0 {
					fmt.Println("Yes")
					return
				}
			}
			if S[2] == 1 || S[8] == 1 {
				if S[3] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[1] == 0 && S[7] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
			}

			if S[5] == 1 {
				if S[3] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[1] == 0 && S[7] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[2] == 0 && S[8] == 0 {
					fmt.Println("Yes")
					return
				}
			}

			if S[9] == 1 {
				if S[3] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[1] == 0 && S[7] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[2] == 0 && S[8] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[5] == 0 {
					fmt.Println("Yes")
					return
				}
			}
		}

		if S[3] == 1 {
			if S[4] == 1 {
				if S[1] == 0 && S[7] == 0 {
					fmt.Println("Yes")
					return
				}
			}
			if S[2] == 1 || S[8] == 1 {
				if S[1] == 0 && S[7] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
			}

			if S[5] == 1 {
				if S[1] == 0 && S[7] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[2] == 0 && S[8] == 0 {
					fmt.Println("Yes")
					return
				}
			}

			if S[9] == 1 {
				if S[1] == 0 && S[7] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[2] == 0 && S[8] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[5] == 0 {
					fmt.Println("Yes")
					return
				}
			}
		}

		if S[1] == 1 || S[7] == 1 {
			if S[2] == 1 || S[8] == 1 {
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
			}

			if S[5] == 1 {
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[2] == 0 && S[8] == 0 {
					fmt.Println("Yes")
					return
				}
			}

			if S[9] == 1 {
				if S[4] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[2] == 0 && S[8] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[5] == 0 {
					fmt.Println("Yes")
					return
				}
			}
		}

		if S[4] == 1 {
			if S[5] == 1 {
				if S[2] == 0 && S[8] == 0 {
					fmt.Println("Yes")
					return
				}
			}

			if S[9] == 1 {
				if S[2] == 0 && S[8] == 0 {
					fmt.Println("Yes")
					return
				}
				if S[5] == 0 {
					fmt.Println("Yes")
					return
				}
			}
		}

		if S[2] == 1 || S[8] == 1 {
			if S[9] == 1 {
				if S[5] == 0 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}

	fmt.Println("No")
}
