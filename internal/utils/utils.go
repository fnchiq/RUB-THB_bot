package utils

import (
	"fmt"
)

func ExchangeRateTHB(n1 []float64, d float64) string {

	rateSb := n1[0] / d
	rateTin := n1[1] / d

	return fmt.Sprintf("Курс Cбербанк:\n%.4f\n\nКурс Тинькофф:\n%.4f\n", rateSb, rateTin)
}

func ExchangeRateKZT(n1 []float64, d float64) string {

	rateSb := d / n1[0]
	rateTin := d / n1[1]

	return fmt.Sprintf("Курс Cбербанк:\n%.4f\n\nКурс Тинькофф:\n%.4f\n", rateSb, rateTin)
}
