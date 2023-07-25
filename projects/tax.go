package projects

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// SalesTaxCalculator ...
func SalesTaxCalculator() {
	taxTotal := .0

	// Torta
	taxTotal += salesTax(.99, .075)

	// Leche
	taxTotal += salesTax(2.75, .015)

	// Mantequilla
	taxTotal += salesTax(.87, .02)

	// tabla
	valueTable()
	// Total
	fmt.Println("Sales tax total: ", taxTotal)
}

func salesTax(cost, taxRate float64) float64 {
	return cost * taxRate
}

func valueTable() {
	// Alinea texto con tabwriter
	w := tabwriter.NewWriter(os.Stdout, 9, 0, 2, ' ', 1)

	fmt.Fprintln(w, "Item\tCost\tSales Tax Rate\t")
	fmt.Fprintln(w, "Torta\t0.99\t75 %\t")
	fmt.Fprintln(w, "Leche\t2.75\t15 %\t")
	fmt.Fprintln(w, "Mantequilla\t.87\t2 %\t")

	// ejecuta el code anterior
	w.Flush()
	fmt.Println()
}
