package roman

// indices holds the key values for Roman numeral conversion in ascending order.
var indices = [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

// ConvertToRoman converts a non-negative integer to its Roman numeral representation.
// Returns an empty string for negative inputs and "N" for zero.
func ConvertToRoman(n int) string {
	if n < 0 {
		return ""
	}
	switch n {
	case 0:
		return "N"
	case 1:
		return "I"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 9:
		return "IX"
	case 10:
		return "X"
	case 40:
		return "XL"
	case 50:
		return "L"
	case 90:
		return "XC"
	case 100:
		return "C"
	case 400:
		return "CD"
	case 500:
		return "D"
	case 900:
		return "CM"
	case 1000:
		return "M"
	default:
		// Find the largest Roman numeral value less than or equal to n
		var d int = indices[0]
		for i := len(indices) - 1; i > 0; i-- {
			if indices[i] <= n {
				d = indices[i]
				break
			}
		}
		// Convert the found value and recursively convert the remainder
		var num string = ConvertToRoman(d)
		var r int = n - d
		if r > 0 {
			num += ConvertToRoman(r)
		}
		return num
	}
}
