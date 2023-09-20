package task5

import "math/big"

func BigStepLittleStep(base, target, modulus *big.Int) *big.Int {
	// Визначаємо межу для великого кроку
	m := new(big.Int).Sqrt(modulus)
	m.Add(m, big.NewInt(1))

	// Знаходимо a^m mod p
	a := new(big.Int).Set(base)
	a.Exp(a, m, modulus)

	// Обчислюємо таблицю малих кроків
	smallStepTable := make(map[string]*big.Int)
	u := new(big.Int)

	for j := big.NewInt(0); j.Cmp(m) < 0; j.Add(j, big.NewInt(1)) {
		u.Exp(base, j, modulus)
		u.Mod(u, modulus)
		smallStepTable[u.String()] = j
	}

	// Шукаємо великий крок
	v := new(big.Int).Set(target)
	for i := big.NewInt(0); i.Cmp(m) < 0; i.Add(i, big.NewInt(1)) {
		// Обчислюємо v * a^(-i) mod p
		v.Mul(v, a)
		v.Mod(v, modulus)

		// Перевіряємо, чи є це значення в таблиці малих кроків
		if j, ok := smallStepTable[v.String()]; ok {
			// Знайшли великий крок
			result := new(big.Int).Mul(i, m)
			result.Add(result, j)
			return result
		}
	}

	return nil // Дискретний логарифм не знайдено
}
