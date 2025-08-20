package template

import "html/template"

// ITemplateService  ...
type ITemplateService interface {

}

// TemplateService ...
type TemplateService struct {
	Template *template.Template
}


// GenerateInspectionTemplate ... returning  