package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	sum := "0"
	var s string
	for i := len(num2) - 1; i >= 0; i-- {
		s = multiplyWith1DigitNumber(num1, num2[i])
		for j := len(num2) - i - 1; j > 0; j-- {
			s = s + "0"
		}
		sum = plus(sum, s)
	}
	return sum
}

func multiplyWith1DigitNumber(num1 string, num2 byte) string {
	sum := ""
	x := 0
	var s int
	for i := len(num1) - 1; i >= 0; i-- {
		s = int(num1[i]-'0') * int(num2-'0')
		sum = string((s+x)%10+'0') + sum
		x = (s + x) / 10
	}
	if x != 0 {
		return string(x+'0') + sum
	}
	return sum
}

func plus(num1 string, num2 string) string {
	for len(num1) < len(num2) {
		num1 = "0" + num1
	}
	for len(num1) > len(num2) {
		num2 = "0" + num2
	}
	sum := ""
	x := 0
	var s int
	for i := len(num1) - 1; i >= 0; i-- {
		s = int(num1[i]-'0') + int(num2[i]-'0')
		sum = string((s+x)%10+'0') + sum
		x = (s + x) / 10
	}
	if x != 0 {
		return string(x+'0') + sum
	}
	return sum
}
