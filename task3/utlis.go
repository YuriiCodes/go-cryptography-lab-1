package task3

import (
	"math/big"
)

// Ця функція обчислює символ Якобі (a/b).
func jacobiSymbol(a, b *big.Int) int {
	if b.Cmp(big.NewInt(1)) == 0 {
		return 1
	}
	if a.Cmp(big.NewInt(0)) == 0 {
		return 0
	}
	if a.Cmp(big.NewInt(2)) == 0 || a.Cmp(big.NewInt(3)) == 0 {
		r := new(big.Int).Mod(b, big.NewInt(8))
		if r.Cmp(big.NewInt(1)) == 0 || r.Cmp(big.NewInt(7)) == 0 {
			return jacobiSymbol(b, a)
		} else if r.Cmp(big.NewInt(3)) == 0 || r.Cmp(big.NewInt(5)) == 0 {
			return -jacobiSymbol(b, a)
		}
	} else if a.Cmp(big.NewInt(1)) == 0 {
		return 1
	}
	if a.Cmp(big.NewInt(0).Set(b)) == 1 {
		a.Mod(a, b)
	}
	if a.Cmp(big.NewInt(2)) == 0 || a.Cmp(big.NewInt(3)) == 0 {
		r := new(big.Int).Mod(b, big.NewInt(8))
		if r.Cmp(big.NewInt(1)) == 0 || r.Cmp(big.NewInt(7)) == 0 {
			return jacobiSymbol(b, a)
		} else if r.Cmp(big.NewInt(3)) == 0 || r.Cmp(big.NewInt(5)) == 0 {
			return -jacobiSymbol(b, a)
		}
	}
	if a.Bit(0) == 0 {
		return jacobiSymbol(big.NewInt(2), b) * jacobiSymbol(a.Div(a, big.NewInt(2)), b)
	}
	if a.Cmp(big.NewInt(1)) == 0 {
		return 1
	}
	return jacobiSymbol(b.Mod(b, a), a) * jacobiSymbol(a, b)
}

func Legendre(a *big.Int, p *big.Int) *big.Int {
	if a.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	e := new(big.Int).Sub(p, big.NewInt(1))
	e.Div(e, big.NewInt(2))

	r := new(big.Int).Exp(a, e, p)
	if r.Cmp(big.NewInt(1)) > 0 {
		return new(big.Int).Sub(r, p)
	}
	return r
}
