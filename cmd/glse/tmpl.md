# リポジトリ一覧

{{range .}}## {{.Path}}

| No | Project Name | Description | Owner | Web URL | Last Activity At |
| :--- | :--- | :--- | :--- | :--- | :--- |
{{range .Projects}}| {{.No}} | {{.Name}} | {{.Description}} | {{.Owner}} | {{.WebURL}} | {{.LastActivityAt}} |
{{end}}
{{end}}
