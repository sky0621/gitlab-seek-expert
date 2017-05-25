# リポジトリ一覧

{{range .}}## {{.Path}}

| No | Project Name | Description | Owner | Web URL | Last Activity At | Commit Count | Commit Users |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .Projects}}| {{.No}} | {{.Name}} | {{.Description}} | {{.Owner}} | {{.WebURL}} | {{.LastActivityAt}} | {{.CommitCount}} | {{.CommitUsers}} |
{{end}}
{{end}}
