/*
 *  main.cpp
 *  Программа для расчёта рулонов обоев под размеры стен
 *
 *  Created by Maksim Sergeev on 31.01.2023.
 */

package main

import (
	"fmt"
	"math"
)

func square(a float64, b float64) float64 {
	return a * b
}

func main() {
	for {

		fmt.Println("Расчёт количества рулонов обоев")

		fmt.Print("Введите количество стен: ")
		var count int
		fmt.Scan(&count)
		var length float64
		var height float64
		var sqRoom float64

		fmt.Println("Высота у всех стен одинаковая? 1 если ДА, 0 если НЕТ")
		var flagH bool
		fmt.Scan(&flagH)

		if flagH {
			fmt.Print("Введите высоту для всех стен: ")
			fmt.Scan(&height)
			for i := 0; i < count; i++ {
				fmt.Printf("Введите длину %d-й стены в см: ", i+1)
				fmt.Scan(&length)
				sqRoom += square(length, height)
			}
		} else {
			for i := 0; i < count; i++ {
				fmt.Printf("Введите длину %d-й стены в см: ", i+1)
				fmt.Scan(&length)
				fmt.Printf("Введите высоту %d-й стены в см: ", i+1)
				fmt.Scan(&height)
				sqRoom += square(length, height)
			}
		}

		var sqRulon float64 = 0
		sqRulon = square(100, 1000)

		fmt.Println("\n\n*******************************************************************")
		fmt.Printf("Рулонов МЕТРОВОЙ ширины:\t%0.2f", sqRoom/sqRulon)
		fmt.Print(";  Придется купить:\t", math.Ceil(sqRoom/sqRulon))
		fmt.Printf("\nРулонов ПОЛУМЕТРОВОЙ ширины:\t%0.2f", sqRoom/(sqRulon/2))
		fmt.Print(";  Придется купить:\t", math.Ceil(sqRoom/(sqRulon/2)))
		fmt.Println("\n*******************************************************************\n\n")
	}
	//	var pause uint
	//	fmt.Scan(&pause)
}
