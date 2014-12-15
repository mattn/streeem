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
const IF = 57353
const ELSE = 57354
const FOR = 57355
const IN = 57356
const EQEQ = 57357
const NEQ = 57358
const GE = 57359
const LE = 57360
const OROR = 57361
const ANDAND = 57362
const NEW = 57363
const TRUE = 57364
const FALSE = 57365
const NIL = 57366
const MODULE = 57367
const PLUSEQ = 57368
const MINUSEQ = 57369
const MULEQ = 57370
const DIVEQ = 57371
const ANDEQ = 57372
const OREQ = 57373
const BREAK = 57374
const CONTINUE = 57375
const PLUSPLUS = 57376
const MINUSMINUS = 57377
const POW = 57378
const SHIFTLEFT = 57379
const SHIFTRIGHT = 57380
const UNARY = 57381

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"RETURN",
	"VAR",
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
	"'|'",
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

//line parser.go.y:361

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 66,
	15, 56,
	16, 56,
	17, 56,
	18, 56,
	19, 56,
	20, 56,
	37, 56,
	38, 56,
	40, 56,
	43, 56,
	44, 56,
	-2, 4,
	-1, 76,
	15, 0,
	16, 0,
	-2, 42,
	-1, 77,
	15, 0,
	16, 0,
	-2, 43,
}

const yyNprod = 62
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 789

var yyAct = []int{

	3, 110, 56, 100, 52, 94, 122, 53, 121, 119,
	116, 57, 58, 59, 111, 60, 95, 112, 55, 53,
	62, 63, 56, 61, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 117, 99, 88,
	89, 90, 91, 1, 98, 7, 21, 53, 96, 2,
	0, 97, 0, 0, 0, 0, 0, 40, 41, 42,
	43, 44, 45, 0, 0, 0, 64, 31, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 28,
	29, 30, 0, 53, 106, 53, 108, 105, 51, 107,
	0, 0, 49, 0, 114, 34, 35, 37, 39, 48,
	50, 0, 0, 0, 0, 0, 40, 41, 42, 43,
	44, 45, 0, 0, 46, 47, 31, 32, 33, 0,
	25, 0, 0, 36, 38, 93, 26, 27, 28, 29,
	30, 0, 0, 115, 0, 0, 0, 51, 0, 0,
	109, 49, 0, 0, 0, 0, 113, 0, 0, 34,
	35, 37, 39, 48, 50, 0, 118, 0, 0, 120,
	40, 41, 42, 43, 44, 45, 0, 0, 46, 47,
	31, 32, 33, 0, 25, 0, 0, 36, 38, 93,
	26, 27, 28, 29, 30, 34, 35, 37, 39, 48,
	50, 51, 104, 0, 0, 49, 40, 41, 42, 43,
	44, 45, 0, 0, 46, 47, 31, 32, 33, 0,
	25, 103, 0, 36, 38, 93, 26, 27, 28, 29,
	30, 34, 35, 37, 39, 48, 50, 51, 0, 0,
	0, 49, 40, 41, 42, 43, 44, 45, 0, 0,
	46, 47, 31, 32, 33, 0, 25, 0, 0, 36,
	38, 93, 26, 27, 28, 29, 30, 0, 0, 102,
	0, 0, 0, 51, 0, 0, 0, 49, 34, 35,
	37, 39, 48, 50, 0, 0, 0, 0, 0, 40,
	41, 42, 43, 44, 45, 0, 0, 46, 47, 31,
	32, 33, 0, 25, 0, 0, 36, 38, 93, 26,
	27, 28, 29, 30, 34, 35, 37, 39, 48, 50,
	51, 0, 0, 101, 49, 40, 41, 42, 43, 44,
	45, 0, 0, 46, 47, 31, 32, 33, 0, 25,
	0, 92, 36, 38, 93, 26, 27, 28, 29, 30,
	34, 35, 37, 39, 48, 50, 51, 0, 0, 0,
	49, 40, 41, 42, 43, 44, 45, 0, 0, 46,
	47, 31, 32, 33, 0, 25, 0, 0, 36, 38,
	93, 26, 27, 28, 29, 30, 34, 35, 37, 39,
	48, 50, 51, 0, 0, 0, 49, 40, 41, 42,
	43, 44, 45, 0, 0, 46, 47, 31, 32, 33,
	0, 25, 0, 0, 36, 38, 24, 26, 27, 28,
	29, 30, 34, 35, 37, 39, 48, 50, 51, 0,
	0, 0, 49, 40, 41, 42, 43, 44, 45, 0,
	0, 46, 47, 31, 32, 33, 0, 25, 0, 0,
	36, 38, 0, 26, 27, 28, 29, 30, 34, 35,
	37, 39, 0, 50, 51, 0, 0, 0, 49, 40,
	41, 42, 43, 44, 45, 0, 0, 46, 47, 31,
	32, 33, 0, 0, 0, 0, 36, 38, 93, 26,
	27, 28, 29, 30, 34, 35, 37, 39, 0, 0,
	51, 0, 0, 0, 49, 40, 41, 42, 43, 44,
	45, 0, 0, 46, 47, 31, 32, 33, 0, 0,
	0, 0, 36, 38, 93, 26, 27, 28, 29, 30,
	0, 0, 37, 39, 0, 0, 51, 0, 0, 0,
	49, 40, 41, 42, 43, 44, 45, 0, 0, 46,
	47, 31, 32, 33, 0, 0, 0, 0, 36, 38,
	93, 26, 27, 28, 29, 30, 0, 0, 0, 0,
	0, 0, 51, 0, 0, 0, 49, 40, 41, 42,
	43, 44, 45, 0, 0, 46, 47, 31, 0, 0,
	0, 0, 0, 0, 0, 0, 93, 26, 27, 28,
	29, 30, 0, 0, 0, 0, 0, 0, 51, 0,
	0, 0, 49, 40, 41, 42, 43, 44, 45, 0,
	0, 46, 47, 31, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 26, 27, 28, 29, 30, 8, 9,
	13, 0, 0, 6, 51, 20, 0, 0, 49, 40,
	41, 42, 43, 44, 45, 0, 14, 15, 16, 31,
	0, 0, 0, 0, 0, 0, 4, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 23,
	51, 10, 0, 0, 49, 0, 22, 17, 0, 11,
	12, 18, 0, 19, 8, 9, 13, 0, 0, 6,
	0, 20, 0, 0, 0, 0, 0, 8, 9, 13,
	0, 0, 14, 15, 16, 0, 0, 0, 0, 0,
	0, 0, 4, 5, 0, 14, 15, 16, 0, 0,
	0, 0, 0, 54, 9, 13, 0, 10, 0, 0,
	0, 0, 0, 17, 0, 11, 12, 18, 0, 19,
	10, 14, 15, 16, 0, 0, 17, 0, 11, 12,
	18, 0, 19, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 17, 0, 11, 12, 18, 0, 19,
}
var yyPact = []int{

	690, -1000, 634, 371, -1000, -1000, 729, 6, -57, -1000,
	703, 703, 703, -1000, -1000, -1000, -1000, -30, 729, 703,
	703, -1000, 690, 703, 703, 703, 703, 703, 703, 703,
	703, 703, 703, 703, 703, 703, 703, 703, 703, 703,
	703, 703, 703, 703, 703, 703, -1000, -1000, 703, 703,
	703, 703, -1000, 299, -37, 5, 729, 623, 623, 623,
	44, -55, 263, 216, -1000, 407, 587, 180, 41, 41,
	623, 623, 623, 335, 551, 551, 515, 515, 551, 551,
	551, 551, 335, 335, 335, 335, 335, 335, 443, 335,
	479, 144, 729, 703, 729, 703, 690, -59, -28, -1000,
	-1000, -1000, 690, 703, -1000, -1000, 587, -1000, 90, -44,
	-1000, 43, 690, -45, 335, 690, -1000, -1000, -46, -1000,
	-48, -1000, -1000,
}
var yyPgo = []int{

	0, 53, 59, 55, 0, 4, 54,
}
var yyR1 = []int{

	0, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 6, 6, 6, 5, 5, 5,
	5, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4,
}
var yyR2 = []int{

	0, 0, 2, 3, 3, 3, 1, 1, 1, 2,
	1, 7, 5, 5, 0, 1, 3, 0, 1, 3,
	3, 1, 1, 2, 2, 2, 1, 1, 1, 1,
	5, 6, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 3, 3, 3, 3,
	4, 4,
}
var yyChk = []int{

	-1000, -1, -2, -4, 32, 33, 9, -3, 4, 5,
	47, 55, 56, 6, 22, 23, 24, 53, 57, 59,
	11, -1, 52, 45, 45, 40, 46, 47, 48, 49,
	50, 36, 37, 38, 15, 16, 43, 17, 44, 18,
	26, 27, 28, 29, 30, 31, 34, 35, 19, 61,
	20, 57, -5, -4, 4, 12, 59, -4, -4, -4,
	45, -5, -4, -4, -1, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, 42, 45, 42, 11, 53, -5, -6, 4,
	58, 60, 53, 41, 58, -5, -4, -5, -4, -1,
	60, 42, 45, -1, -4, 53, 54, 4, -1, 54,
	-1, 54, 54,
}
var yyDef = []int{

	1, -2, 1, 6, 7, 8, 17, 10, 21, 22,
	0, 0, 0, 26, 27, 28, 29, 0, 17, 0,
	0, 2, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 0, 9, 18, 21, 0, 17, 23, 24, 25,
	14, 0, 0, 0, 3, 5, -2, 0, 34, 35,
	36, 37, 38, 39, 40, 41, -2, -2, 44, 45,
	46, 47, 48, 49, 50, 51, 52, 53, 57, 58,
	59, 0, 17, 0, 17, 0, 1, 0, 0, 15,
	32, 33, 1, 0, 61, 19, 56, 20, 0, 0,
	60, 0, 1, 0, 30, 1, 12, 16, 0, 13,
	0, 31, 11,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 3, 3, 50, 61, 3,
	59, 60, 48, 46, 42, 47, 3, 49, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 41, 52,
	44, 39, 43, 40, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 57, 3, 58, 56, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 53, 45, 54,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 51,
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
		//line parser.go.y:44
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:51
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
		//line parser.go.y:64
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPosition(l.pos)
			}
		}
	case 4:
		//line parser.go.y:73
		{
			yyVAL.stmt = &ast.PipeLine{Exprs: []ast.Expr{yyS[yypt-2].expr, yyS[yypt-0].expr}}
			yyVAL.stmt.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 5:
		//line parser.go.y:78
		{
			yyS[yypt-2].stmt.(*ast.PipeLine).Exprs = append(yyS[yypt-2].stmt.(*ast.PipeLine).Exprs, yyS[yypt-0].expr)
			yyVAL.stmt.SetPosition(yyS[yypt-2].stmt.Position())
		}
	case 6:
		//line parser.go.y:83
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			yyVAL.stmt.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 7:
		//line parser.go.y:88
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 8:
		//line parser.go.y:93
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 9:
		//line parser.go.y:98
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
			yyVAL.stmt.SetPosition(yyS[yypt-1].tok.Position())
		}
	case 10:
		//line parser.go.y:103
		{
			yyVAL.stmt = yyS[yypt-0].stmt_if
			yyVAL.stmt.SetPosition(yyS[yypt-0].stmt_if.Position())
		}
	case 11:
		//line parser.go.y:109
		{
			yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
		}
	case 12:
		//line parser.go.y:113
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyS[yypt-1].stmts...)
			}
		}
	case 13:
		//line parser.go.y:121
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
		}
	case 14:
		//line parser.go.y:126
		{
			yyVAL.expr_idents = []string{}
		}
	case 15:
		//line parser.go.y:130
		{
			yyVAL.expr_idents = []string{yyS[yypt-0].tok.Lit}
		}
	case 16:
		//line parser.go.y:134
		{
			yyVAL.expr_idents = append(yyS[yypt-2].expr_idents, yyS[yypt-0].tok.Lit)
		}
	case 17:
		//line parser.go.y:139
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 18:
		//line parser.go.y:143
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 19:
		//line parser.go.y:147
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 20:
		//line parser.go.y:151
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.Lit}}, yyS[yypt-0].exprs...)
		}
	case 21:
		//line parser.go.y:157
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 22:
		//line parser.go.y:162
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 23:
		//line parser.go.y:167
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 24:
		//line parser.go.y:172
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 25:
		//line parser.go.y:177
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 26:
		//line parser.go.y:182
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 27:
		//line parser.go.y:187
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 28:
		//line parser.go.y:192
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 29:
		//line parser.go.y:197
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 30:
		//line parser.go.y:202
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-4].expr.Position())
		}
	case 31:
		//line parser.go.y:207
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-3].expr_idents, Stmts: yyS[yypt-1].stmts}
		}
	case 32:
		//line parser.go.y:211
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 33:
		//line parser.go.y:216
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 34:
		//line parser.go.y:221
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 35:
		//line parser.go.y:226
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 36:
		//line parser.go.y:231
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 37:
		//line parser.go.y:236
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 38:
		//line parser.go.y:241
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 39:
		//line parser.go.y:246
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 40:
		//line parser.go.y:251
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 41:
		//line parser.go.y:256
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 42:
		//line parser.go.y:261
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 43:
		//line parser.go.y:266
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 44:
		//line parser.go.y:271
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 45:
		//line parser.go.y:276
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 46:
		//line parser.go.y:281
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 47:
		//line parser.go.y:286
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 48:
		//line parser.go.y:291
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "+=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 49:
		//line parser.go.y:296
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "-=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 50:
		//line parser.go.y:301
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "*=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 51:
		//line parser.go.y:306
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "/=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 52:
		//line parser.go.y:311
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "&=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 53:
		//line parser.go.y:316
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "|=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 54:
		//line parser.go.y:321
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyS[yypt-1].expr.Position())
		}
	case 55:
		//line parser.go.y:326
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyS[yypt-1].expr.Position())
		}
	case 56:
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 57:
		//line parser.go.y:336
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 58:
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 59:
		//line parser.go.y:346
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 60:
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.Lit, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 61:
		//line parser.go.y:356
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			yyVAL.expr.SetPosition(yyS[yypt-3].expr.Position())
		}
	}
	goto yystack /* stack new state and value */
}
