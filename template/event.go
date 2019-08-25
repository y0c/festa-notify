package template

import (
	"bytes"
	"fmt"
	"github.com/y0c/festa-notify/festa"
	"html/template"
)

// Variables is HTML template variable
type Variables struct {
	Events []festa.Event
}

const mailTemplate = `
	<html>
		<head>
			<meta charset="utf-8" />
		</head>
		<body>
			<ul>
			 {{range .Events}}
					<li>
						<h3><a href="https://festa.io/events/{{ .EventID }}">{{ .Name }}</a></h3>
						<ul>
							{{range .Tickets}}
								<li>
									<strong>{{ .Name }}</strong>
									<div>
										<span>마감여부 - </span><span>{{ if .Registable }} 등록가능 {{ else }} 마감 {{ end }}<span>
									</div>
									<div>
										<span>가격 - </span><span>{{ .Price }}원<span>
									</div>
									<div>
										<span>{{ .LimitPerUser }} 개 남음</span>
									</div>
								</li>
							{{end}}
						</ul>
					</li>		
			 {{end}}
			</ul>
		</body>
	</html>
`

// GenerateEventTemplate convert events to mail html template
func GenerateEventTemplate(events []festa.Event) (string, error) {

	var buf bytes.Buffer
	t := template.New("MailTemplate")

	tmpl, err := t.Parse(mailTemplate)

	if err != nil {
		return "", fmt.Errorf("template error %v", err)
	}

	err = tmpl.Execute(&buf, Variables{
		events,
	})

	if err != nil {
		return "", fmt.Errorf("template error %v", err)
	}

	return buf.String(), nil
}
