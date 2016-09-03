package numberinfo

import "math/big"

// Number is an interface to number helpers
type Number interface {
	Int64Number() (*Int64Number, error)
	Float64Number() (*Float64Number, error)

	BigFactorial() (*big.Int, error)
	IsPrime() bool
	String() string
	Sqrt() float64
}
