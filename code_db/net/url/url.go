package main

import (
	"log"
	"net/url"
)

func TestPathEscapeUnescape() {
	src := []string{
		"a/b/c",
		"a b c",
	}

	for _, s := range src {
		escaped := url.PathEscape(s)

		unescaped, err := url.PathUnescape(escaped)
		if err != nil {
			log.Fatalf("failed to unescape: %s, %s, %v", s, escaped, err)
		}

		log.Printf("escaped: [%s], [%s], [%s]\n", s, escaped, unescaped)
	}
}

func main() {
	TestPathEscapeUnescape()
}
