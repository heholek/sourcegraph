package app

import "testing"

func TestGuessRepoURIFromRemoteURL(t *testing.T) {
	tests := map[string]string{
		"git@github.com:a/b.git":          "github.com/a/b",
		"ssh://git@github.com/a/b.git":    "github.com/a/b",
		"ssh://github.com/a/b.git":        "github.com/a/b",
		"ssh://github.com:1234/a/b.git":   "github.com/a/b",
		"https://github.com:1234/a/b.git": "github.com/a/b",
		"http://alice@foo.com:1234/a/b":   "foo.com/a/b",
	}
	for input, want := range tests {
		got := guessRepoURIFromRemoteURL(input)
		if got != want {
			t.Errorf("%s: got %q, want %q", input, got, want)
		}
	}
}
