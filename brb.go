// Package brb (Basic Big.Rat) make using golang's big.Rats less of a pain in
// the butt.
// It provides a natural language interface to big.Rats instead of having to
// deal with assembly like mnemonics.
package brb

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
)

var (
	// ErrVariableNotFound is returned if the requested variable does not
	// exist.
	ErrVariableNotFound = errors.New("variable not found")

	// ErrNoResult is returned if the script has not return value but was.
	// successful.
	// Variables may or may not have been set.
	ErrNoResult = errors.New("no result")
)

// Interpreter is the context that is required to run expressions.
type Interpreter struct {
	lexer *yylexer
	mtx   sync.Mutex
}

// New returns an Interpreter context.
func New() *Interpreter {
	i := Interpreter{}
	return &i
}

// Run executes script and returns a *big.Rat if there is a result.
// If there is no result Run returns ErrNoResult.
// The following example returns ErrNoResult:
//	br, err := new(Interpreter).Run("a=12.5")
func (i *Interpreter) Run(script string, args ...interface{}) (*big.Rat, error) {
	r := bufio.NewReader(strings.NewReader(script))

	// make sure args are *big.Rat
	for k, v := range args {
		switch v.(type) {
		case *big.Rat:
		default:
			return nil, fmt.Errorf("invalid arg %v type %T", k+1, v)
		}
	}

	// for now don't allow reentrant calls
	i.mtx.Lock()
	defer i.mtx.Unlock()

	i.lexer = newLexer(r, args...)
	result := yyParse(i.lexer)
	if result == 0 {
		if i.lexer.result != nil {
			v := *i.lexer.result
			return &v, nil
		}
		return nil, ErrNoResult
	}
	return nil, i.lexer.lastError
}

// GetVar returns a variable if it exists.
// If the requested variable does not exist it returns ErrVariableNotFound.
// The following example returns a *big.Rat that is set to 12/1
//	i := New()
//	_, err := i.Run("a=12.0")
//	myvar, err := i.GetVar("a")
func (i *Interpreter) GetVar(name string) (*big.Rat, error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	val := i.lexer.getVar(name)
	if val == nil {
		return nil, ErrVariableNotFound
	}
	v := *val
	return &v, nil
}

// GetResult returns the result if it exists.
func (i *Interpreter) GetResult(name string) (*big.Rat, error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	val := i.lexer.result
	if val == nil {
		return nil, ErrNoResult
	}
	v := *val
	return &v, nil
}
