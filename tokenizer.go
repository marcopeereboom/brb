// THIS FILE IS AUTOGENERATED, DO NOT EDIT!
package brb

func (y *yylexer) Lex(val *yySymType) int {
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

yystate0:

	y.buf = y.buf[:0]

	goto yystart1

	goto yystate1 // silence unused label error
yystate1:
	c = y.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '$':
		goto yystate4
	case c == '.':
		goto yystate6
	case c == '\n' || c == '\r':
		goto yystate3
	case c == '\t' || c == ' ':
		goto yystate2
	case c == 'i':
		goto yystate13
	case c >= '0' && c <= '9':
		goto yystate11
	case c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate12
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	goto yyrule2

yystate4:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate5
	}

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9':
		goto yystate5
	}

yystate6:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate7:
	c = y.getc()
	switch {
	default:
		goto yyrule5
	case c == 'E' || c == 'e':
		goto yystate8
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate8:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate9
	case c >= '0' && c <= '9':
		goto yystate10
	}

yystate9:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate10
	}

yystate10:
	c = y.getc()
	switch {
	default:
		goto yyrule5
	case c >= '0' && c <= '9':
		goto yystate10
	}

yystate11:
	c = y.getc()
	switch {
	default:
		goto yyrule5
	case c == '.':
		goto yystate7
	case c == 'E' || c == 'e':
		goto yystate8
	case c >= '0' && c <= '9':
		goto yystate11
	}

yystate12:
	c = y.getc()
	switch {
	default:
		goto yyrule6
	case c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate13:
	c = y.getc()
	switch {
	default:
		goto yyrule6
	case c == 'n':
		goto yystate14
	case c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate12
	}

yystate14:
	c = y.getc()
	switch {
	default:
		goto yyrule6
	case c == 'v':
		goto yystate15
	case c >= 'a' && c <= 'u' || c >= 'w' && c <= 'z':
		goto yystate12
	}

yystate15:
	c = y.getc()
	switch {
	default:
		goto yyrule3
	case c >= 'a' && c <= 'z':
		goto yystate12
	}

yyrule1: // [ \t]+

	goto yystate0
yyrule2: // [\n\r]
	{
		y.line++
		goto yystate0
	}
yyrule3: // "inv"
	{
		return INV
	}
yyrule4: // "$"{DIGIT}
	{
		return y.positional(val, string(y.buf))
	}
yyrule5: // {NUMBER}
	{
		return y.number(val, string(y.buf))
	}
yyrule6: // {VARIABLE}
	{
		return y.variable(val, string(y.buf))
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	y.empty = true
	return int(c)
}