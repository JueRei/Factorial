// Package factorial provides a number of implementation of factorial, the mathematical function defined by
// the recurrence relation:
//
//  f_0 = 1
//  f_n = n . f_{n - 1}
//
// The variants show different algorithms. For each algorithm there are two implementations, one for a uint,
// one for a big.Int. The return value in all cases is a big.Int.
package factorial

import "math/big"

var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)
var bigTwo = big.NewInt(2)

// Iterative_uint is an imperative, loop-based implementation with the parameter being an uint value.
func Iterative_uint(n uint) (total *big.Int) {
	total = big.NewInt(1)
	for i := uint(2); i <= n; i++ {
		total.Mul(total, big.NewInt(int64(i)))
	}
	return
}

// Iterative_bigInt is an imperative, loop-based implementation with the parameter being a big.Int value.
func Iterative_bigInt(n *big.Int) (total *big.Int) {
	total = big.NewInt(1)
	for i := new(big.Int).Set(n); i.Cmp(bigOne) > 0; i.Sub(i, bigOne) {
		total.Mul(total, i)
	}
	return
}

// Recursive_uint is the naïve recursive implementation, i.e. it is not tail recursive, with a uint as parameter.
func Recursive_uint(n uint) *big.Int {
	return Recursive_bigInt(big.NewInt(int64(n)))
}

// Recursive_uint is the naïve recursive implementation, i.e. it is not tail recursive, with a big.Int as parameter.
func Recursive_bigInt(n *big.Int) (value *big.Int) {
	value = big.NewInt(1)
	if n.Cmp(bigOne) > 0 {
		value.Mul(n, Recursive_bigInt(new(big.Int).Sub(n, bigOne)))
	}
	return
}

// TailRecursive_uint is the tail recursive implementation with a uint as parameter. There is though no tail
// call optimization in Go so this is not equivalent to the imperative version.
func TailRecursive_uint(n uint) *big.Int {
	return TailRecursive_bigInt(big.NewInt(int64(n)))
}

func tailRecursive_bigInt_iterate(n, result *big.Int) *big.Int {
	if n.Cmp(bigTwo) < 0 {
		return result
	}
	return tailRecursive_bigInt_iterate(new(big.Int).Sub(n, bigOne), new(big.Int).Mul(result, n))
}

// TailRecursive_bigInt is the tail recursive implementation with a big.Int as parameter. There is though no
// tail call optimization in Go so this is not equivalent to the imperative version.
func TailRecursive_bigInt(n *big.Int) (value *big.Int) {
	value = big.NewInt(1)
	if n.Cmp(bigOne) > 0 {
		return tailRecursive_bigInt_iterate(new(big.Int).Set(n), value)
	}
	return
}

var algorithms_uint = []func(uint) *big.Int{
	Iterative_uint,
	Recursive_uint,
	TailRecursive_uint,
}

var algorithms_bigInt = []func(*big.Int) *big.Int{
	Iterative_bigInt,
	Recursive_bigInt,
	TailRecursive_bigInt,
}
