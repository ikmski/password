package password

import "testing"

func TestPasswordVerify(t *testing.T) {

	pp := NewPasswordPolicy()

	pp.Length = 8
	pp.Digits = 0
	pp.Symbols = 0
	ok := pp.Verify("abcd")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok = pp.Verify("abcdEFGH")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

	pp.Length = 8
	pp.Digits = 2
	pp.Symbols = 0
	ok = pp.Verify("abcdEFGH")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok = pp.Verify("abcdEF12")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

	pp.Length = 8
	pp.Digits = 0
	pp.Symbols = 2
	ok = pp.Verify("abcdEFGH")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok = pp.Verify("@#$%^&*()-_=+,.?/:;{}[]~")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}

func TestPasswordLength(t *testing.T) {

	pp := NewPasswordPolicy()

	pp.Length = 8
	pp.Digits = 0
	pp.Symbols = 0
	ok := pp.HasEnoughLength("abcd")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok = pp.HasEnoughLength("abcdEFGH")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}
	ok = pp.HasEnoughLength("abcdEFGHij")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}

func TestPasswordDigit(t *testing.T) {

	pp := NewPasswordPolicy()

	pp.Length = 8
	pp.Digits = 2
	pp.Symbols = 0
	ok := pp.HasEnoughDigits("abcd")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok = pp.HasEnoughDigits("abcdEF12")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}
	ok = pp.HasEnoughDigits("abcdEF1234")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}

func TestPasswordSymbol(t *testing.T) {

	pp := NewPasswordPolicy()

	pp.Length = 8
	pp.Digits = 0
	pp.Symbols = 2
	ok := pp.HasEnoughSymbols("abcd")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok = pp.HasEnoughSymbols("abcdEF!@")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}
	ok = pp.HasEnoughSymbols("abcdEF&*()")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}
