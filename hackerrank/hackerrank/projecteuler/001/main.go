package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// MultiplesSum sums up the multiples of 3 and 5 less than n
func MultiplesSum(n int32) int64 {
	// every n % 15 the pattern repeats, with 7 values:
	vs := [7]int32{3, 5, 6, 9, 10, 12, 15}
	var r int64
	var c int32
	var cycles int32
	for {
		c = cycles * 15
		for _, v := range vs {
			current := c + v
			if current >= n {
				return r
			}
			r += int64(current)
		}
		cycles++
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	testCases, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var i int64
	for ; i < testCases; i++ {
		n, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		fmt.Fprintf(writer, "%d\n", MultiplesSum(int32(n)))
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
