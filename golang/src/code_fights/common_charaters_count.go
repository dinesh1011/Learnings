package main

func commonCharacterCount(s1 string, s2 string) int {

	commonChars := 0
	charsS2 := []rune(s2)

	for _, chars1 := range s1 {

		for index, chars2 := range charsS2 {

			if chars1 == chars2 {
				commonChars++
				charsS2[index] = 0
				//charsS2= append(charsS2[0:index], charsS2[index+1 : ]...)
				break
			}

		}
	}

	return commonChars
}
