# リポジトリ一覧

{{range .}}## {{.Path}}

| No | Avatar | Project Name | Description | Last Activity At | Commit Users |
| :--- | :--- | :--- | :--- | :--- | :--- |
{{range .Projects}}| {{.No}} | <img src="{{.AvatarURL}}" alt="No Image" width="100"> | [{{.Name}}]({{.WebURL}}) | {{range .Descriptions}}{{.}}<br>{{end}} | {{.LastActivityAt}} | TotalCount: {{.CommitCount}}<br><br>{{range .Committers}}{{.CommitterName}}({{.CommitterEmail}}): {{.CommitCount}}<br>{{end}} |
{{end}}
{{end}}
