package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Password struct {
	password string
}

var testCases []Password = []Password{{"Ban#ana"}, {"B2ANANA"}, {"bjdf#$%D89"}}

func checkForLengthyPassword(password string) bool {
	return len(password) >= 8
}

func containsUppercaseLetter(password string) bool {
	return strings.ToLower(password) != password
}

func containsLowercaseLetter(password string) bool {
	return strings.ToUpper(password) != password
}

func containsDigit(password string) bool {
	re := regexp.MustCompile(`\d`)
	return re.MatchString(password)
}

func containsSpecialCharacter(password string) bool {
	re := regexp.MustCompile(`[!#/()?_]`)
	return re.MatchString(password)
}

func ratePassword(password string) int {
	rating := 0
	if checkForLengthyPassword(password) {
		rating++
	}
	if containsLowercaseLetter(password) {
		rating++
	}
	if containsUppercaseLetter(password) {
		rating++
	}
	if containsDigit(password) {
		rating++
	}
	if containsSpecialCharacter(password) {
		rating++
	}

	return rating
}

func main() {
	for _, testCase := range testCases {
		passwordRating := ratePassword(testCase.password)
		fmt.Println(passwordRating)
	}
}
