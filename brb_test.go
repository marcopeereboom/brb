package brb

import (
	"math/big"
	"testing"
)

func TestMath(t *testing.T) {
	i := New()
	rv, err := i.Run("-12+13*(15+1)")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("result: %v", rv)
	if c := rv.Cmp(new(big.Rat).SetFloat64(196.0)); c != 0 {
		t.Error("expected 196")
		return
	}
}

func TestInv(t *testing.T) {
	i := New()
	rv, err := i.Run("inv(3/2)")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("result: %v", rv)
	er, _ := new(big.Rat).SetString("2/3")
	if c := rv.Cmp(er); c != 0 {
		t.Error("expected 2/3")
		return
	}
}

func TestVar(t *testing.T) {
	i := New()
	_, err := i.Run("a=12.5")
	if err != ErrNoResult {
		t.Error(err)
		return
	}
	er, err := i.GetVar("a")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", er)
	rv, _ := new(big.Rat).SetString("12.5")
	if c := rv.Cmp(er); c != 0 {
		t.Error("expected 12.5")
		return
	}
}

func TestVarVar(t *testing.T) {
	i := New()
	_, err := i.Run("ia=1+12\nis=ia+1\nis=is+12\nis=ia+12")
	if err != ErrNoResult {
		t.Error(err)
		return
	}
}

func TestScriptNoArgs(t *testing.T) {
	i := New()
	rv, err := i.Run("1+2")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("result: %v", rv)
}

func TestScriptArgs(t *testing.T) {
	i := New()
	rv, err := i.Run("$1+$2",
		new(big.Rat).SetFloat64(3.0),
		new(big.Rat).SetFloat64(2.0))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("result: %v", rv)
}

func TestScriptArgsFail(t *testing.T) {
	i := New()
	_, err := i.Run("$1+$2",
		new(big.Rat).SetFloat64(3.0),
		"1+2")
	if err == nil {
		t.Error("string args should fail")
		return
	}
}

func TestScriptMissingVar(t *testing.T) {
	i := New()
	_, err := i.Run("$1+in",
		new(big.Rat).SetFloat64(3.0),
		new(big.Rat).SetFloat64(2.0))
	if err == nil {
		t.Errorf("unknown var")
		return
	}
}
