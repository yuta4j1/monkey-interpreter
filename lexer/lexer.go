package lexer

import "github.com/yuta4j1/monkey-interpreter/token"

// 字句解析器
type Lexer struct {
	input        string // 対象となる文字列
	position     int    // 現在の文字へのカーソル位置
	readPosition int    // 次の文字へのカーソル位置
	ch           byte   // 読み取った文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// カーソルを最初の文字に設定
	l.readChar()
	return l
}

// 一文字読む。readPositionは次の文字に設定する
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// 空白をスキップする
// 現在カーソルの文字が空白文字でなくなるまで、読み取りを進める
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// 空白文字は字句解析の対象にならない
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			// 次の文字が '=' ならば、"==" リテラルとして識別する
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// 識別子を読み、非英字に到達するまで字句解析の位置を進める
func (l *Lexer) readIdentifier() string {
	// 現在のカーソル位置
	position := l.position
	for isLetter(l.ch) {
		// 英字でなくなるまで、文字を読み続ける
		l.readChar()
	}
	return l.input[position:l.position]
}

// 識別子（数値）を読みこむ
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		// 数値でなくなるまで、文字を読み続ける
		l.readChar()
	}
	return l.input[position:l.position]
}

// 文字の先読みを行う。
// 正しく字句解析するには、'=' と '==' の判別など、先読みする必要がある。
// この関数では一文字先を先読みする（Lexerオブジェクトのpositionは変わらない。）
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// トークンを作成する
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 文字判定
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 数値判定
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
