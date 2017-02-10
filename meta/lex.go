package meta

import "github.com/davyxu/golexer"

// 自定义的token id
const (
	Token_EOF = iota
	Token_Unknown
	Token_LineEnd
	Token_Numeral
	Token_String
	Token_WhiteSpace
	Token_Identifier
	Token_Colon       // :
	Token_ParenL      // (
	Token_ParenR      // )
	Token_CurlyBraceL // {
	Token_CurlyBraceR // }
	Token_Star        // *
	Token_Dot         // .
)

type sprotoParser struct {
	*golexer.Parser
}

func newSProtoParser() *sprotoParser {

	l := golexer.NewLexer()

	// 匹配顺序从高到低

	l.AddMatcher(golexer.NewNumeralMatcher(Token_Numeral))
	l.AddMatcher(golexer.NewStringMatcher(Token_String))

	l.AddIgnoreMatcher(golexer.NewWhiteSpaceMatcher(Token_WhiteSpace))
	l.AddIgnoreMatcher(golexer.NewLineEndMatcher(Token_LineEnd))

	l.AddMatcher(golexer.NewSignMatcher(Token_CurlyBraceL, "{"))
	l.AddMatcher(golexer.NewSignMatcher(Token_CurlyBraceR, "}"))
	l.AddMatcher(golexer.NewSignMatcher(Token_ParenL, "("))
	l.AddMatcher(golexer.NewSignMatcher(Token_ParenR, ")"))
	l.AddMatcher(golexer.NewSignMatcher(Token_Star, "*"))
	l.AddMatcher(golexer.NewSignMatcher(Token_Dot, "."))
	l.AddMatcher(golexer.NewSignMatcher(Token_Colon, ":"))

	l.AddMatcher(golexer.NewIdentifierMatcher(Token_Identifier))

	l.AddMatcher(golexer.NewUnknownMatcher(Token_Unknown))

	return &sprotoParser{
		Parser: golexer.NewParser(l),
	}
}
