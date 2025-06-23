package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Print("Введите число:")
	fmt.Scanln(&a)
	fmt.Print("Введите число:")
	fmt.Scanln(&b)
	fmt.Print("Введите число:")
	fmt.Scanln(&c)
	fmt.Print("Введите число:")
	fmt.Scanln(&d)
	fmt.Print("Введите число:")
	fmt.Scanln(&e)

	myArray := [5]int{a, b, c, d, e}
	n := len(myArray)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if myArray[j] < myArray[j+1] {
				myArray[j], myArray[j+1] = myArray[j+1], myArray[j]
			}
		}
	}
	var average = float64(a+b+c+d+e) / 5
	fmt.Printf("%d, %d, %d, %d, %d\n", myArray[0], myArray[1], myArray[2], myArray[3], myArray[4])
	fmt.Println("Самое большое число:", myArray[0])
	fmt.Println("Самое маленькое число:", myArray[4])
	fmt.Println("Среднее арифметическое число:", average)
}

