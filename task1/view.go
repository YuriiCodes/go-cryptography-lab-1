package task1

import "fmt"

func View() {
	n := 7 // Замініть це число на те, яке вам потрібно обчислити
	euler := EulerPhi(n)
	mob := Mobius(n)

	fmt.Printf("Функція Ейлера для %d: %d\n", n, euler)
	fmt.Printf("Функція Мьобіуса для %d: %d\n", n, mob)

	numbers := []int{12, 18, 24} // Замініть цей набір чисел на той, який вам потрібно обчислити НСК
	lcmResult := FindLCM(numbers)

	fmt.Printf("НСК для набору чисел %v: %d\n", numbers, lcmResult)
}
