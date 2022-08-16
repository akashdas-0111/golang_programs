package math
func Add(num1 int, num2 int) int {
	return num1 + num2
}
func Mul(num1 int, num2 int) int {
	return num1 * num2
}
func Sub(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}
func ReverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}
func Wordc(s string) int {
	return len(s)
}
