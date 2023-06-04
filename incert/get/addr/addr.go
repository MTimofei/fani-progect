package addr

import (
	"fmt"
	"log"
	"net/http"
)

type URL struct {
	url string
	//storeg *data.Data
}

func New(url string) *URL {
	return &URL{url: url}
}

func (addr *URL) Start() {
	// template.
	var result []byte
	var client = http.Client{}

	req, err := http.NewRequest(http.MethodGet, addr.url, nil)
	if err != nil {
		log.Println("err func get.start:", err)
		return
	}

	res, err := client.Do(req)

	if err != nil {
		log.Println("err func get.start:", err)
		return
	}
	// err = URL.storeg.Set()
	// if err != nil {
	// 	log.Println("err func get.start:", err)
	// 	return
	// }

	defer res.Body.Close()

	res.Body.Read(result)

	fmt.Println(res.Status)
	fmt.Println(res.Request.Proto)

	fmt.Println(res.Request.Form)
	fmt.Println(result)
}
