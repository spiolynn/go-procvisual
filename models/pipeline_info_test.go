package models

import (
	"fmt"
	"testing"
)

func TestGetPipelineById(t *testing.T) {

	beforeTest()
	pl, err := GetPipelineById(1)
	if err == nil {
		fmt.Printf("%+v", *pl)
	} else {
		fmt.Printf("%v", err)
	}
}

