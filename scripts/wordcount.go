package main

import (
	"fmt"
	"strings"
)

func count(words []string) (wc map[string]int) {
	wc = make(map[string]int)
	for i := 0; i < len(words); i++ {
		wc[words[i]]++
	}
	return
}

func pluralise(x int, w string) (plural string) {
	plural = w
	if x != 1 {
		plural = w + "s"
	}
	return
}

func main() {
	s := "I am a string of words and I have many words in me"
	wrds := strings.Fields(s)

	wc := count(wrds)

	for k, v := range wc {
		fmt.Printf("The word [%s] appeared [%d] %v\n", k, v, (pluralise(v, "time")))
	}

}
