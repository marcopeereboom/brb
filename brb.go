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
	ErrVariableNotFound = errors.New("variable not found")
	ErrNoResult         = errors.New("no result")
)

type Interpreter struct {
	lexer *yylexer
	mtx   sync.Mutex
}

func New() *Interpreter {
	i := Interpreter{}
	return &i
}

func (i *Interpreter) Run(script string, args ...interface{}) (*big.Rat, error) {
	r := bufio.NewReader(strings.NewReader(script))

	// make sure args are either *big.Rat
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
