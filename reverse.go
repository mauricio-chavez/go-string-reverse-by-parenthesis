package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Checks if parenthesis are balanced with a stack
func areParenthesisBalanced(str string) bool {
	var stack []rune
	for _, char := range str {
		switch char {
		case '(':
			stack = append(stack, char) // push
		case ')':
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		}
	}
	return len(stack) == 0 // If there are no parenthesis left, these are balanced
}

// Ulititary reverse function
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Takes a string and returns same string with its inner substring
// in parenthesis in reverse and its length, starting to count
// from specified starting index
func hereBeDragons(str string, start int) (string, int) {
	i := start
loop:
	for i < len(str) {
		switch str[i] {
		case '(':
			// If we have an opening parenthesis, we pass the original
			// string but starting one index later so we can pivot
			// starting parenthesis. This returns same string with its
			// inner substring in parenthesis in reverse and its length,
			// so we can assign this new string to our original string,
			// and we can know how many steps we should skip to keep
			// algorithm's linear complexity
			newStr, skip := hereBeDragons(str, i+1)
			str = newStr
			i += skip
		case ')':
			// If we encounter a closing parenthesis, we reverse string
			// from opening to closing parenthesis
			subStr := reverse(str[start:i])

			// We sliced and append substring using the following caveats
			// - (start - 1) is index before opening parenthesis
			// - (i+1) is index after closing parenthesis
			// Then, we create a new substring, replacing its inner substring
			// in parenthesis with this reversed substring
			newStr := str[:start-1] + subStr + str[i+1:]

			// Here, I printed out every middle step :P
			fmt.Printf("> %v", newStr)

			// This means we reached the last closing parenthesis
			// and we can break the loop
			if i == len(newStr) {
				break loop
			}

			// So I returned string without parenthesis and its length
			return newStr, len(subStr)
		default:
			// Here, we check if we finished our string and we didn't encountered
			// a closing parenthesis, so we reverse the remaining string
			if i == len(str)-1 {
				fmt.Println(reverse(str))
			}
			// This is here in order to increment our index
			i++
		}
	}

	// And this is here because I had to return something, even if
	// in the end I didn't used it :P
	return "", 0
}

// Rinnegan is just a wrapper for my magical function :P
func Rinnegan(str string) {
	balanced := areParenthesisBalanced(str)
	if !balanced {
		log.Fatalln("Parenthesis are not balanced")
	}
	// Starts my function at index = 0
	hereBeDragons(str, 0)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string to be reversed:\n-> ")

	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	Rinnegan(text)
}
