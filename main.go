package main

import "fmt"

func main() {
	// var firstName string = "Anggi"
	// var lastName = "Dian"
	// // fmt.Printf("halo %s %s!\n", firstName, lastName)
	// fmt.Println("halo", firstName, lastName+"!")

	// var positiveNumber uint8 = 89
	// var negativeNumber = -12343423644
	// fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan positif: %d\n", negativeNumber)

	// var decimalNumber = 2.62
	// fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// var message = `Nama saya "Dian Anggi".
	// Mari belajar "Golang".`
	// fmt.Println(message)

	// const firstName string = "Anggi"
	// fmt.Print("halo", firstName, "!\n")
	// var value = (((2+6)%3)*4 - 2) / 3
	// var isEqual = (value == 2)

	// fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// var left = false
	// var right = true
	// var leftAndRight = left && right
	// fmt.Printf("left && right\t(%t) \n", leftAndRight)

	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

}
