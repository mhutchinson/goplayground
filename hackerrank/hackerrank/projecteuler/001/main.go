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
	r3 := arithmeticSum(3, (n-1)/3)
	r5 := arithmeticSum(5, (n-1)/5)
	r15 := arithmeticSum(15, (n-1)/15)

	return r3 + r5 - r15
}

func arithmeticSum(m, i int32) int64 {
	first := m
	last := m * i
	return int64(i) * int64(first+last) >> 1
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
