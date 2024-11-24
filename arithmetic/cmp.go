package arithmetic

import (
	"github.com/swarit-pandey/go-fast/common"
)

// Max returns a greater number between the two simply
// its a single instruction on x86 systems (CMOV on x86)
func Max[T common.Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns a lesser number between the two simply
// its a single instruction on x86 systems (CMOV on x86)
func Min[T common.Number](a, b T) T {
	if a > b {
		return b
	}
	return a
}

// MaxOf returns the max integer out of a slice, runs in O(n)
// works for all the std lib numbers supported
func MaxOf[T common.Number](a ...T) T {
	if len(a) == 0 {
		panic("input is an empty slice")
	}

	r := a[0]
	for _, v := range a {
		if v > r {
			r = v
		}
	}

	return T(r)
}

func MinOf[T common.Number](a ...T) T {
	if len(a) == 0 {
		panic("input is an empty slice")
	}

	r := a[0]
	for _, v := range a {
		if v < r {
			r = v
		}
	}

	return T(r)
}

// MinMaxOf gets the min and max from the given array
// runs in O(n) but makes ~1.5n comparisons, instead of 2n
func MinMaxOf[T common.Number](a []T) (min, max T) {
	n := len(a)
	if len(a) < 2 {
		panic("input slice must have 2 elements")
	}

	i := 0
	if (n & 1) == 0 {
		if a[0] > a[1] {
			max = a[0]
			min = a[1]
		} else {
			max = a[1]
			min = a[0]
		}
		i = 2
	} else {
		max = a[0]
		min = a[0]
		i = 1
	}

	if n <= 2 {
		return min, max
	}

	for ; i < n-1; i += 2 {
		if a[i] > a[i+1] {
			max = Max(a[i], max)
			min = Min(a[i+1], min)
		} else {
			max = Max(a[i+1], max)
			min = Min(a[i], min)
		}
	}

	return min, max
}

// FastMax is only for positive integers and uses bit mask
// to get the greater number, it avoids branching
func FastMax[T common.Positive64](a, b T) T {
	diff := a - b
	mask := diff >> 63
	return b + (diff & ^mask)
}

// Clamp constraints the value betwen min and max
func Clamp[T common.Number](val, min, max T) T {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// Abs returns absolute value
func Abs[T common.Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sign returns
// -1 if x < 0
// 0 if x == 0
// 1 if x > 0
func Sign[T common.Number](x T) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}
