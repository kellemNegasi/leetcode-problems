package solutions

import "strings"

func ReverseWords(s string) string {
	words := strings.Fields(s)
	l := len(words)
	for i:=0; i<l/2; i++{
		currentWord := words[i]
		words[i] = words[l-1-i]
		words[l-1-i]= currentWord
	}

	return strings.Join(words, " ")
}
