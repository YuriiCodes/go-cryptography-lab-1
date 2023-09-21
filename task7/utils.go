package task7

import (
	"errors"
	"math/big"
	"math/rand"
	"time"
)

// MillerRabin performs the Miller-Rabin primality test with k iterations, true if n is probably prime, false if n is composite
func MillerRabin(n *big.Int, k int) (bool, error) {
	if n.Cmp(big.NewInt(2)) <= 0 {
		return false, errors.New("n must be greater than 2")
	}

	// n - 1 = 2^s * d, where d is odd
	s := 0
	d := new(big.Int).Sub(n, big.NewInt(1))
	for d.Bit(0) == 0 {
		s++
		d.Rsh(d, 1)
	}

	// Random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < k; i++ {
		// Pick a random a in the range [2, n-1]
		a := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(n, big.NewInt(2)))
		a.Add(a, big.NewInt(2))

		x := new(big.Int).Exp(a, d, n)

		if x.Cmp(big.NewInt(1)) == 0 || x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
			continue // Pass the next iteration
		}

		for j := 0; j < s-1; j++ {
			x.Exp(x, big.NewInt(2), n)
			if x.Cmp(big.NewInt(1)) == 0 {
				return false, nil
			}
			if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
				break // Pass the next iteration
			}
		}

		if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
			return false, nil
		}
	}

	return true, nil
}
