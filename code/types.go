package code

import (
	"fmt"
	"math"
	"math/big"
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

// BigNumbers ...
// sum of very large numbers
func BigNumbers() {
	// sumamos un int grande
	// ser convierte en un numero negativo
	intA := math.MaxInt64
	intA = intA + 1

	// sumamos un int grande
	// retorna el valor correcto de la suma
	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int     :", intA)
	fmt.Println("Big Int : ", bigA.String())
	// MaxInt64:  9223372036854775807
	// Int     : -9223372036854775808
	// Big Int :  9223372036854775808

	// Si tiene una situación en la que tiene un número cuyo valor es más alto de lo que Go puede administrar,
	// entonces el paquete big de la biblioteca estándar es lo que necesita.
}

// -------------------------------------------------------------------------------------

// Byte
// El tipo byte en Go es solo un alias para uint8, que es un número que tiene 8 bits de almacenamiento.
// En realidad, byte es un tipo significativo y lo verás en muchos lugares.
// Un bit es un único valor binario, un único interruptor de encendido/apagado.
// La agrupación de bits en grupos de 8 era un estándar común en la computación temprana
// y se convirtió en una forma casi universal de codificar datos.
// 8 bits tienen 256 combinaciones posibles de "apagado" y "encendido",
// uint8 tiene 256 valores enteros posibles de 0 a 255.
// Todas las combinaciones de encendido y apagado se representan con este tipo.

// Verá que byte se usa al leer y escribir datos hacia y desde una conexión de red y al leer y escribir datos en archivos.

// -------------------------------------------------------------------------------------

// Runas ...
func Runas() {
	username := "Sir_King_Uber"
	s := 'S'
	i := 'i'
	r := 'r'
	o := '_'

	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}

	fmt.Println()
	fmt.Printf("S type %T value %v\n", s, s)
	fmt.Printf("S type %T value %v\n", i, i)
	fmt.Printf("S type %T value %v\n", r, r)
	fmt.Printf("S type %T value %v\n", o, o)

	for idx, v := range username {
		fmt.Println(idx, string(v))
	}

	// Cuando el texto está en el tipo string, para permitir esta variabilidad,
	// Go almacena todas las cadenas como una colección de byte.
	// Para poder realizar operaciones de forma segura con texto de cualquier tipo de codificación,
	// de un solo byte o de varios bytes, debe convertirse de una colección de bytes
	// a una colección de runas.

	// Cuando trabaje con cadenas, asegúrese de verificar con el paquete strings primero.
	// Está lleno de herramientas útiles que ya pueden hacer lo que necesita.
}
