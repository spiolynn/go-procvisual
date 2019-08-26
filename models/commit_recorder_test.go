package models

import (
	"fmt"
	"testing"
)

func TestGetCommitRecoderByChangeNum(t *testing.T) {

	beforeTest()
	cr, err := GetCommitRecoderByChangeNum(14835)
	if err == nil {
		fmt.Printf("%+v", *cr)
	} else {
		fmt.Printf("%v", err)
	}

}

func TestGetCommitNumbyTime(t *testing.T) {
	beforeTest()
	cr, err := GetCommitNumbyTime("PJ1901400","BK2C/pecf/pecf-prod-app","dev_core","2019-08-01","2019-08-30")
	if err == nil {
		fmt.Printf("%+v", cr)
	} else {
		fmt.Printf("%v", err)
	}
	}


func TestGetCommiShared(t *testing.T) {
	beforeTest()
	cr, err := GetCommiShared("PJ1901400","BK2C/pecf/pecf-prod-app","dev_core")
	if err == nil {
		fmt.Printf("%+v", cr)
	} else {
		fmt.Printf("%v", err)
	}
}