package forms

// This will contain all the helper functions for the controllers

import (
	// "encoding/json"
	"fmt"
	pongo2 "github.com/flosch/pongo2/v4"
	"github.com/yuin/gluamapper"
	// "bitbucket.org/selenesoftware/humboldt/template"
	// "github.com/yuin/gopher-lua"
)

type tagFormStartNode struct {
	position *pongo2.Token
	name     string
}

func (node *tagFormStartNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	// fmt.Println(ctx)
	writer.WriteString("<form>JASON HAS A MUTHAFUCKIN FORM</form>")
	return nil
}

func formStart(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	formStartNode := &tagFormStartNode{
		position: start,
		name:     "form",
	}

	// var form Form
	formatToken := arguments.MatchType(pongo2.TokenString)
	// fmt.Println(arguments)
	if err := gluamapper.Map(formatToken, &formStartNode); err != nil {
		panic(err)
	}

	fmt.Println(arguments)

	// if countToken := arguments.MatchType(TokenNumber); countToken != nil {
	// 	loremNode.count = AsValue(countToken.Val).Integer()
	// }

	return formStartNode, nil
}

func init() {
	pongo2.RegisterTag("formstart", formStart)
}
