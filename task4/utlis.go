package task4

import (
	"math/big"
)

func PollardRho(n *big.Int) *big.Int {
	// Ініціалізуємо змінні x, y і d
	x := big.NewInt(2)
	y := big.NewInt(2)
	d := big.NewInt(1)

	// Функція для обчислення f(x) = x^2 + 1 (модуль n)
	f := func(x *big.Int) *big.Int {
		xSquared := new(big.Int).Set(x)
		xSquared.Mul(x, x)
		xSquared.Add(xSquared, big.NewInt(1))
		return xSquared.Mod(xSquared, n)
	}

	// Пошук спільного дільника
	for d.Cmp(big.NewInt(1)) == 0 {
		x.Set(f(x))
		y.Set(f(f(y)))

		// Використовуємо великий спільний дільник для n та |x - y|
		temp := new(big.Int).Set(x)
		temp.Sub(temp, y)
		d.GCD(nil, nil, temp, n)
	}

	return d
}
