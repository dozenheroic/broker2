package sdk

import (
	"fmt"
	"io"
	"net/http"
)

func Publish(topic, msg string) {
	url := fmt.Sprintf("http://localhost:8080/publish?topic=%s&msg=%s", topic, msg)
	_, _ = http.Get(url)
}

func Poll(topic, group string) string {
	url := fmt.Sprintf("http://localhost:8080/poll?topic=%s&group=%s", topic, group)
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
