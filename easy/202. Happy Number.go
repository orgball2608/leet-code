package easy

func isHappy(n int) bool {
	slow := n
	fast := sumDigits(n)
	for fast != 0 && slow != fast {
		slow = sumDigits(slow)
		fast = sumDigits(sumDigits(fast))
	}
	if fast == 1 {
		return true
	}

	return false
}

func sumDigits(number int) int {
	sum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		sum += digit * digit
	}
	return sum
}
