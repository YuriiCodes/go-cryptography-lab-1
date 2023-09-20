package task4

import (
	"fmt"
	"math/big"
)

func View() {
	n := new(big.Int)
	n.SetString("1234567890123456789012345678901234567890", 10)

	factor := PollardRho(n)
	fmt.Printf("Фактор n: %s\n", factor.String())
}
