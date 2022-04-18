package main

/*Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных. */
import (
	"fmt"
	"math/rand"
	"time"
)

func evenOdd(input []int) (even []int, odd []int) {
	for _, elem := range input {
		if (elem % 2) != 0 {
			odd = append(odd, elem)
		} else {
			even = append(even, elem)
		}
	}
	return
}

func main() {
	var size int
	fmt.Println("Введите размер массива")
	fmt.Scan(&size)
	mySlice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size-1; i++ {
		mySlice[i] = rand.Intn(10 * size)
	}
	even, odd := evenOdd(mySlice)
	fmt.Printf("Массив четных чисел: %d\nМассив нечётных чисел: %d\n", even, odd)
}
