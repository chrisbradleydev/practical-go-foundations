package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	/*
	a := 1
	{
		a := 3                  // shadows outer a
		// a = 2                // change outer a
		fmt.Println("inner", a) // print inner a
	}
	fmt.Println("outer", a)
	return
	*/

	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)

	sig, err = sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

func sha1Sum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close() // defer called in LIFO order

	var r io.Reader = file
	if strings.HasSuffix(filename, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}

	// io.CopyN(os.Stdout, r, 100)
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}