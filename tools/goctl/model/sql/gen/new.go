package gen

import (
	"github.com/sjclijie/go-zero/tools/goctl/model/sql/template"
	"github.com/sjclijie/go-zero/tools/goctl/util"
	"github.com/sjclijie/go-zero/tools/goctl/util/stringx"
)

func genNew(table Table, withCache bool) (string, error) {
	text, err := util.LoadTemplate(category, modelNewTemplateFile, template.New)
	if err != nil {
		return "", err
	}

	output, err := util.With("new").
		Parse(text).
		Execute(map[string]interface{}{
			"table":                 wrapWithRawString(table.Name.Source()),
			"withCache":             withCache,
			"upperStartCamelObject": table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
