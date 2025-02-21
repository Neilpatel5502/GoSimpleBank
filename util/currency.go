package util

const (
	USD = "USD"
	AUD = "AUD"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, AUD, CAD:
		return true
	}
	return false
}
