package task2

import (
	"math/big"
)

// Функція для знаходження оберненого числа a за модулем m
func modInverse(a, m *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, m)
}

// Функція для знаходження x за китайською теоремою про лишки
func SolveCRT(congruences [][2]*big.Int) *big.Int {
	n := len(congruences)
	M := new(big.Int)
	result := new(big.Int)

	// Знаходимо загальний модуль M, який є добутком усіх m[i]
	M.SetUint64(1)
	for i := 0; i < n; i++ {
		M.Mul(M, congruences[i][1])
	}

	// Знаходимо x за кожним конгруентним рівнянням
	for i := 0; i < n; i++ {
		Mi := new(big.Int).Div(M, congruences[i][1])
		MiInv := modInverse(Mi, congruences[i][1])
		result.Add(result, new(big.Int).Mul(congruences[i][0], new(big.Int).Mul(Mi, MiInv)))
	}

	// Повертаємо x за модулем M
	return new(big.Int).Mod(result, M)
}
