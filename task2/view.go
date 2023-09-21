package task2

import (
	"fmt"
	"math/big"
)

func View() {

	eq1x := big.NewInt(2)
	eq1m := big.NewInt(3)

	eq2x := big.NewInt(3)
	eq2m := big.NewInt(5)

	eq3x := big.NewInt(2)
	eq3m := big.NewInt(7)

	fmt.Println("Система:")
	fmt.Printf("x = %v (mod %v)\n", eq1x, eq1m)
	fmt.Printf("x = %v (mod %v)\n", eq2x, eq2m)
	fmt.Printf("x = %v (mod %v)\n", eq3x, eq3m)
	congruences := [][2]*big.Int{
		{eq1x, eq1m},
		{eq2x, eq2m},
		{eq3x, eq3m},
	}

	x := SolveCRT(congruences)
	fmt.Printf("Розв'язок системи: x = %v\n", x)

}
