package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

type IssueItem struct {
	ID       int
	User     string
	Title    string
	CreateAt time.Time
}

type Issues struct {
	Count  int
	Issues []IssueItem
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func TestTextTemplate() {
	issues := Issues{
		Count: 2,
		Issues: []IssueItem{
			IssueItem{1, "foo", "title-foo", time.Now().AddDate(-1, 0, 0)},
			IssueItem{2, "bar", "title-bar", time.Now().AddDate(0, -1, 0)},
		},
	}

	const templ = `{{.Count}} issues:
{{range .Issues}}------------------
ID: {{.ID}}
User: {{.User}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, issues); err != nil {
		log.Fatal(err)
	}
}

func main() {
	TestTextTemplate()
}
