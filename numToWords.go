package numToWords

func Convert(num uint64) string {
	switch {
	case num == 0:
		return "zero"
	case num == 1:
		return "one"
	case num == 2:
		return "two"
	case num == 3:
		return "three"
	case num == 4:
		return "four"
	case num == 5:
		return "five"
	case num == 6:
		return "six"
	case num == 7:
		return "seven"
	case num == 8:
		return "eight"
	case num == 9:
		return "nine"

	case num == 10:
		return "ten"
	case num == 11:
		return "eleven"
	case num == 12:
		return "twelve"
	case num == 13:
		return "thirteen"
	case num == 15:
		return "fifteen"
	case num == 18:
		return "eighteen"

	case num >= 14 && num <= 19:
		return Convert(num%10) + "teen"

	case num == 20:
		return "twenty"
	case num == 30:
		return "thirty"
	case num == 40:
		return "forty"
	case num == 50:
		return "fifty"
	case num == 80:
		return "eighty"

	case num >= 60 && num <= 90 && num % 10 == 0:
		return Convert(num/10) + "ty"

	case num >= 21 && num <= 99:
		return Convert(num-num%10) + "-" + Convert(num%10)

	case num >= 100 && num < 1000:
		return convertHundredAndAbove(num, 100, "hundred")
	case num >= 1000 && num < 1_000_000:
		return convertHundredAndAbove(num, 1000, "thousand")
	case num >= 1_000_000 && num < 1_000_000_000:
		return convertHundredAndAbove(num, 1_000_000, "million")
	case num >= 1_000_000_000 && num < 1_000_000_000_000:
		return convertHundredAndAbove(num, 1_000_000_000, "billion")
	default:
		return "Not Implemented"
	}
}

func convertHundredAndAbove(num, divisor uint64, placeValue string) string {
	if num % divisor == 0 {
		return Convert(num/divisor) + " " + placeValue
	} else {
		return Convert(num-(num%divisor)) + " " + Convert(num%divisor)
	}
}
