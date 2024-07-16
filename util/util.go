package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/leekchan/accounting"
)

func FormatCurrency(amount int) string {
	ac := accounting.Accounting{
	    Symbol:    "Rp ",
	    Precision: 0,
	    Thousand:  ".",
	    Decimal:   ",",
	}
	return ac.FormatMoney(amount)
}

func RandomWithProbability(values []int, probabilities []int) (int, error) {
	if len(values) != len(probabilities) {
		return 0, fmt.Errorf("length of values and probabilities must match")
	}

	rand.Seed(time.Now().UnixNano())

	totalProbability := 0
	for _, probability := range probabilities {
		totalProbability += probability
	}

	r := rand.Intn(totalProbability)

	accumulatedProbability := 0
	for i, probability := range probabilities {
		accumulatedProbability += probability
		if r < accumulatedProbability {
			return values[i], nil
		}
	}

	return 0, fmt.Errorf("unexpected error in randomWithProbability")
}

func RandomInRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}

func RandomWithDigitRange(minDigit, maxDigit int) int {
	rand.Seed(time.Now().UnixNano())

	min := pow(10, minDigit-1)
	max := pow(10, maxDigit) - 1

	return rand.Intn(max-min+1) + min
}

func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}