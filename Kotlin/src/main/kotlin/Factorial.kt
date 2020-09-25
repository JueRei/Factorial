package uk.org.winder.maths.factorial

import java.lang.IllegalArgumentException
import java.math.BigInteger

val zero: BigInteger = BigInteger.ZERO
val one: BigInteger = BigInteger.ONE
val two = 2.bigint

fun validate(x: BigInteger): Unit {
	if (x < zero) { throw IllegalArgumentException("Factorial not defined for negative integers.") }
}

fun iterative(x: BigInteger): BigInteger {
	validate(x)
	var total = one
	for (i in two .. x) { total *= i }
	return total
}
fun iterative(x: Int): BigInteger = iterative(x.bigint)
fun iterative(x: Long): BigInteger = iterative(x.bigint)

fun reductive(x: BigInteger): BigInteger {
	validate(x)
	if (x < two) { return one }
	return (one .. x).reduce{t, i -> t * i}
}
fun reductive(x: Int): BigInteger = reductive(x.bigint)
fun reductive(x: Long): BigInteger = reductive(x.bigint)

fun foldive(x: BigInteger): BigInteger {
	validate(x)
	return (two .. x).fold(one){t, i -> t * i}
}
fun foldive(x: Int): BigInteger = foldive(x.bigint)
fun foldive(x: Long): BigInteger = foldive(x.bigint)

fun naiveRecursive(x: BigInteger): BigInteger {
	validate(x)
	if (x < two) { return one }
	return x * naiveRecursive(x - one)
}
fun naiveRecursive(x: Int): BigInteger = naiveRecursive(x.bigint)
fun naiveRecursive(x: Long): BigInteger = naiveRecursive(x.bigint)

fun tailRecursive(x: BigInteger): BigInteger {
	validate(x)
	tailrec fun iterate(n: BigInteger, t: BigInteger = one): BigInteger {
		if (n < two) { return t }
		return iterate(n - one, n * t)
	}
	return iterate(x)
}
fun tailRecursive(x: Int): BigInteger = tailRecursive(x.bigint)
fun tailRecursive(x: Long): BigInteger = tailRecursive(x.bigint)
