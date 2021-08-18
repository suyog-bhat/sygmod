package mulpack

import "errors"

func Mulnums(a int, b int) int {
	return a * b
}
func Divide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("cant divide by zero")
	} else {
		return float64(a) / float64(b), nil
	}
}
