package task1

import (
	"strconv"
	"strings"
)

// Функція для обчислення функції Ейлера для числа n
func EulerPhi(n int) int {
	result := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}

// Функція для обчислення функції Мьобіуса для числа n
func Mobius(n int) int {
	if n == 1 {
		return 1
	}
	result := 1
	count := 0
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			n /= i
			count++
			if n%i == 0 {
				return 0
			}
		}
		if count > 1 {
			return 0
		}
		if count == 1 {
			result = -result
		}
	}
	return result
}

// Функція для обчислення найменшого спільного кратного для двох чисел a і b
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Функція для обчислення найбільшого спільного дільника для двох чисел a і b
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Функція для знаходження НСК для набору чисел
func FindLCM(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

// a func that returns an error and []int:
// e.g. input is  "1, 2, 3" => []int{1, 2, 3}, nil
// e.g. input is  "4,5,6" => []int{4, 5, 6}, nil
func ParseIntArray(input string) ([]int, error) {
	var result []int
	for _, s := range strings.Split(input, ",") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil

}
