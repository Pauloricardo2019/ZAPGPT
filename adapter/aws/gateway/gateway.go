package gateway

import (
	"encoding/base64"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"net/url"
	"zapgpt/adapter/provider"
)

func Process(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, err := parseBase64RequestData(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	text, err := provider.GenerateGPTText(result)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       text,
	}, nil
}

// O twillio sempre manda o Body como Query's, por exemplo "x=0&Body=Test&y=b", por isso usamos a func "url.ParseQuery"
func parseBase64RequestData(req string) (string, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(req)
	if err != nil {
		return "", err
	}

	data, err := url.ParseQuery(string(dataBytes))
	if err != nil {
		return "", err
	}

	if data.Has("Body") {
		return data.Get("Body"), nil
	}

	return "", errors.New("Body not found")
}
