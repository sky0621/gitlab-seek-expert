# リポジトリ一覧

{{range .}}## {{.Path}}

| No | Project Name | Description | Web URL | Last Activity At | Commit Count | Commit Users |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .Projects}}| {{.No}} | {{.Name}} | {{range .Descriptions}}{{.}}<br>{{end}} | {{.WebURL}} | {{.LastActivityAt}} | {{.CommitCount}} | {{range .Committers}}{{.CommitterName}}({{.CommitterEmail}}):{{.CommitCount}}<br>{{end}} |
{{end}}
{{end}}
