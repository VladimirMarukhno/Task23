package main
/*Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а возвращает 2D-массив,
где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении i (строку надо разбить на слова и взять последнее).*/
import (
	"fmt"
	"strings"
)

func parseTest(sentences []string, chars []rune) [][]interface{} {
	var interfaceArray [][]interface{}
	var ch int
	for a := 0; a < len(sentences); a++ {
		sentences[a] = strings.ToUpper(sentences[a])
		st := sentences[a]
		for index := range chars {
			for i, r := range st {
				if chars[index] == r {
					ch = i
				}
			}
			if ch > -1 {
				interfaceArray = append(interfaceArray, []interface{}{chars[index], ch})
			}
			ch = -1
		}
	}
	return interfaceArray
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'M'}
	interfaceArray := parseTest(sentences, chars)
	for i := 0; i < len(interfaceArray); i++ {
		fmt.Printf("%q position: %d\n", interfaceArray[i][0], interfaceArray[i][1])
	}
}