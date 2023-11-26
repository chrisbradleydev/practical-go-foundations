package nlp

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

// var tokenizeCases = []struct{
// 	text string
// 	tokens []string
// } {
// 	{"Who's on first?", []string{"who", "s", "on", "first"}},
// 	{"", nil},
// }

type tokenizeCase struct {
	Text   string
	Tokens []string
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	// data, err := os.ReadFile("tokenize_cases.toml")
	// require.NoError(t, err)

	var testCases struct {
		Cases []tokenizeCase
	}

	// err = toml.Unmarshal(data, &testCases)
	_, err := toml.DecodeFile("tokenize_cases.toml", &testCases)
	require.NoError(t, err, "Unmarshal TOML")
	return testCases.Cases
}

func TestTokenizeTable(t *testing.T) {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
	/* Before testify
	// if tokens != expected {} // can't compare slices with == in Go (only to nil)
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
	*/
}