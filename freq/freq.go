package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	mapDemo()

	w, err := mostCommon(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(w)
}

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", nil
	}
	return maxWord(freqs)
}

// "Who's on first?" -> [Who s on first]
var re = regexp.MustCompile(`[a-zA-Z]+`)

func mapDemo() {
	var stocks map[string]float64 // symbol -> price
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("%s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	// stocks = make(map[string]float64)
	// stocks[sym] = 136.73
	stocks = map[string]float64{
		sym:    136.73,
		"AAPL": 172.35,
	}
	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	for k := range stocks { // keys
		fmt.Println(k)
	}

	for k, v := range stocks { // key & value
		fmt.Println(k, "->", v)
	}

	for _, v := range stocks { // values
		fmt.Println(v)
	}

	delete(stocks, "AAPL")
	fmt.Println(stocks)
	delete(stocks, "AAPL") // no panic
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}
	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count
	lnum := 0
	for s.Scan() {
		lnum++
		words := re.FindAllString(s.Text(), -1) // current line
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	fmt.Println("num lines:", lnum)

	return freqs, nil
}