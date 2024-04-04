package validator

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Implements Luhn Algorithm (modulus 10 algorithm)

func Validate(card string) (bool, error) {

	card, err := checkCardNumber(card)

	if err != nil {
		log.Fatalln(err)
		return false, err
	}

	checkDigit, err := strconv.Atoi(card[len(card)-1:])

	if err != nil {
		log.Fatalln(err)
		return false, err
	}

	calculatedCheckDigit, err := calculateCheckDigit(card)

	if err != nil {
		log.Fatalln(err)
		return false, err
	}

	return checkDigit == calculatedCheckDigit, nil
}

func checkCardNumber(card string) (string, error) {
	pattern := "^(\\d{4}\\s\\d{4}\\s\\d{4}\\s\\d{4}|\\d{16})$"
	ok, err := regexp.Match(pattern, []byte(card))

	if err != nil {
		log.Fatalln(err)
		return "", errors.New("unknown error")
	}

	if !ok {
		return "", errors.New("invald card format")
	}

	return strings.ReplaceAll(card, " ", ""), nil
}

func calculateCheckDigit(card string) (int, error) {
	runningSum := 0
	for i := (len(card) - 2); i >= 0; i-- {
		digit, err := strconv.Atoi(card[i : i+1])
		// Not expected error, but just in case
		if err != nil {
			log.Fatalln(err)
			return 0, err
		}

		if i%2 == 0 {
			digit *= 2
			for digit > 0 {
				runningSum += digit % 10
				digit /= 10
			}
		} else {
			runningSum += digit
		}
	}

	return 10 - (runningSum % 10), nil
}
