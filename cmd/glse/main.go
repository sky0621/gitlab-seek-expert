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

	namespaces, res, err := git.Namespaces.ListNamespaces(&gitlab.ListNamespacesOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 1000,
		},
	})
	if err != nil {
		panic(err)
	}
	if res.Status != "200 OK" {
		panic(errors.New("not 200 OK"))
	}

	glNss := []*gitlabNameSpace{}
	for _, ns := range namespaces {
		projects, res, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{
			OrderBy: gitlab.String("name"),
			Sort:    gitlab.String("asc"),
			ListOptions: gitlab.ListOptions{
				Page:    1,
				PerPage: 1000,
			},
		})
		if err != nil {
			panic(err)
		}
		if res.Status != "200 OK" {
			panic(errors.New("not 200 OK"))
		}

		glPrjs := []*gitlabProject{}
		no := 1
		for _, p := range projects {
			if ns.Path != p.Namespace.Path {
				continue
			}

			commits, res, err := git.Commits.ListCommits(p.ID, &gitlab.ListCommitsOptions{})
			if err != nil {
				panic(err)
			}
			if res.Status != "200 OK" {
				panic(errors.New("not 200 OK"))
			}

			cmtMap := map[string]int{}
			for _, cmt := range commits {
				if cnt, ok := cmtMap[cmt.AuthorName]; ok {
					cmtMap[cmt.AuthorName] = cnt + 1
				} else {
					cmtMap[cmt.AuthorName] = 1
				}
			}

			glPrjs = append(glPrjs, &gitlabProject{
				No:             no,
				ID:             p.ID,
				NamespaceID:    p.Namespace.ID,
				Namespace:      p.Namespace.Name,
				Name:           p.Name,
				Description:    p.Description,
				Owner:          getOwnerName(p.Owner),
				WebURL:         p.WebURL,
				LastActivityAt: p.LastActivityAt.Format("2006-01-02 15:04:05"),
				CommitCount:    len(commits),
				CommitUsers:    fmt.Sprintf("%v", cmtMap),
			})
			no = no + 1
		}

		glNss = append(glNss, &gitlabNameSpace{Path: ns.Path, Projects: glPrjs})
	}

	tmpl := template.Must(template.ParseFiles("./tmpl.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, glNss)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}

type gitlabNameSpace struct {
	Path     string
	Projects []*gitlabProject
}

type gitlabProject struct {
	No             int
	ID             int
	NamespaceID    int
	Namespace      string
	Name           string
	Description    string
	Owner          string
	WebURL         string
	LastActivityAt string
	CommitCount    int
	CommitUsers    string
}

func getOwnerName(o *gitlab.User) string {
	if o == nil {
		return "-"
	}
	return o.Name
}
