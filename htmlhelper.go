package htmlhelper

import (
	"regexp"
	"strings"
)

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

	// Make sure all table tags are lower cases
	tagsMap := map[string]string{
		"TABLE": "table",
		"TH":    "th",
		"TR":    "tr",
		"TD":    "td",
		"TBODY": "tbody",
	}
	for k, v := range tagsMap {
		output = strings.Replace(output, k, v, -1)
	}

	return output
}

func getColumns(htmlStr string, isHead bool) (columns []string) {
	p := ``

	if isHead {
		p = `<th.*?>(.*?)</th>`
	} else {
		p = `<td.*?>(.*?)</td>`
	}

	re := regexp.MustCompile(p)
	matched := re.FindAllStringSubmatch(htmlStr, -1)

	for _, m := range matched {
		columns = append(columns, m[1])
	}

	return columns
}

func getRows(htmlStr string) (rows []string) {
	p := `<tr.*?>(.*?)</tr>`
	re := regexp.MustCompile(p)
	matched := re.FindAllStringSubmatch(htmlStr, -1)

	for _, m := range matched {
		rows = append(rows, m[1])
	}

	return rows
}

func getTableHead(htmlStr string) (headStr string) {
	p := `<thead.*?>(.*?)</thead>`
	re := regexp.MustCompile(p)
	matched := re.FindStringSubmatch(htmlStr)
	if len(matched) != 2 {
		return ""
	}
	return matched[1]
}

func getTableBody(htmlStr string) (bodyStr string) {
	p := `<tbody.*?>(.*?)</tbody>`
	re := regexp.MustCompile(p)
	matched := re.FindStringSubmatch(htmlStr)
	if len(matched) != 2 {
		return ""
	}
	return matched[1]
}

func getTables(htmlStr string) (tables []string) {
	p := `<table.*?>(.*?)</table>`
	re := regexp.MustCompile(p)
	matched := re.FindAllStringSubmatch(htmlStr, -1)

	for _, m := range matched {
		tables = append(tables, m[1])
	}

	return tables
}

// TablesToCSVs finds all tables in HTML string and converts each table to CSV records.
//
// Params:
//     htmlStr: input HTML string.
// Return:
//     csvs: slice contains CSV records.
//           See https://godoc.org/encoding/csv for more info.
func TablesToCSVs(htmlStr string) (csvs [][][]string) {
	str := Clean(htmlStr)

	for _, t := range getTables(str) {
		records := [][]string{}

		head := getTableHead(t)
		body := ""

		// Table has <thead> and <tbody>.
		if len(head) != 0 {
			for _, r := range getRows(head) {
				records = append(records, getColumns(r, true))
			}
			body = getTableBody(t)
		} else {
			body = t
		}

		for _, r := range getRows(body) {
			records = append(records, getColumns(r, false))
		}

		csvs = append(csvs, records)
	}

	return csvs
}
