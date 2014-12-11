package parser

import (
	"errors"
	"fmt"
	"github.com/mattn/streeem/ast"
)

const (
	EOF = -1   // End of file.
	EOL = '\n' // End of line.
)

var opName = map[string]int{
	"return":   RETURN,
	"if":       IF,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"in":       IN,
	"else":     ELSE,
	"true":     TRUE,
	"false":    FALSE,
	"nil":      NIL,
}

type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

func (s *Scanner) Scan() (tok int, lit string, pos ast.Position, err error) {
retry:
	s.skipBlank()
	pos = s.pos()
	switch ch := s.peek(); {
	case isLetter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			return
		}
		if name, ok := opName[lit]; ok {
			tok = name
		} else {
			tok = IDENT
		}
	case isDigit(ch):
		tok = NUMBER
		lit, err = s.scanNumber()
		if err != nil {
			return
		}
	case ch == '"':
		tok = STRING
		lit, err = s.scanString('"')
		if err != nil {
			return
		}
	case ch == '\'':
		tok = STRING
		lit, err = s.scanString('\'')
		if err != nil {
			return
		}
	case ch == '`':
		tok = STRING
		lit, err = s.scanRawString()
		if err != nil {
			return
		}
	default:
		switch ch {
		case -1:
			tok = EOF
		case '#':
			for !isEOL(s.peek()) {
				s.next()
			}
			s.next()
			goto retry
		case '!':
			s.next()
			switch s.peek() {
			case '=':
				tok = NEQ
				lit = "!="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '=':
			s.next()
			switch s.peek() {
			case '=':
				tok = EQEQ
				lit = "=="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '+':
			s.next()
			switch s.peek() {
			case '+':
				tok = PLUSPLUS
				lit = "++"
			case '=':
				tok = PLUSEQ
				lit = "+="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '-':
			s.next()
			switch s.peek() {
			case '-':
				tok = MINUSMINUS
				lit = "--"
			case '=':
				tok = MINUSEQ
				lit = "-="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '*':
			s.next()
			switch s.peek() {
			case '*':
				tok = POW
				lit = "**"
			case '=':
				tok = MULEQ
				lit = "*="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '/':
			s.next()
			switch s.peek() {
			case '=':
				tok = DIVEQ
				lit = "/="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '>':
			s.next()
			switch s.peek() {
			case '=':
				tok = GE
				lit = ">="
			case '>':
				tok = SHIFTRIGHT
				lit = ">>"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '<':
			s.next()
			switch s.peek() {
			case '=':
				tok = LE
				lit = "<="
			case '<':
				tok = SHIFTLEFT
				lit = "<<"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '|':
			s.next()
			switch s.peek() {
			case '|':
				tok = OROR
				lit = "||"
			case '=':
				tok = OREQ
				lit = "|="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '&':
			s.next()
			switch s.peek() {
			case '&':
				tok = ANDAND
				lit = "&&"
			case '=':
				tok = ANDEQ
				lit = "&="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '.':
			s.next()
			if s.peek() == '.' {
				s.next()
				if s.peek() == '.' {
					tok = VARARG
				} else {
					err = fmt.Errorf(`syntax error "%s"`, "..")
					return
				}
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '(', ')', ':', ';', '%', '?', '{', '}', ',', '[', ']', '\n', '^':
			tok = int(ch)
			lit = string(ch)
		default:
			err = fmt.Errorf(`syntax error "%s"`, string(ch))
			return
		}
		s.next()
	}
	return
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isHex(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

func isBrank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (s *Scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		return -1
	}
}

func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

func (s *Scanner) current() int {
	return s.offset
}

func (s *Scanner) set(o int) {
	s.offset = o
}

func (s *Scanner) back() {
	s.offset--
}

func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

func (s *Scanner) pos() ast.Position {
	return ast.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *Scanner) skipBlank() {
	for isBrank(s.peek()) {
		s.next()
	}
}

func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for {
		if !isLetter(s.peek()) && !isDigit(s.peek()) {
			break
		}
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

func (s *Scanner) scanNumber() (string, error) {
	var ret []rune
	ch := s.peek()
	ret = append(ret, ch)
	s.next()
	if ch == '0' && s.peek() == 'x' {
		ret = append(ret, s.peek())
		s.next()
		for isHex(s.peek()) {
			ret = append(ret, s.peek())
			s.next()
		}
	} else {
		for isDigit(s.peek()) || s.peek() == '.' || s.peek() == 'e' {
			ret = append(ret, s.peek())
			s.next()
		}
		if isLetter(s.peek()) {
			return "", errors.New("identifier starts immediately after numeric literal")
		}
	}
	return string(ret), nil
}

func (s *Scanner) scanRawString() (string, error) {
	var ret []rune
	for {
		s.next()
		if s.peek() == EOF {
			return "", errors.New("unexpected EOF")
			break
		}
		if s.peek() == '`' {
			s.next()
			break
		}
		ret = append(ret, s.peek())
	}
	return string(ret), nil
}

func (s *Scanner) scanString(l rune) (string, error) {
	var ret []rune
eos:
	for {
		s.next()
		switch s.peek() {
		case EOL:
			return "", errors.New("unexpected EOL")
		case EOF:
			return "", errors.New("unexpected EOF")
		case l:
			s.next()
			break eos
		case '\\':
			s.next()
			switch s.peek() {
			case 'b':
				ret = append(ret, '\b')
				continue
			case 'f':
				ret = append(ret, '\f')
				continue
			case 'r':
				ret = append(ret, '\r')
				continue
			case 'n':
				ret = append(ret, '\n')
				continue
			case 't':
				ret = append(ret, '\t')
				continue
			}
			ret = append(ret, s.peek())
			continue
		default:
			ret = append(ret, s.peek())
		}
	}
	return string(ret), nil
}

type Lexer struct {
	s     *Scanner
	lit   string
	pos   ast.Position
	e     error
	stmts []ast.Stmt
}

type Error struct {
	Message  string
	Pos      ast.Position
	Filename string
	Fatal    bool
}

func (e *Error) Error() string {
	return e.Message
}

func (l *Lexer) Error(msg string) {
	l.e = &Error{Message: msg, Pos: l.pos, Fatal: false}
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos, err := l.s.Scan()
	if err != nil {
		l.e = &Error{Message: fmt.Sprintf("%s", err.Error()), Pos: pos, Fatal: true}
	}
	if tok == EOF {
		return 0
	}
	lval.tok = ast.Token{Tok: tok, Lit: lit}
	lval.tok.SetPosition(pos)
	l.lit = lit
	l.pos = pos
	return tok
}

func Parse(s *Scanner) ([]ast.Stmt, error) {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		return nil, l.e
	}
	return l.stmts, l.e
}
