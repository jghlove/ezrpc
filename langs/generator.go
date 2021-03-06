package langs

import "github.com/Wuvist/go-thrift/parser"

type ApiGen interface {
	Generate(output string, parsedThrift map[string]*parser.Thrift)
}

var Langs = make(map[string]ApiGen)
