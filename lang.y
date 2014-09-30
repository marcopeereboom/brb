%{
// THIS FILE IS AUTOGENERATED, DO NOT EDIT!
package brb

import (
    "math/big"
)

%}

%union{
    lval *big.Rat
    variable string
}

%token	NUM
%token	INV
%token	VAR

%left	'-' '+'
%left	'*' '/'
%left	NEG     /* negation--unary minus */

%type	<lval>		NUM, exp, INV
%type	<variable>	VAR

%%

input:    /* empty */
        | input line
;

line:     exp         { yylex.(*yylexer).result = $1 }
	| VAR '=' exp { yylex.(*yylexer).variables[$1] = $3 }
;

exp:	  '(' exp ')'        { $$ = new(big.Rat).Set($2)     }
        | exp '+' exp        { $$ = new(big.Rat).Add($1, $3) }
        | exp '-' exp        { $$ = new(big.Rat).Sub($1, $3) }
        | exp '*' exp        { $$ = new(big.Rat).Mul($1, $3) }
        | exp '/' exp        { $$ = new(big.Rat).Quo($1, $3) }
        | INV '(' exp ')'    { $$ = new(big.Rat).Inv($3)     }
        | '-' exp  %prec NEG { $$ = new(big.Rat).Neg($2)     }
	| NUM                { $$ = new(big.Rat).Set($1)     }
	| VAR {
		val := yylex.(*yylexer).getVar($1)
		if val == nil {
			yylex.(*yylexer).Errorf("unknown variable: %v", $1);
			goto ret1
		}
		$$=val
	}
;

%%
