package color_test

import (
	"pttep-vr-api/pkg/constant/color"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Color", func(t *testing.T) {
		text := "text"
		color.Red(text)
		color.Green(text)
		color.Yellow(text)
		color.Blue(text)
		color.Purple(text)
		color.CadetBlue(text)
		color.White(text)
	})
}
