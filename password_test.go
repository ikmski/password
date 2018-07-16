package password

import "testing"

func TestPasswordVerify(t *testing.T) {

	pp := &Policy{8, 0, 0}

	ok, _ := pp.Verify("abcd")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok, _ = pp.Verify("abcdEFGH")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

	pp = &Policy{8, 2, 0}
	ok, _ = pp.Verify("abcdEFGH")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok, _ = pp.Verify("abcdEF12")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

	pp = &Policy{8, 0, 2}
	ok, _ = pp.Verify("abcdEFGH")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok, _ = pp.Verify("@#$%^&*()-_=+,.?/:;{}[]~")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}

func TestPasswordLength(t *testing.T) {

	pp := &Policy{8, 0, 0}
	ok, _ := pp.hasEnoughLength("abcd")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok, _ = pp.hasEnoughLength("abcdEFGH")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}
	ok, _ = pp.hasEnoughLength("abcdEFGHij")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}

func TestPasswordDigit(t *testing.T) {

	pp := &Policy{8, 2, 0}
	ok, _ := pp.hasEnoughDigits("abcd")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok, _ = pp.hasEnoughDigits("abcdEF12")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}
	ok, _ = pp.hasEnoughDigits("abcdEF1234")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}

func TestPasswordSymbol(t *testing.T) {

	pp := &Policy{8, 0, 2}
	ok, _ := pp.hasEnoughSymbols("abcd")
	if ok {
		t.Errorf("got %v\nwant %v", ok, false)
	}
	ok, _ = pp.hasEnoughSymbols("abcdEF!@")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}
	ok, _ = pp.hasEnoughSymbols("abcdEF&*()")
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}
