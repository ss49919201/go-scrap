package ses

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func newSession(region string) *session.Session {
	return session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	))
}

func GetTemplate(region, templateName string) (*ses.Template, error) {
	sess := newSession(region)
	sesClient := ses.New(sess)

	req, resp := sesClient.GetTemplateRequest(&ses.GetTemplateInput{
		TemplateName: aws.String(templateName),
	})

	err := req.Send()
	if err != nil {
		return nil, err
	}

	return resp.Template, nil
}

func SendTemplatedEmail(region, source, templateName string, toAddresses []string, data interface{}) error {
	sess := newSession(region)
	sesClient := ses.New(sess)

	dataJson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, _ := sesClient.SendTemplatedEmailRequest(
		&ses.SendTemplatedEmailInput{
			Destination: &ses.Destination{
				ToAddresses: aws.StringSlice(toAddresses),
			},
			Source:       aws.String(source),
			Template:     aws.String(templateName),
			TemplateData: aws.String(string(dataJson)),
		},
	)

	return req.Send()
}
