package roman

import (
	"fmt"
	"testing"
)

// TestConvertToRoman verifies that ConvertToRoman correctly converts integers to Roman numerals.
func TestConvertToRoman(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		// Basic cases
		{2023, "MMXXIII"},
		{15, "XV"},
		{1, "I"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		// Edge cases
		{0, "N"},
		{-1, ""},
		{-100, ""},
		// Additional cases
		{3999, "MMMCMXCIX"}, // Maximum valid Roman numeral
		{1994, "MCMXCIV"},   // Common year
		{49, "XLIX"},        // Combination of subtractive and additive
		{99, "XCIX"},        // Double subtractive
		{2025, "MMXXV"},     // Current year
		{1001, "MI"},        // Simple addition
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d to %s", tt.n, tt.want), func(t *testing.T) {
			got := ConvertToRoman(tt.n)
			if got != tt.want {
				t.Errorf("ConvertToRoman(%d) = %s; want: %s", tt.n, got, tt.want)
			}
		})
	}
}
