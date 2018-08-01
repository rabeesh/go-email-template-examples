package main

import (
	"bytes"
	"fmt"
	"html/template"
)

type TplParser struct {
	Tpls map[string]*template.Template
}

func NewTplParser(tplNames []string) *TplParser {
	es := &TplParser{Tpls: make(map[string]*template.Template, 0)}

	for _, tplName := range tplNames {
		t, err := template.ParseFiles(
			"./tpls/emails/layout.html",
			"./tpls/emails/orderItems.html",
			"./tpls/emails/"+tplName+".html",
		)
		if err != nil {
			fmt.Println("Error in loading email templates")
			panic(err)
		}
		es.Tpls[tplName] = t
	}
	return es
}

func (es *TplParser) Parse(emailType string, data Order) (string, error) {
	t := es.Tpls[emailType]
	var tplBytes bytes.Buffer

	if err := t.ExecuteTemplate(&tplBytes, "layout", data); err != nil {
		return "", err
	}

	return tplBytes.String(), nil
}
