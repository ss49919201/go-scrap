package main

import (
	"bufio"
	"fmt"
	"math"
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
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, toInt(v))
	}
	return slice
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

var mem = make(map[int]int)

func main() {
	A := readIntSlice()
	B := readIntSlice()
	C := readIntSlice()
	D := readIntSlice()

	y := "Yes"
	n := "No"

	// ABCの内角
	ab := []float64{float64(B[1] - A[1]), float64(B[0] - A[0])}
	ad := []float64{float64(D[1] - A[1]), float64(D[0] - A[0])}
	fmt.Println(math.Atan2(math.Abs(ab[1]-ad[1]), math.Abs(ab[0]-ad[0])) * 360 / math.Pi)
	if math.Abs((math.Atan2(float64(C[1]-B[1]), float64(C[0]-B[0]))-math.Atan2(float64(A[1]-B[1]), float64(A[0]-B[0])))/math.Pi*180) >= 180 {
		fmt.Println(n)
		return
	}
	// BCDの内角
	// BCA+ACD
	a := math.Abs((math.Atan2(float64(A[1]-C[1]), float64(A[0]-C[0])) - math.Atan2(float64(B[1]-C[1]), float64(B[0]-C[0]))) / math.Pi * 180)
	b := math.Abs((math.Atan2(float64(D[1]-C[1]), float64(D[0]-C[0])) - math.Atan2(float64(A[1]-C[1]), float64(A[0]-C[0]))) / math.Pi * 180)
	if a+b == 360 {
		if (360-math.Max(a, b))+math.Min(a, b) >= 180 {
			fmt.Println(n)
			return
		}
	} else {
		if a+b >= 180 {
			fmt.Println(n)
			return
		}
	}
	// CDAの内角
	if math.Abs((math.Atan2(float64(A[1]-D[1]), float64(A[0]-D[0]))-math.Atan2(float64(C[1]-D[1]), float64(C[0]-D[0])))/math.Pi*180) >= 180 {
		fmt.Println(n)
		return
	}
	// DABの内角
	// DAC+CAB
	a = math.Abs((math.Atan2(float64(C[1]-A[1]), float64(C[0]-A[0])) - math.Atan2(float64(D[1]-A[1]), float64(D[0]-A[0]))) / math.Pi * 180)
	b = math.Abs((math.Atan2(float64(B[1]-A[1]), float64(B[0]-A[0])) - math.Atan2(float64(C[1]-A[1]), float64(C[0]-A[0]))) / math.Pi * 180)
	if a+b == 360 {
		if (360-math.Max(a, b))+math.Min(a, b) >= 180 {
			fmt.Println(n)
			return
		}
	} else {
		if a+b >= 180 {
			fmt.Println(n)
			return
		}
	}

	fmt.Println(y)
}
