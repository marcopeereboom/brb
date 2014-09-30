brb
===

Basic Big.Rat make using golang's big.Rats less of a pain in the butt.

When using lots of big.Rats in go it quickly becomes a painful.
Brb works around that by allowing natural expressions that deal with big.Rats.
For example:
	rv, err := i.Run("-12+13*(15+1)")

After running this rv contains a big.Rat that is set to 196/1.

Look at brb_test.go for more examples.

The repo comes with pre-generated lex and yacc target files.
One can generate those themselves by running make however, this should not be
needed.
