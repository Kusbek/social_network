package entity_test

import (
	"fmt"
	"testing"

	"git.01.alem.school/Kusbek/social-network/entity"
)

func TestStringToTime(t *testing.T) {
	date1 := "1994-09-18"
	date2 := "1994-09-18 00:00:00+00:00 "

	res1 := entity.StringToTime(date1)
	res2 := entity.StringToTime(date2)

	fmt.Println(res1)
	fmt.Println(res2)
}
