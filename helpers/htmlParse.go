package helpers

import (
	"bytes"
	"github.com/arioprima/jobseekers_api/schemas"
	"github.com/arioprima/jobseekers_api/templates"
	"github.com/sirupsen/logrus"
	"html/template"
)

func ParseHtml(fileName string, data map[string]string) string {
	// Sesuaikan path jika perlu
	html, errParse := template.ParseFS(templates.TemplateFS, fileName+".html")
	if errParse != nil {
		logrus.Fatal(errParse.Error())
	}

	body := schemas.SchemaHTMLRequest{
		To:   data["to"],
		Otp:  data["otp"],
		Date: data["date"],
		Year: data["year"],
	}

	buf := new(bytes.Buffer)
	errExecute := html.Execute(buf, body)
	if errExecute != nil {
		logrus.Fatal(errExecute.Error())
	}

	return buf.String()
}
