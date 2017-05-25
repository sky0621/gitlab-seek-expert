package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"

	"errors"

	gitlab "github.com/xanzy/go-gitlab"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	host := flag.String("host", "example.com", "GitLab Host Name")
	pkey := flag.String("pkey", "xxxxxxxxxx", "Your GitLab Private Key")
	flag.Parse()
	git := gitlab.NewClient(nil, *pkey)
	git.SetBaseURL(fmt.Sprintf("%s/api/v3", *host))
	ob := "name"
	st := "asc"
	projects, res, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{
		OrderBy: &ob,
		Sort: &st,
		ListOptions: gitlab.ListOptions{
			Page: 1,
			PerPage: 1000,
		},
	})
	if err != nil {
		panic(err)
	}
	if res.Status != "200 OK" {
		panic(errors.New("not 200 OK"))
	}

	infos := []*gitlabInfo{}
	for idx, p := range projects {
		infos = append(infos, &gitlabInfo{
			No:          idx+1,
			ID:          p.ID,
			NamespaceID: p.Namespace.ID,
			Namespace:   p.Namespace.Name,
			Name:        p.Name,
			Description: p.Description,
		})
	}

	tmpl := template.Must(template.ParseFiles("./tmpl.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, infos)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}

type gitlabInfo struct {
	No	int
	ID          int
	NamespaceID int
	Namespace   string
	Name        string
	Description string
}
