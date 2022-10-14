package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	logger "github.com/sirupsen/logrus"
)

func CheckForNumbers(str string) bool {
	numCheck := regexp.MustCompile(`^[0-9]+$`)
	return numCheck.MatchString(str)
}

func IsMPinValid(mpin string) (bool, error) {
	var (
		letterPresent      bool
		numberPresent      bool
		specialCharPresent bool
		passLen            int
		errorString        string
	)
	MPinLength, err := StringToNumber(os.Getenv("MPIN_LENGTH"))
	if err != nil {
		logger.Error("IsPasswordValid: Parse password length. Error: ", err)
		return false, err
	}

	for _, ch := range mpin {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsLetter(ch):
			letterPresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if letterPresent {
		appendError("MPin cannot conatin character")
		logger.Error("MPin cannot conatin character")
	}
	if !numberPresent {
		appendError("numeric character required")
		logger.Error("numeric character required")
	}
	if specialCharPresent {
		appendError("MPin cannot conatin special character")
		logger.Error("MPin cannot conatin special character")
	}
	if len(mpin) != MPinLength {
		appendError(fmt.Sprintf("MPin length must be between %d digits long", MPinLength))
	}
	if len(errorString) != 0 {
		return false, errors.New(errorString)
	}
	return true, nil
}
func StringToNumber(key string) (int, error) {
	nkey, _ := strconv.Atoi(key)
	return nkey, nil
}

func NumberToString(key int) string {
	skey := strconv.Itoa(key)
	return skey
}

func CheckMPin(mpin, rempin string) int {
	comp := strings.Compare(mpin, rempin)
	return comp
}

func GenOTP(phone string) (string, string, error) {
	max := 9999
	min := 1111
	otp := rand.Intn(max-min) + min
	fmt.Println("otp==============================>", otp)
	strOTP := NumberToString(otp)
	encOTP, err := Encrypt(strOTP, os.Getenv("OTP_ENC_KEY"))
	if err != nil {
		logger.Error("func_CreateUser: Error in encrypt password: ", err)
		return encOTP, strOTP, err
	}
	return encOTP, strOTP, err
}
