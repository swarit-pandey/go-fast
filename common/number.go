package common

// Number is a an interface that represents all integers supported by std lib
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Positive64 is a constraint for positive 64-bit integers
type Positive64 interface {
	~int64 | ~uint64
}
