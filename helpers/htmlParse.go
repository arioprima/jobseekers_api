package helpers

import (
	"bytes"
	"fmt"
	"github.com/arioprima/jobseekers_api/schemas"
	"github.com/sirupsen/logrus"
	"html/template"
)

func ParseHtml(fileName string, data map[string]string) string {
	html, errParse := template.ParseFiles("templates/" + fileName + ".html")

	if errParse != nil {
		defer fmt.Println("parser file html failed")
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
		defer fmt.Println("execute html file failed")
		logrus.Fatal(errExecute.Error())
	}

	return buf.String()
}
