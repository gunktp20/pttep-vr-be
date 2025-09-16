package random_test

import (
	"pttep-vr-api/pkg/utils/random"
	"testing"
)

func Test(t *testing.T) {

	t.Run("Random", func(t *testing.T) {
		random.New([]random.CharacterSet{random.CharacterBig, random.CharacterSmall, random.Number, random.Special}, 16)
	})
}
