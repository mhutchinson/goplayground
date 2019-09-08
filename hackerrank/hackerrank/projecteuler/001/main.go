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
	var r int64
	var next3 int32 = 3
	var next5 int32 = 5
	for next3 < n || next5 < n {
		if next3 == next5 {
			r += int64(next3)
			next3 += 3
			next5 += 5
		} else if next3 < next5 {
			r += int64(next3)
			next3 += 3
		} else {
			r += int64(next5)
			next5 += 5
		}
	}
	return r
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
