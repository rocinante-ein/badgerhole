package presenters

import (
	"html/template"
	"net/http"
	"strings"
)

// FakeLoginPresenterImpl BadgerHole login page presenter
type FakeLoginPresenterImpl struct {
	TemplateFile string
}

// NewFakeLoginPresenter create fake login presenter instance
func NewFakeLoginPresenter(templateFile string) *FakeLoginPresenterImpl {
	return &FakeLoginPresenterImpl{
		TemplateFile: templateFile,
	}
}

// Response rendering fake login page
func (flp FakeLoginPresenterImpl) Response(w http.ResponseWriter, params map[string]string) {

	// fail message area rendering word
	var failMessage = ""

	// if http method is "POST" ... authentication failure
	if strings.ToUpper(params["http_method"]) == "POST" {
		failMessage = "authentication failure"
	}

	// set fake login template file
	tmpl, err := template.ParseFiles(flp.TemplateFile)
	if err != nil {
		panic(err)
	}

	// fake servername
	w.Header().Set("server", params["server"])

	// response fake login page
	err = tmpl.Execute(w, map[string]string{"message": failMessage})
	if err != nil {
		panic(err)
	}
}
