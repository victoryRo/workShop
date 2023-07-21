package code

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// IfStatement ...
func IfStatement() {

	for i := 0; i < 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// SwitchExpression ...
func SwitchExpression() {

	// podemos usar operadores de comparacion
	// dentro del case
	switch borne := time.Sunday; {
	case borne == time.Friday || borne == time.Sunday:
		fmt.Printf("her borne in a weekend %v\n", borne)
	default:
		fmt.Println("Born some other day")
	}
}

// MapSong ...
func MapSong() {

	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	quantity := []int{}
	// ordena un slice de int de mayor a menor
	for _, times := range words {
		quantity = append(quantity, times)
		sort.Slice(quantity, func(i, j int) bool {
			return quantity[i] > quantity[j]
		})
	}

	// fmt.Printf("slice of quantities: %#v\n", quantity)

	// Recorre el map words y compara los valores con el valor mas alto dado por quantity[0]
	for w, t := range words {
		if words[w] == quantity[0] {
			fmt.Printf("The most popular word is %s\n", w)
			fmt.Printf("With a count of: %d\n", t)
		}
	}
}

// RandNumber ...
func RandNumber() {

	// for infinito
	for {
		// va dentro del for e itera un num diff cada vez
		num := rand.Intn(8)

		if num%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if num%2 == 0 {
			fmt.Println("Stop")
			break
		}

		// la salida va dentro retorna un valor x iteracion
		fmt.Println(num)
	}
}

// SortSlice ...
func SortSlice() {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Printf("Slice normal %v\n", nums)

	// toma el slice y lo ordena
	// modifica el slice dado
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Printf("Slice after sort %v\n", nums)
}
