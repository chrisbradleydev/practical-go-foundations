package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
	/* Before testify
	// if tokens != expected {} // can't compare slices with == in Go (only to nil)
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
	*/
}