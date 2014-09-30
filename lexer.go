package brb

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

type yylexer struct {
	src       *bufio.Reader
	buf       []byte
	empty     bool
	current   byte
	lastError error
	result    *big.Rat // result of the script
	line      int      // line we are parsing
	args      []interface{}
	variables map[string]*big.Rat
}

func newLexer(src *bufio.Reader, args ...interface{}) *yylexer {
	y := yylexer{
		line:      1,
		src:       src,
		args:      args,
		variables: make(map[string]*big.Rat),
	}

	if b, err := src.ReadByte(); err == nil {
		y.current = b
	}

	return &y
}

func (y *yylexer) getc() byte {
	if y.current != 0 {
		y.buf = append(y.buf, y.current)
	}
	y.current = 0
	if b, err := y.src.ReadByte(); err == nil {
		y.current = b
	}
	return y.current
}

func (y *yylexer) Error(e string) {
	y.lastError = fmt.Errorf("line %v: %v", y.line, e)
}

func (y *yylexer) Errorf(format string, args ...interface{}) {
	line := fmt.Sprintf("line %v: ", y.line)
	y.lastError = fmt.Errorf(line+format, args...)
}

// lexer functions
func (y *yylexer) positional(val *yySymType, p string) int {
	idx, _ := strconv.Atoi(p[1:])
	idx-- // args are 1 base
	if idx < 0 || idx > len(y.args) {
		y.Errorf("invalid argument index %v", idx)
	}

	switch y.args[idx].(type) {
	case *big.Rat:
		val.lval = y.args[idx].(*big.Rat)
		return NUM
	default:
		y.Errorf("invalid argument type %T", y.args[idx])
	}

	panic("notreached")
	return NUM
}

func (y *yylexer) number(val *yySymType, s string) int {
	var ok bool
	val.lval, ok = new(big.Rat).SetString(s)
	if !ok {
		log.Fatal("invalid number")
	}
	return NUM
}

func (y *yylexer) variable(val *yySymType, s string) int {
	val.variable = string(y.buf)
	return VAR
}

// lang helpers
func (y *yylexer) getVar(s string) *big.Rat {
	return y.variables[s]
}
