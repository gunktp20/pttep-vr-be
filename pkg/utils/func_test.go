package utils_test

import (
	"pttep-vr-api/pkg/utils"
	"testing"
)

func Test(t *testing.T) {

	t.Run("PrintJson", func(t *testing.T) {
		v := "string"
		utils.PrintJson(v)
	})
	t.Run("PrintJson", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic but got none")
			}
		}()
		v := make(chan int)
		utils.PrintJson(v)
	})

}
