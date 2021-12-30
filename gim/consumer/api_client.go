package consumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type ApiClient struct {
	forwardUrl string
	httpClient http.Client
}

func NewApiClient(forwardUrl string, timeout time.Duration) ApiClient {
	client := http.Client{Timeout: timeout}
	return ApiClient{
		forwardUrl: forwardUrl,
		httpClient: client,
	}
}

func (client ApiClient) Send(msg IncomingMessage) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, client.forwardUrl, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	request.Header.Add("X-User-Agent", "GIM")

	response, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Message processing failed with response status code %d", response.StatusCode))
	}
	return nil
}
