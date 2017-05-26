# リポジトリ一覧

{{range .}}## {{.Path}}

| No | Project Name | Description | Web URL | Last Activity At | Commit Count | Commit Users |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .Projects}}| {{.No}} | {{.Name}} | {{.Description}} | {{.WebURL}} | {{.LastActivityAt}} | {{.CommitCount}} | {{.CommitUsers}} |
{{end}}
{{end}}
