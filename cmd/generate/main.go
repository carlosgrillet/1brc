package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

type station struct {
	name string
	mean float64
}

func loadStations(path string) ([]station, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var stations []station
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		parts := strings.SplitN(line, ";", 2)
		if len(parts) != 2 {
			continue
		}
		mean, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			continue
		}
		stations = append(stations, station{name: parts[0], mean: mean})
	}
	return stations, sc.Err()
}

func main() {
	stationsFile := flag.String("stations", "data/weather_stations.csv", "stations CSV file")
	outputFile := flag.String("out", "data/measurements.txt", "output file")
	rows := flag.Int64("rows", 1_000_000_000, "number of rows to generate")
	flag.Parse()

	stations, err := loadStations(*stationsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load stations: %v\n", err)
		os.Exit(1)
	}
	if len(stations) == 0 {
		fmt.Fprintln(os.Stderr, "no stations loaded")
		os.Exit(1)
	}

	out, err := os.Create(*outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create output: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	w := bufio.NewWriterSize(out, 1<<20) // 1 MB buffer
	n := int64(len(stations))

	for i := int64(0); i < *rows; i++ {
		s := stations[rand.Int64N(n)]
		temp := s.mean + rand.NormFloat64()*10.0
		// clamp to realistic range
		if temp < -99.9 {
			temp = -99.9
		} else if temp > 99.9 {
			temp = 99.9
		}
		fmt.Fprintf(w, "%s;%.1f\n", s.name, temp)
	}

	if err := w.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "flush: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "generated %d rows -> %s\n", *rows, *outputFile)
}
