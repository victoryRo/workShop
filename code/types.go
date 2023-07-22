package code

import (
	"fmt"
	"runtime"
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

// IntegerMemory ...
// Este código crea una colección de int o int8.
// Luego agrega 10 millones de valores a la colección.
// Una vez hecho esto, usa el paquete de tiempo de ejecución
// para darnos una lectura de cuánta memoria en montón se está usando.
// Convertimos esa lectura a MB y luego la imprimimos:
func IntegerMemory() {
	// var list []int
	var list []int8

	// llenamos el slice con 10 millones de valores
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}

	// struct Memoria Stadisticas
	var m runtime.MemStats
	// lee estadísticas de memoria (ReadMemStats)
	// rellena m con estadísticas del asignador de memoria
	runtime.ReadMemStats(&m)
	// m.TotalAlloc son bytes acumulados asignados para objetos de montón
	fmt.Printf("TotalAllocated (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}

// FloatPrecision ...
func FloatPrecision() {
	var a int = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println(a / 3) // 33
	fmt.Println(b / 3) // 33.333332
	fmt.Println(c / 3) // 33.333333333333336

	// Veamos qué sucede si intentamos que nuestro número vuelva a 100 multiplicándolo por 3
	fmt.Println(a / 3 * 3) // 99
	fmt.Println(b / 3 * 3) // 100
	fmt.Println(c / 3 * 3) // 100
	// A primera vista, las matemáticas de punto flotante pueden parecer simples,
	// pero se complican rápidamente. Al definir sus variables de punto flotante,
	// normalmente float64 debería ser su primera opción a menos que necesite ser más eficiente con la memoria.
}
