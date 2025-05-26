package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	INR = "INR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR, CAD:
		return true
	}
	return false
}
