package password

import (
	"math/rand"
	"regexp"
	"time"
)

var cnstCharas = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var cnstDigits = []rune("0123456789")
var cnstSymbols = []rune("!@#$%^&*()-_=+,.?/:;{}[]~")

var matchDigits = regexp.MustCompile(`[0-9]`)
var matchSymbols = regexp.MustCompile(`[\!\@\#\$\%\^\&\*\(\\\)\-_\=\+\,\.\?\/\:\;\{\}\[\]~]`)

type PasswordPolicy struct {
	Length  int
	Digits  int
	Symbols int
}

func NewPasswordPolicy() *PasswordPolicy {
	pp := new(PasswordPolicy)
	pp.Length = 8
	pp.Digits = 2
	pp.Symbols = 0
	return pp
}

func (pp *PasswordPolicy) Random() string {

	rand.Seed(time.Now().UTC().UnixNano())

	result := make([]rune, pp.Length)

	i := 0
	for i < pp.Symbols {
		result[i] = cnstSymbols[rand.Intn(len(cnstSymbols))]
		i++
	}
	for i < pp.Symbols+pp.Digits {
		result[i] = cnstDigits[rand.Intn(len(cnstDigits))]
		i++
	}
	for i < pp.Length {
		result[i] = cnstCharas[rand.Intn(len(cnstCharas))]
		i++
	}
	shuffle(result)
	return string(result)
}

func shuffle(data []rune) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func (pp *PasswordPolicy) Verify(pass string) bool {

	if pp.HasEnoughLength(pass) &&
		pp.HasEnoughDigits(pass) &&
		pp.HasEnoughSymbols(pass) {
		return true
	}

	return false
}

func (pp *PasswordPolicy) HasEnoughLength(pass string) bool {

	if len(pass) >= pp.Length {
		return true
	}

	return false
}

func (pp *PasswordPolicy) HasEnoughDigits(pass string) bool {

	s := matchDigits.FindAllString(pass, -1)
	if len(s) >= pp.Digits {
		return true
	}

	return false
}

func (pp *PasswordPolicy) HasEnoughSymbols(pass string) bool {

	s := matchSymbols.FindAllString(pass, -1)
	if len(s) >= pp.Symbols {
		return true
	}

	return false
}
