package main

import (
	"fmt"
	"log"
	"math"
)

//Here are some simple examples of exercises done in Go, they can serve as an example for someone who wants to learn Go.

func main() {
	var i, j int
	//Make a function that takes a number and returns if it is perfect. A number is perfect when it is equal to the sum of its divisors except itself
	fmt.Println("Enter a number you want to know if it's perfect.")
	if _, err := fmt.Scan(&i); err != nil {
		log.Print("\nError scanning the value, due to:", err)
		return
	}
	if PerfectNumber(i) {
		fmt.Println("The number is perfect")
	} else {
		fmt.Println("The number is not perfect.")
	}
	//Using binary search method, create a function that allows you to search for a value in an array of N elements and return whether it was found or not.
	numbers := []int{1, 4, 5, 17, 22, 34, 46, 58, 90}
	fmt.Println("Enter a number to search using binary method.")
	if _, err := fmt.Scan(&i); err != nil {
		log.Print("\nError scanning the value, due to:", err)
		return
	}
	fmt.Println(BinarySearch(i, numbers))

	//Create a method that takes a number and returns its factorial using recursion.
	fmt.Println("Enter the value for which you want to find the factorial.")
	if _, err := fmt.Scan(&i); err != nil {
		log.Print("\nError scanning the value, due to:", err)
		return
	}

	fmt.Scanf("%d", &i)
	fmt.Println("The result of the factorial is: ", FactorialNumber(i))
	//Create a method that takes in two numbers and returns whether they are consecutive primes (there are no primes between them).
	fmt.Println("Enter two numbers to identify if they are consecutive primes.")
	if _, err := fmt.Scan(&i); err != nil {
		log.Print("\nError scanning the value, due to:", err)
		return
	}

	if _, err := fmt.Scan(&j); err != nil {
		log.Print("\nError scanning the value, due to:", err)
		return
	}

	if ConsecutivePrimeNumbers(i, j) {
		fmt.Println("The numbers are consecutive primes.")
	} else {
		fmt.Println("The numbers are not consecutive primes.")
	}
}

// Functions
// We have to make a function to find the divisors of it and count if it is perfect.
func PerfectNumber(num int) bool {

	sum := 0
	res := 0
	for i := 1; i < num; i++ {
		res = num % i
		if res == 0 {
			sum = sum + i
		}
	}
	if sum == num {
		return true
	} else {
		return false
	}

}

// Create the binary function to find the value in the array.
func BinarySearch(number int, vector []int) bool {

	low := 0
	high := len(vector) - 1

	for low <= high {
		mean := (low + high) / 2

		if vector[mean] < number {
			low = mean + 1
		} else {
			high = mean - 1
		}
	}

	if low == len(vector) || vector[low] != number {
		return false
	}

	return true
}

// Now we create the function to return the factorial of a number
func FactorialNumber(number int) int {
	var value int
	if number == 1 {
		return 1
	}
	value = FactorialNumber(number-1) * number
	return value
}

// This is the most complicated function, as we first need to identify if the numbers are prime and then if they are consecutive.
func ConsecutivePrimeNumbers(num1, num2 int) bool {
	if Prime(num1) {
		if Prime(num2) {
			comp := num2 - 1
			for num1 < comp {
				num1++
				if Prime(num1) {
					return false
				}
			}
			return true
		} else {
			fmt.Println(num2, "Not a prime number ")
			return false
		}
	} else {
		fmt.Println(num1, "Not a prime number ")
		return false
	}
}

// With this function we identify if the numbers are prime.
func Prime(num int) bool {
	isPrime := true
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime = false
		}
	}
	if isPrime {
		return true
	} else {
		return false
	}
}
