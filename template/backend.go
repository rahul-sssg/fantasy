package template

import (
	"html/template"
	"log"
	"os"
)

// declaring struct
type Student struct {
	// defining struct fields
	Name  string
	Marks int
	Id    string
}

// main function
func GenerateHtml() error {
	// defining struct instance
	std1 := Student{
		Name:  "Rahul",
		Marks: 100,
		Id:    "999",
	}

	// Parsing the required html
	// file in same directory
	t, err := template.ParseFiles("template/indexTemplateInput.html")

	if err != nil {
		return nil
	}
	// standard output to print merged data
	err = t.Execute(os.Stdout, std1)
	if err != nil {
		return nil
	}
	
	f, err := os.Create("./template/indexTemplateOutput.html")
	if err != nil {
		log.Println("create file: ", err)
		return err
	}
	err = t.Execute(f, std1)
	if err != nil {
		log.Print("execute: ", err)
		return err
	}
	f.Close()
	return nil

}
