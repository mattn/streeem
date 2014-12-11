//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/streeem/ast"
)

//line parser.go.y:16
type yySymType struct {
	yys         int
	stmt_if     ast.Stmt
	stmts       []ast.Stmt
	stmt        ast.Stmt
	expr        ast.Expr
	exprs       []ast.Expr
	expr_idents []string
	tok         ast.Token
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VARARG = 57350
const RETURN = 57351
const VAR = 57352
const THROW = 57353
const IF = 57354
const ELSE = 57355
const FOR = 57356
const IN = 57357
const EQEQ = 57358
const NEQ = 57359
const GE = 57360
const LE = 57361
const OROR = 57362
const ANDAND = 57363
const NEW = 57364
const TRUE = 57365
const FALSE = 57366
const NIL = 57367
const MODULE = 57368
const PLUSEQ = 57369
const MINUSEQ = 57370
const MULEQ = 57371
const DIVEQ = 57372
const ANDEQ = 57373
const OREQ = 57374
const BREAK = 57375
const CONTINUE = 57376
const PLUSPLUS = 57377
const MINUSMINUS = 57378
const POW = 57379
const SHIFTLEFT = 57380
const SHIFTRIGHT = 57381
const UNARY = 57382

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"RETURN",
	"VAR",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"OROR",
	"ANDAND",
	"NEW",
	"TRUE",
	"FALSE",
	"NIL",
	"MODULE",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"ANDEQ",
	"OREQ",
	"BREAK",
	"CONTINUE",
	"PLUSPLUS",
	"MINUSMINUS",
	"POW",
	"SHIFTLEFT",
	"SHIFTRIGHT",
	"'='",
	"'?'",
	"':'",
	"','",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:354

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 74,
	16, 0,
	17, 0,
	-2, 41,
	-1, 75,
	16, 0,
	17, 0,
	-2, 42,
}

const yyNprod = 61
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 776

var yyAct = []int{

	3, 108, 55, 51, 99, 120, 92, 52, 119, 117,
	114, 56, 57, 58, 98, 54, 93, 115, 97, 52,
	62, 63, 61, 55, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 96, 109, 86, 87, 88,
	89, 90, 1, 7, 2, 21, 52, 94, 0, 95,
	110, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 0, 0, 0, 64, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 32, 33, 35, 37,
	47, 49, 52, 52, 106, 104, 105, 38, 39, 40,
	41, 42, 43, 112, 0, 44, 45, 29, 30, 31,
	0, 23, 0, 0, 34, 36, 24, 25, 26, 27,
	28, 0, 0, 113, 0, 0, 0, 46, 50, 0,
	0, 0, 48, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 107, 0, 0,
	0, 0, 0, 0, 111, 0, 32, 33, 35, 37,
	47, 49, 0, 116, 0, 0, 118, 38, 39, 40,
	41, 42, 43, 0, 0, 44, 45, 29, 30, 31,
	0, 23, 0, 0, 34, 36, 24, 25, 26, 27,
	28, 0, 0, 0, 0, 0, 0, 46, 50, 103,
	0, 0, 48, 32, 33, 35, 37, 47, 49, 0,
	0, 0, 0, 0, 38, 39, 40, 41, 42, 43,
	0, 0, 44, 45, 29, 30, 31, 0, 23, 102,
	0, 34, 36, 24, 25, 26, 27, 28, 0, 0,
	0, 0, 0, 0, 46, 50, 0, 0, 0, 48,
	32, 33, 35, 37, 47, 49, 0, 0, 0, 0,
	0, 38, 39, 40, 41, 42, 43, 0, 0, 44,
	45, 29, 30, 31, 0, 23, 0, 0, 34, 36,
	24, 25, 26, 27, 28, 0, 0, 101, 0, 0,
	0, 46, 50, 0, 0, 0, 48, 32, 33, 35,
	37, 47, 49, 0, 0, 0, 0, 0, 38, 39,
	40, 41, 42, 43, 0, 0, 44, 45, 29, 30,
	31, 0, 23, 0, 0, 34, 36, 24, 25, 26,
	27, 28, 0, 0, 0, 0, 0, 0, 46, 50,
	0, 0, 100, 48, 32, 33, 35, 37, 47, 49,
	0, 0, 0, 0, 0, 38, 39, 40, 41, 42,
	43, 0, 0, 44, 45, 29, 30, 31, 0, 23,
	0, 91, 34, 36, 24, 25, 26, 27, 28, 0,
	0, 0, 0, 0, 0, 46, 50, 0, 0, 0,
	48, 32, 33, 35, 37, 47, 49, 0, 0, 0,
	0, 0, 38, 39, 40, 41, 42, 43, 0, 0,
	44, 45, 29, 30, 31, 0, 23, 0, 0, 34,
	36, 24, 25, 26, 27, 28, 0, 0, 0, 0,
	0, 0, 46, 50, 0, 0, 0, 48, 32, 33,
	35, 37, 0, 49, 0, 0, 0, 0, 0, 38,
	39, 40, 41, 42, 43, 0, 0, 44, 45, 29,
	30, 31, 0, 0, 0, 0, 34, 36, 24, 25,
	26, 27, 28, 0, 32, 33, 35, 37, 0, 46,
	50, 0, 0, 0, 48, 38, 39, 40, 41, 42,
	43, 0, 0, 44, 45, 29, 30, 31, 0, 0,
	0, 0, 34, 36, 24, 25, 26, 27, 28, 0,
	0, 0, 35, 37, 0, 46, 50, 0, 0, 0,
	48, 38, 39, 40, 41, 42, 43, 0, 0, 44,
	45, 29, 30, 31, 0, 0, 0, 0, 34, 36,
	24, 25, 26, 27, 28, 38, 39, 40, 41, 42,
	43, 46, 50, 44, 45, 29, 48, 0, 0, 0,
	0, 0, 0, 0, 24, 25, 26, 27, 28, 8,
	9, 13, 0, 0, 6, 46, 50, 20, 0, 0,
	48, 0, 0, 0, 0, 0, 0, 0, 14, 15,
	16, 0, 0, 0, 0, 0, 0, 0, 4, 5,
	0, 38, 39, 40, 41, 42, 43, 0, 8, 9,
	13, 29, 10, 6, 0, 0, 20, 0, 17, 0,
	11, 12, 59, 18, 0, 19, 0, 14, 15, 16,
	0, 46, 50, 0, 0, 0, 48, 4, 5, 0,
	0, 0, 0, 0, 0, 0, 0, 8, 9, 13,
	0, 10, 6, 0, 0, 20, 22, 17, 0, 11,
	12, 0, 18, 0, 19, 0, 14, 15, 16, 0,
	0, 0, 0, 0, 0, 0, 4, 5, 0, 0,
	0, 0, 0, 38, 39, 40, 41, 42, 43, 0,
	10, 0, 0, 29, 0, 0, 17, 0, 11, 12,
	0, 18, 0, 19, 26, 27, 28, 8, 9, 13,
	0, 0, 0, 46, 50, 0, 0, 0, 48, 53,
	9, 13, 0, 0, 0, 0, 14, 15, 16, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 14, 15,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	10, 0, 0, 0, 0, 0, 17, 0, 11, 12,
	0, 18, 10, 19, 0, 0, 0, 0, 17, 0,
	11, 12, 0, 18, 0, 19,
}
var yyPact = []int{

	643, -1000, 604, 375, -1000, -1000, 715, 2, -58, -1000,
	703, 703, 703, -1000, -1000, -1000, -1000, 565, 715, 703,
	703, -1000, 643, 703, 703, 703, 703, 703, 703, 703,
	703, 703, 703, 703, 703, 703, 703, 703, 703, 703,
	703, 703, 703, 703, -1000, -1000, 703, 703, 703, 703,
	703, -1000, 328, -37, 4, 715, 574, 574, 574, 14,
	-40, -55, 281, 234, -1000, 187, 656, 656, 574, 574,
	574, 375, 518, 518, 494, 494, 518, 518, 518, 518,
	375, 375, 375, 375, 375, 375, 375, 422, 375, 458,
	140, 715, 715, 703, 643, -60, 3, -1000, -1000, -1000,
	-1000, 643, 703, -1000, -1000, -1000, 70, -44, -1000, 13,
	643, -45, 375, 643, -1000, -1000, -46, -1000, -49, -1000,
	-1000,
}
var yyPgo = []int{

	0, 52, 54, 53, 0, 3, 45,
}
var yyR1 = []int{

	0, 1, 1, 1, 2, 2, 2, 2, 2, 3,
	3, 3, 6, 6, 6, 5, 5, 5, 5, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4,
}
var yyR2 = []int{

	0, 0, 2, 3, 1, 1, 1, 2, 1, 7,
	5, 5, 0, 1, 3, 0, 1, 3, 3, 1,
	1, 2, 2, 2, 1, 1, 1, 1, 5, 6,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 3, 3, 3, 3, 4,
	4,
}
var yyChk = []int{

	-1000, -1, -2, -4, 33, 34, 9, -3, 4, 5,
	47, 55, 56, 6, 23, 24, 25, 53, 58, 60,
	12, -1, 52, 41, 46, 47, 48, 49, 50, 37,
	38, 39, 16, 17, 44, 18, 45, 19, 27, 28,
	29, 30, 31, 32, 35, 36, 57, 20, 62, 21,
	58, -5, -4, 4, 13, 60, -4, -4, -4, 57,
	-1, -5, -4, -4, -1, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, 43, 43, 12, 53, -5, -6, 4, 54, 59,
	61, 53, 42, 59, -5, -5, -4, -1, 61, 43,
	57, -1, -4, 53, 54, 4, -1, 54, -1, 54,
	54,
}
var yyDef = []int{

	1, -2, 1, 4, 5, 6, 15, 8, 19, 20,
	0, 0, 0, 24, 25, 26, 27, 1, 15, 0,
	0, 2, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 53, 54, 0, 0, 0, 0,
	0, 7, 16, 19, 0, 15, 21, 22, 23, 12,
	0, 0, 0, 0, 3, 0, 33, 34, 35, 36,
	37, 38, 39, 40, -2, -2, 43, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 55, 56, 57, 58,
	0, 15, 15, 0, 1, 0, 0, 13, 30, 31,
	32, 1, 0, 60, 17, 18, 0, 0, 59, 0,
	1, 0, 28, 1, 10, 14, 0, 11, 0, 29,
	9,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 3, 3, 50, 62, 3,
	60, 61, 48, 46, 43, 47, 3, 49, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 42, 52,
	45, 40, 44, 41, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 58, 3, 59, 56, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 53, 57, 54,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 51,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:43
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:50
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				for _, s := range yyS[yypt-0].stmts {
					if yyS[yypt-1].stmt.Position().Line == s.Position().Line {
						l.pos = yyS[yypt-1].stmt.Position()
						yylex.Error("syntax error")
					}
				}
			}
		}
	case 3:
		//line parser.go.y:63
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPosition(l.pos)
			}
		}
	case 4:
		//line parser.go.y:72
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			yyVAL.stmt.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 5:
		//line parser.go.y:77
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 6:
		//line parser.go.y:82
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 7:
		//line parser.go.y:87
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
			yyVAL.stmt.SetPosition(yyS[yypt-1].tok.Position())
		}
	case 8:
		//line parser.go.y:92
		{
			yyVAL.stmt = yyS[yypt-0].stmt_if
			yyVAL.stmt.SetPosition(yyS[yypt-0].stmt_if.Position())
		}
	case 9:
		//line parser.go.y:98
		{
			yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
		}
	case 10:
		//line parser.go.y:102
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyS[yypt-1].stmts...)
			}
		}
	case 11:
		//line parser.go.y:110
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
		}
	case 12:
		//line parser.go.y:115
		{
			yyVAL.expr_idents = []string{}
		}
	case 13:
		//line parser.go.y:119
		{
			yyVAL.expr_idents = []string{yyS[yypt-0].tok.Lit}
		}
	case 14:
		//line parser.go.y:123
		{
			yyVAL.expr_idents = append(yyS[yypt-2].expr_idents, yyS[yypt-0].tok.Lit)
		}
	case 15:
		//line parser.go.y:128
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 16:
		//line parser.go.y:132
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 17:
		//line parser.go.y:136
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 18:
		//line parser.go.y:140
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.Lit}}, yyS[yypt-0].exprs...)
		}
	case 19:
		//line parser.go.y:146
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 20:
		//line parser.go.y:151
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 21:
		//line parser.go.y:156
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 22:
		//line parser.go.y:161
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 23:
		//line parser.go.y:166
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 24:
		//line parser.go.y:171
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 25:
		//line parser.go.y:176
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 26:
		//line parser.go.y:181
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 27:
		//line parser.go.y:186
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 28:
		//line parser.go.y:191
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-4].expr.Position())
		}
	case 29:
		//line parser.go.y:196
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-3].expr_idents, Stmts: yyS[yypt-1].stmts}
		}
	case 30:
		//line parser.go.y:200
		{
			yyVAL.expr = &ast.FuncExpr{Stmts: yyS[yypt-1].stmts}
		}
	case 31:
		//line parser.go.y:204
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 32:
		//line parser.go.y:209
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 33:
		//line parser.go.y:214
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 34:
		//line parser.go.y:219
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 35:
		//line parser.go.y:224
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 36:
		//line parser.go.y:229
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 37:
		//line parser.go.y:234
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 38:
		//line parser.go.y:239
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 39:
		//line parser.go.y:244
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 40:
		//line parser.go.y:249
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 41:
		//line parser.go.y:254
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 42:
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 43:
		//line parser.go.y:264
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 44:
		//line parser.go.y:269
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 45:
		//line parser.go.y:274
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 46:
		//line parser.go.y:279
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 47:
		//line parser.go.y:284
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "+=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 48:
		//line parser.go.y:289
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "-=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 49:
		//line parser.go.y:294
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "*=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 50:
		//line parser.go.y:299
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "/=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 51:
		//line parser.go.y:304
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "&=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 52:
		//line parser.go.y:309
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "|=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 53:
		//line parser.go.y:314
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyS[yypt-1].expr.Position())
		}
	case 54:
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyS[yypt-1].expr.Position())
		}
	case 55:
		//line parser.go.y:324
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 56:
		//line parser.go.y:329
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 57:
		//line parser.go.y:334
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 58:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 59:
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.Lit, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 60:
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			yyVAL.expr.SetPosition(yyS[yypt-3].expr.Position())
		}
	}
	goto yystack /* stack new state and value */
}
