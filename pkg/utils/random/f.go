package random

import (
	"fmt"
	"math/rand"
)

type CharacterSet string

const (
	Number         CharacterSet = "NUMBER"
	CharacterSmall CharacterSet = "CHARACTER_SMALL"
	CharacterBig   CharacterSet = "CHARACTER_BIG"
	Special        CharacterSet = "SPECIAL"
)

func New(characterSet []CharacterSet, length int) string {
	var charSet string

	for _, c := range characterSet {
		switch c {
		case CharacterBig:
			charSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		case CharacterSmall:
			charSet += "abcdefghijklmnopqrstuvwxyz"
		case Number:
			charSet += "0123456789"
		case Special:
			charSet += "!@#?$&"
		}
	}
	var letterRunes = []rune(fmt.Sprintf("%s", charSet))

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
