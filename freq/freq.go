package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	mapDemo()

	// w, err := mostCommon(file)
	w, err := getFreqs(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	pl := rankByWordCount(w, 5)
	for _, p := range pl {
		fmt.Printf("%s: %d\n", p.Key, p.Value)
	}
}

// func mostCommon(r io.Reader) (string, error) {
// 	freqs, err := wordFrequency(r)
// 	if err != nil {
// 		return "", nil
// 	}
// 	return maxWord(freqs)
// }
func getFreqs(r io.Reader) (map[string]int, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return make(map[string]int), err
	}
	return freqs, nil
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

// func maxWord(freqs map[string]int) (string, error) {
// 	if len(freqs) == 0 {
// 		return "", fmt.Errorf("empty map")
// 	}

// 	maxN, maxW := 0, ""
// 	for word, count := range freqs {
// 		if count > maxN {
// 			maxN, maxW = count, word
// 		}
// 	}
// 	return maxW, nil
// }

func rankByWordCount(wordFrequencies map[string]int, total int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl[:total]
}

type Pair struct {
	Key string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}
func (p PairList) Swap(i, j int){
	p[i], p[j] = p[j], p[i]
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