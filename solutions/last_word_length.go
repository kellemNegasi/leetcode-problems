package solutions

func LengthOfLastWord(s string) int {
    // words := strings.Fields(s)
	// r := len(words[len(words)-1])
	r := getResult(s)
	return r
}

func getResult(s string) int{
	l := len(s)
	count := 0
	stringFound := false
	for i:=l-1; i>=0; i--{
		currentChar := string(s[i])
		if currentChar != " "{
			stringFound = true
			count++
		}
		
		if stringFound && currentChar==" "{
			break
		}
	}
	return count
}