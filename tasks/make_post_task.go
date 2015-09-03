package main
import (
	"net/http"
	"io"
	"fmt"
	"github.com/heckfer/fala-com-meu-carro/model"
)

func main() {
	post := &model.Post{
		CarPlate: "IWI-5585",
		Message: "I saw a credit card falling from this car! Thats crazy",
	}

	resp, _ := newRequest("POST", "http://localhost:8080/create_post", post)

	fmt.Println(resp)
}

func newRequest(method, urlStr string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, urlStr, body)
	return http.DefaultClient.Do(req)
}