package boolconv

import "testing"

func TestNewBoolByBytes(t *testing.T) {

	ret := NewBoolByBytes([]byte{0})
	if ret == True {
		t.Error("expected boolconv.Flase got boolconv.True")
	}

	ret = NewBoolByBytes([]byte{1})
	if ret == False {
		t.Error("expected boolconv.True got boolconv.False")
	}

}

func TestNewBoolBybool(t *testing.T) {

	ret := NewBoolBybool(false)
	if ret == True {
		t.Error("expected boolconv.Flase got boolconv.True")
	}

	ret = NewBoolBybool(true)
	if ret == False {
		t.Error("expected boolconv.True got boolconv.False")
	}

}

func TestNewBoolByInterface(t *testing.T) {

	// Elem

	_, err := NewBoolByInterface(true)
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(false)
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(True)
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(False)
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(True.Bytes())
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(False.Bytes())
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(True.String())
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(False.String())
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(byte(1))
	if err != nil {
		t.Error("error should not be nil")
	}

	_, err = NewBoolByInterface(byte(0))
	if err != nil {
		t.Error("error should not be nil")
	}

	// pointer
	v := "true"
	_, err = NewBoolByInterface(&v)
	if err != nil {
		t.Error("error should not be nil")
	}

	// unsupported

	_, err = NewBoolByInterface([]string{"test"})
	if err == nil {
		t.Error("error is occurred")
	}

}

func TestTob(t *testing.T) {

	if False.Tob() {
		t.Error("expected false got true")
	}

	if !True.Tob() {
		t.Error("expected true got false")
	}

}

func TestBytes(t *testing.T) {

	if False.Bytes()[0] == byte(True) {
		t.Error("expected boolconv.Flase got boolconv.True")
	}

	if True.Bytes()[0] == byte(False) {
		t.Error("expected boolconv.True got boolconv.False")
	}

}

func TestString(t *testing.T) {

	if False.String() != "false" {
		t.Error("expected \"false\" got ", False.String())
	}

	if True.String() != "true" {
		t.Error("expected \"true\" got ", True.String())
	}

}
