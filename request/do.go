package request

import (
	"fmt"
	"io"
	"net/http"
)

func Get(address string) (string, error) {
	resp, err := http.DefaultClient.Get(address)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("发生了 Get 请求")
	return string(bytesBody), nil
}
