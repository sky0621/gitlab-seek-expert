# リポジトリ一覧

{{range .}}## {{.Path}}

| No | Project Name | Description | Web URL | Last Activity At | Commit Count | Commit Users |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .Projects}}| {{.No}} | {{.Name}} | {{range .Descriptions}} {{.}}<br> {{end}} | {{.WebURL}} | {{.LastActivityAt}} | {{.CommitCount}} | {{.Committers}} |
{{end}}
{{end}}
