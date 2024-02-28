package api

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

func SendFormDataRequest(url string, start string, end string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("numStart", start)
	_ = writer.WriteField("numEnd", end)

	writer.Close()
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	responce, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer responce.Body.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(responce.Body)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil

}
