package main

import "fmt"

func main() {
	fmt.Println("go")
	fmt.Println(reverseSentence("ф1эхъус апщи дуней"))
}

// проходит по строке с конца и добавляем символы в массив,
// если символ пробел, запускаем еще один цикл,
// который добавляет слово в выводимую строку
// второй цикл начинается с конца, т.к. слово в массив добавлено в перевернутом виде
func reverseSentence(arg string) string {
	var word []string
	var outSentence string
	arg = " " + arg

	for i := len(arg) - 1; i >= 0; i-- {
		if string(arg[i]) == " " || i == 0 {
			for j := len(word) - 1; j >= 0; j-- {
				outSentence += word[j]
			}
			outSentence += " "
			word = word[:0]
			continue
		}
		word = append(word, string(arg[i]))
	}
	return outSentence
}
