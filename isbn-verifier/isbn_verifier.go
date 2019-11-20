package isbn

func IsValidISBN(isbn string) bool {
	if len(isbn) < 10 {
		return false
	}

	var sum int
	factor := 10
	for i, r := range isbn {
		if r == 'X' && i == len(isbn)-1 {
			sum += 10
		}
		if '0' <= r && r <= '9' {
			sum += int(r-'0') * factor
			factor--
		}
		if factor < 0 {
			return false
		}
	}

	return sum%11 == 0
}
