package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

/*
 * Complete the 'BurstyRateLimiter' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 *  3. int batchSize
 *  4. int toAdd
 */
var sum int
var mux sync.Mutex

func BurstyRateLimiter(requestChan chan bool, resultChan chan int, batchSize int, toAdd int) {
	for {
		_, ok := <-requestChan
		if !ok {
			close(requestChan)
		}
		for i := 0; i < batchSize; i++ {
			resultChan <- sum
			mux.Lock()
			sum += toAdd
			mux.Unlock()
		}
	}
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//skipBatches := int(skipTemp)
	skipBatches := 0
	//
	//totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//printBatches := int(totalTemp)
	printBatches := 4

	//batchSizeTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//batchSize := int(batchSizeTemp)
	batchSize := 2

	//toAddTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//toAdd := int(toAddTemp)
	toAdd := 3

	resultChan := make(chan int)
	requestChan := make(chan bool)
	go BurstyRateLimiter(requestChan, resultChan, batchSize, toAdd)
	for i := 0; i < skipBatches+printBatches; i++ {
		start := time.Now().UnixNano()
		requestChan <- true
		for j := 0; j < batchSize; j++ {
			new := <-resultChan
			if i < skipBatches {
				continue
			}
			fmt.Println(new)
		}
		if i >= skipBatches && i != skipBatches+printBatches-1 {
			fmt.Println("-----")
		}
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000000
		if timeDiff < 3 {
			fmt.Println("Rate is too high")
			break
		}
	}
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
