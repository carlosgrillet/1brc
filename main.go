package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// storage = [count, min, avg, max]
var storage = make(map[string]*[]float64)

const (
	COUNT int = iota
	MIN
	AVG
	MAX
)

func main() {
	data := flag.String("file", "data/measurements.txt", "weather file to parse")
	flag.Parse()

	file, err := os.Open(*data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Printf("cities: %d\n", len(storage))
			return
		}
		values := strings.Split(string(line), ";")
		city := values[0]
		temp, _ := strconv.ParseFloat(values[1], 64)

		makeCalculation(city, temp)
	}
}

func makeCalculation(city string, temp float64) {
	currentValues, ok := storage[city]
	if !ok {
		storage[city] = &[]float64{1, temp, temp, temp}
		return
	}

	(*currentValues)[COUNT] += 1

	if temp < (*currentValues)[MIN] {
		(*currentValues)[MIN] = temp
		(*currentValues)[AVG] = ((*currentValues)[AVG] + temp) / (*currentValues)[COUNT]
		return
	}

	if temp > (*currentValues)[MAX] {
		(*currentValues)[MAX] = temp
	}

	(*currentValues)[AVG] = ((*currentValues)[AVG] + temp) / (*currentValues)[COUNT]
}

func findHalves(file *os.File) (int64, int64) {
	var offset int64

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Printf("could not stat file. %v", err)
		os.Exit(1)
	}

	halfFile := fileStat.Size() / 2
	file.Seek(halfFile, 1)

// offsetTillNewLine will advance the file pointer and return the ammount of
// bytes scanned till it find a new line character
func offsetTillNewLine(file *os.File) int {
	currentChar := make([]byte, 1)
	var offset int
	for {
		if file.Read(currentChar); currentChar[0] != '\n' {
			offset += 1
			continue
		}
		break
	}
	return offset
}
