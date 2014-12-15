package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/beefsack/go-geekdo"
)

func main() {
	stderr := log.New(os.Stderr, "", 0)
	r := bufio.NewReader(os.Stdin)
	client := geekdo.NewClient()
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	if err := w.Write([]string{
		"ID",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
	}); err != nil {
		stderr.Fatalf("Error writing line of CSV, %v", err)
	}
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			stderr.Fatalf("Error reading input, %v", err)
		}
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		id, err := strconv.Atoi(trimmed)
		if err != nil {
			stderr.Fatalf("Error getting ID from %s, %v", trimmed, err)
		}
		ratings, err := client.RatingPercentages(id)
		if err != nil {
			stderr.Fatalf("Error getting ratings for %d, %v", id, err)
		}
		if err := w.Write([]string{
			strconv.Itoa(id),
			strconv.Itoa(ratings[1]),
			strconv.Itoa(ratings[2]),
			strconv.Itoa(ratings[3]),
			strconv.Itoa(ratings[4]),
			strconv.Itoa(ratings[5]),
			strconv.Itoa(ratings[6]),
			strconv.Itoa(ratings[7]),
			strconv.Itoa(ratings[8]),
			strconv.Itoa(ratings[9]),
			strconv.Itoa(ratings[10]),
		}); err != nil {
			stderr.Fatalf("Error writing line of CSV, %v", err)
		}
		w.Flush()
	}
}
