package utils

import "fmt"

func FormatAsDollars(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}
