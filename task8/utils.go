package task8

import (
	"math/big"
)

func GenerateKeys(p, q *big.Int) (n, e, d *big.Int) {
	// Compute n = p * q
	n = new(big.Int).Mul(p, q)

	// Compute Euler's totient function: phi(n) = (p-1)*(q-1)
	phiN := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	// Choose public exponent e, such that 1 < e < phi(n) and gcd(e, phi(n)) = 1
	e = big.NewInt(17) // a common choice

	// Compute private exponent d, the modular multiplicative inverse of e (mod phi(n))
	d = new(big.Int).ModInverse(e, phiN)

	return n, e, d
}

func Encrypt(m, e, n *big.Int) *big.Int {
	// Encryption: c ≡ m^e (mod n)
	return new(big.Int).Exp(m, e, n)
}

func Decrypt(c, d, n *big.Int) *big.Int {
	// Decryption: m ≡ c^d (mod n)
	return new(big.Int).Exp(c, d, n)
}
