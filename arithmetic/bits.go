package arithmetic

import "github.com/swarit-pandey/go-fast/common"

// BitsRequired returns you the number of bits required to
// represent a number (only for types supported by std lib Go) in binary system
// runs in O(log n) where n is the number given
func BitsRequired[T common.Number](a T) int {
	if a == 0 {
		return 1
	}

	if a < 0 {
		a = -a
	}

	r := 0
	for a > 0 {
		r++
		a >>= 1
	}

	return r
}
