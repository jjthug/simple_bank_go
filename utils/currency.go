package utils

const (
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR:
		return true
	default:
		return false
	}
}
