package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	fmt.Println(githubInfo(ctx, "chrisbradleydev"))
}

func Demo() {
	res, err := http.Get("https://api.github.com/users/chrisbradleydev")
	if err != nil {
		log.Fatalf("error: %s", err)
		/* equivalent
		log.Printf("error: %s", err)
		os.Exit(1)
		*/
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", res.Status)
	}

	fmt.Printf("Content-Type: %s\n", res.Header.Get("Content-Type"))

	// if _, err := io.Copy(os.Stdout, res.Body); err != nil {
	// 	log.Fatalf("error: can't copy - %s", err)
	// }

	var r Reply
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	fmt.Println(r)
	// fmt.Printf("%#v\n", r)
}

func githubInfo(ctx context.Context, username string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	// res, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if res.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, res.Status)
	}

	defer res.Body.Close()

	var r struct { // anonymous struct
		Name        string `json:"name"`
		PublicRepos int    `json:"public_repos"`
	}
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.PublicRepos, nil
}

type Reply struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

/* JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int64, int32, int16, int8, uint64, ...
array <-> []any ([]interface{})
object <-> map[string]any struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/