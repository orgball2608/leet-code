package easy

func isHappy(n int) bool {
	slow := n
	fast := sumDigits(n)

	for fast != 1 && fast != slow {
		slow = sumDigits(slow)
		fast = sumDigits(sumDigits(fast))
	}

	if fast == 1 {
		return true
	}

	return false
}

func pow(digit, power int) int {
	r := 1

	for i := 0; i < power; i++ {
		r *= digit
	}

	return r
}

func sumDigits(number int) int {
	sum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		sum += pow(digit, 2)
	}

	return sum
}
