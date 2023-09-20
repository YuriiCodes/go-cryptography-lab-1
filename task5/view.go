package task5

import (
	"fmt"
	"math/big"
)

func View() {
	base := big.NewInt(3)     // Основа
	target := big.NewInt(10)  // Значення, для якого шукаємо дискретний логарифм
	modulus := big.NewInt(13) // Модуль поля

	result := BigStepLittleStep(base, target, modulus)
	if result != nil {
		fmt.Printf("Дискретний логарифм для %s за основою %s в полі %s: %s\n", target.String(), base.String(), modulus.String(), result.String())
	} else {
		fmt.Printf("Дискретний логарифм не знайдено.\n")
	}
}
