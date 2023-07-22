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

// -------------------------------------------------------------------------------------

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

// -------------------------------------------------------------------------------------

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

// -------------------------------------------------------------------------------------

// Overflow ...
// Desbordamiento y envolvente
// Cuando intenta inicializar un número con un valor que es demasiado grande
// para el tipo que estamos usando, obtiene un error de desbordamiento
func Overflow() {
	// el numero mas alto aceptado por int8 es (127)
	// var a int8 = 128
	// fmt.Println(a)
}

// NumberWraparound ...
// Desencadenamiento de números envolventes
func NumberWraparound() {
	var a int8
	var b uint8
	a = 125
	b = 253

	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}

	// En este ejercicio, vimos que, para los enteros con signo,
	// terminaría con un número negativo y para los enteros sin signo, se vuelve a 0.
	// Siempre debe considerar el número máximo posible para su variable
	// y asegurarse de tener el tipo adecuado para respaldar ese número.
}

// -------------------------------------------------------------------------------------

// Numeros Grandes
