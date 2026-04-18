package internal
import "errors"
// Add додає два числа
func Add(a, b int) int {
	return a + b
}
// Divide ділить числа, повертає помилку при діленні на нуль
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}