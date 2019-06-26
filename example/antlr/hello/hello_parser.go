// Code generated from Hello.g4 by ANTLR 4.7.2. DO NOT EDIT.

package hello // Hello
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 5, 8, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 6, 2, 4, 3, 2, 2,
	2, 4, 5, 7, 3, 2, 2, 5, 6, 7, 4, 2, 2, 6, 3, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'hello'",
}
var symbolicNames = []string{
	"", "", "ID", "WS",
}

var ruleNames = []string{
	"r",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type HelloParser struct {
	*antlr.BaseParser
}

func NewHelloParser(input antlr.TokenStream) *HelloParser {
	this := new(HelloParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Hello.g4"

	return this
}

// HelloParser tokens.
const (
	HelloParserEOF  = antlr.TokenEOF
	HelloParserT__0 = 1
	HelloParserID   = 2
	HelloParserWS   = 3
)

// HelloParserRULE_r is the HelloParser rule.
const HelloParserRULE_r = 0

// IRContext is an interface to support dynamic dispatch.
type IRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRContext differentiates from other interfaces.
	IsRContext()
}

type RContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRContext() *RContext {
	var p = new(RContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HelloParserRULE_r
	return p
}

func (*RContext) IsRContext() {}

func NewRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RContext {
	var p = new(RContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HelloParserRULE_r

	return p
}

func (s *RContext) GetParser() antlr.Parser { return s.parser }

func (s *RContext) ID() antlr.TerminalNode {
	return s.GetToken(HelloParserID, 0)
}

func (s *RContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.EnterR(s)
	}
}

func (s *RContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.ExitR(s)
	}
}

func (p *HelloParser) R() (localctx IRContext) {
	localctx = NewRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HelloParserRULE_r)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		p.Match(HelloParserT__0)
	}
	{
		p.SetState(3)
		p.Match(HelloParserID)
	}

	return localctx
}
