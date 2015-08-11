package main
import (
	"net/http"
	"io"
	"fmt"
)

func main() {
	resp, _ := newRequest("POST", "http://localhost:8080/create_post", nil)
	fmt.Println(resp)
}

func newRequest(method, urlStr string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, urlStr, body)
	return http.DefaultClient.Do(req)
}