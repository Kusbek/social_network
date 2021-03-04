package entity

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ID = int

const layoutISO = "2006-01-02"

func StringToTime(date string) time.Time {
	d, err := time.Parse(layoutISO, date)
	if err != nil {
		fmt.Println(err)
	}
	return d
}

func TimeToString(date time.Time) string {
	return date.Format(layoutISO)
}

func encrypt(str string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
