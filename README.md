brb
===

Basic Big.Rat make using golang's big.Rats less of a pain in the butt.

When using lots of big.Rats in go it quickly becomes a painful.
Brb works around that by allowing natural expressions that deal with big.Rats.
For example:
```Go
	rv, err := new(brb.Interpreter).Run("-12+13*(15+1)")
```
After running this rv contains a big.Rat that is set to 196/1.

It supports positional arguments as well as variables.
For example:
```Go
	rv, err := new(brb.Interpreter).Run(`
		a = 99
		$1 + a
	`, new(big.Rat).SetFloat64(1.0))
```
After running this rv contains a big.Rat that is set to 100/1.

Variables can be  interrogated after the fact.
For example:
```Go
	i := brb.New()
	_, err := i.Run(`
		a = 112.8
		a = a + 42.2
	`)
	er, err := i.GetVar("a")
```
After running this er contains a big.Rat that is set to 155/1.

Look at brb_test.go for more examples.

The repository comes with pre-generated lex and yacc target files.
One can generate those themselves by running make however, this should not be
needed.
