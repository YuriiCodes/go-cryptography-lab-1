package task3

import (
	"fmt"
	"math/big"
)

func View() {
	a := big.NewInt(5)
	b := big.NewInt(7)

	legendre := Legendre(a, b)
	fmt.Println(legendre)

	jacobi := big.NewInt(17)
	jacobi.Mod(jacobi, b)
	jacRes := JacobiSymbol(jacobi, b)
	fmt.Printf("Символ Якобі (%v/%v) = %v\n", jacobi, b, jacRes)
}
