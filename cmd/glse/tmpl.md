# リポジトリ一覧

{{range .}}## {{.Path}}

| No | Avatar | Project Name | Description | Web URL | Last Activity At | Commit Count | Commit Users |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .Projects}}| {{.No}} | ![No Image]({{.AvatarURL}}) | {{.Name}} | {{range .Descriptions}}{{.}}<br>{{end}} | {{.WebURL}} | {{.LastActivityAt}} | {{.CommitCount}} | {{range .Committers}}{{.CommitterName}}({{.CommitterEmail}}): {{.CommitCount}}<br>{{end}} |
{{end}}
{{end}}
