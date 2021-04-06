package openacademy

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/pool/h"
)

var fields_OpenAcademyCourse = map[string]models.FieldDefinition{
	"Name":        fields.Char{},
	"Description": fields.Text{},
}

func init() {
	models.NewModel("OpenAcademyCourse")
	h.OpenAcademyCourse().AddFields(fields_OpenAcademyCourse)
}
