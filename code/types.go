package code

import (
	"fmt"
	"unicode"
)

// passwordValid ...
func passwordValid(password string) bool {

	// convertimos el string a slice de runas
	passToRune := []rune(password)

	// comprobamos la validacion de cantidad de caracteres
	if len(passToRune) < 8 {
		return false
	}
	if len(passToRune) > 15 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, r := range passToRune {

		// pkg unicode verifica runas
		// verifica si la runa dada comple con la validación
		if unicode.IsUpper(r) {
			hasUpper = true
		}
		if unicode.IsLower(r) {
			hasLower = true
		}
		if unicode.IsNumber(r) {
			hasNumber = true
		}
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

// CheckValidation ...
func CheckValidation(password string) {

	if passwordValid(password) {
		fmt.Println("Password good")
	} else {
		fmt.Println("Password bad")
	}
}

// code.CheckValidation("")
// code.CheckValidation("This!I5A")
// code.CheckValidation("Auto7&los")

// -------------------------------------------------------------------------------------
