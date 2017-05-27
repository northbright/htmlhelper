package htmlhelper

import (
	//"fmt"
	"regexp"
	"strings"
)

type Row struct {
	Columns []string
}

type Table struct {
	Rows []Row
}

// Clean removes white spaces, blank lines and line breakers to clean input HTML string.
func Clean(input string) (output string) {
	var p = ``
	var re *regexp.Regexp

	output = input

	// Remove all '\r'.
	output = strings.Replace(output, "\r", "", -1)

	// Remove white spaces and line breakers after '<xx '.
	p = `<(\S+)\s+`
	re = regexp.MustCompile(p)
	output = re.ReplaceAllString(output, "<$1 ")

	// Remove white spaces and link breakers before '>'.
	p = `\s+>`
	re = regexp.MustCompile(p)
	output = re.ReplaceAllString(output, ">")

	// Remove blank lines.
	p = `\s*\n\s*`
	re = regexp.MustCompile(p)
	output = re.ReplaceAllString(output, "")

	// Remove white spaces and line breaks after '>'.
	p = `>\s+`
	re = regexp.MustCompile(p)
	output = re.ReplaceAllString(output, ">")

	// Remove white spaces and line breaks before '<'.
	p = `\s+<`
	re = regexp.MustCompile(p)
	output = re.ReplaceAllString(output, "<")

	return output
}

/*
func getTable(htmlStr string) (t Table) {
}

func GetTables(htmlStr string) (tables []Table) {
	var p = `<table.*`
	var re *regexp.Regexp

	str := Clean(htmlStr)

}
*/
