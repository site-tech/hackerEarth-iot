package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	_ "github.com/quasilyte/go-ruleguard/dsl"
)

func main() {
	// url and REST method to send data to
	url := "http://localhost:49565/api/v1/resource/RESTDevice/int16_reading"
	method := "POST"
	// buld REST client
	client := &http.Client{}

	for i := 0; i < 10; i++ {

		// generate random data
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 999
		value := rand.Intn(max-min+1) + min
		payload := strings.NewReader(fmt.Sprint(value))

		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			fmt.Println(err)
			return
		}

		// add content header
		req.Header.Add("Content-Type", "text/plain")

		// perform resquest and check for error
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		// read in response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		// print status code
		fmt.Println(res.StatusCode)
		// print response body
		fmt.Println(string(body))
	}
}
