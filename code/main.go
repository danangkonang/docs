package main

import (
	"fmt"
	"strings"
)

func sortArray(arr []int) []int {
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func isPalindrome(inputString string) bool {
	// Menghapus spasi dan mengubah huruf menjadi huruf kecil
	cleanedString := strings.ReplaceAll(strings.ToLower(inputString), " ", "")

	// Mengonversi string menjadi slice of runes
	runes := []rune(cleanedString)

	// Membalikkan slice of runes
	reversedRunes := make([]rune, len(runes))
	for i, j := len(runes)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversedRunes[j] = runes[i]
	}

	// Membandingkan slice of runes dengan slice of runes yang telah dibalikkan
	return string(runes) == string(reversedRunes)
}

func fizzBuzz(input uint) string {
	if input%15 == 0 {
		return "FizzBuzz"
	}
	if input%3 == 0 {
		return "Fizz"
	}
	if input%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(input)
}

func main() {
	fmt.Println(sortArray([]int{2, 8, 6, 3, 1}))
	fmt.Println(isPalindrome("A man a plan a canal Panama"))
	fmt.Println(fizzBuzz(25))
}
