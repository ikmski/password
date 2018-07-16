package password

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var cnstCharas = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var cnstDigits = []rune("0123456789")
var cnstSymbols = []rune("!@#$%^&*()-_=+,.?/:;{}[]~")

var matchDigits = regexp.MustCompile(`[0-9]`)
var matchSymbols = regexp.MustCompile(`[\!\@\#\$\%\^\&\*\(\\\)\-_\=\+\,\.\?\/\:\;\{\}\[\]~]`)

type Policy struct {
	Length  int
	Digits  int
	Symbols int
}

func Default() *Policy {
	pp := new(Policy)
	pp.Length = 12
	pp.Digits = 2
	pp.Symbols = 2
	return pp
}

func (pp *Policy) Random() string {

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

func (pp *Policy) Verify(pass string) (bool, []string) {

	var result = true
	var messages = make([]string, 0)

	ok, message := pp.hasEnoughLength(pass)
	if !ok {
		result = false
		messages = append(messages, message)
	}

	ok, message = pp.hasEnoughDigits(pass)
	if !ok {
		result = false
		messages = append(messages, message)
	}

	ok, message = pp.hasEnoughSymbols(pass)
	if !ok {
		result = false
		messages = append(messages, message)
	}

	return result, messages
}

func (pp *Policy) hasEnoughLength(pass string) (bool, string) {

	if len(pass) >= pp.Length {
		return true, ""
	}

	message := fmt.Sprintf("Your password does not have enough length.\nIt needs %d or more.", pp.Length)
	return false, message
}

func (pp *Policy) hasEnoughDigits(pass string) (bool, string) {

	s := matchDigits.FindAllString(pass, -1)
	if len(s) >= pp.Digits {
		return true, ""
	}

	message := fmt.Sprintf("Your password does not have enough digits.\nIt needs %d or more.", pp.Digits)
	return false, message
}

func (pp *Policy) hasEnoughSymbols(pass string) (bool, string) {

	s := matchSymbols.FindAllString(pass, -1)
	if len(s) >= pp.Symbols {
		return true, ""
	}

	message := fmt.Sprintf("Your password does not have enough symbols.\nIt needs %d or more.", pp.Digits)
	return false, message
}
