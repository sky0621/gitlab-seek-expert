package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"

	"errors"

	gitlab "github.com/xanzy/go-gitlab"
)

const (
	perPage = 99999
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
			PerPage: perPage,
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
				PerPage: perPage,
			},
		})
		if err != nil {
			panic(err)
		}
		if res.Status != "200 OK" {
			panic(errors.New("not 200 OK"))
		}

		if len(projects) == 0 {
			continue
		}

		glPrjs := []*gitlabProject{}
		no := 1
		for _, p := range projects {
			if ns.Path != p.Namespace.Path {
				continue
			}

			commits, res, err := git.Commits.ListCommits(p.ID, &gitlab.ListCommitsOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: perPage,
				},
			})
			if err != nil {
				panic(err)
			}
			if res.Status != "200 OK" {
				panic(errors.New("not 200 OK"))
			}

			cmtMap := map[string]*gitlabCommitter{}
			for _, cmt := range commits {
				setCnt := 1
				if committer, ok := cmtMap[cmt.CommitterEmail]; ok {
					setCnt = committer.CommitCount + 1
				}
				cmtMap[cmt.CommitterEmail] = &gitlabCommitter{
					CommitterEmail: cmt.CommitterEmail,
					CommitterName:  cmt.CommitterName,
					CommitCount:    setCnt,
				}
			}

			committers, allCnt := toSlice(cmtMap)
			glPrjs = append(glPrjs, &gitlabProject{
				No:             no,
				ID:             p.ID,
				NamespaceID:    p.Namespace.ID,
				Namespace:      p.Namespace.Name,
				Name:           p.Name,
				Description:    toMdDescription(p.Description),
				WebURL:         p.WebURL,
				LastActivityAt: p.LastActivityAt.Format("2006-01-02 15:04:05"),
				CommitCount:    allCnt,
				Committers:     committers,
			})
			no = no + 1
		}

		if len(glPrjs) == 0 {
			continue
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
	WebURL         string
	LastActivityAt string
	CommitCount    int
	Committers     []*gitlabCommitter
}

type gitlabCommitter struct {
	CommitterEmail string
	CommitterName  string
	CommitCount    int
}

func toMdDescription(d string) string {
	return d
}

func toSlice(cmtMap map[string]*gitlabCommitter) ([]*gitlabCommitter, int) {
	committers := []*gitlabCommitter{}
	allCnt := 0
	for _, v := range cmtMap {
		committers = append(committers, v)
		allCnt = allCnt + v.CommitCount
	}
	return committers, allCnt
}
