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
	url := "http://localhost:49565/api/v1/resource/RESTDevice/float32_reading"
	method := "POST"
	// buld REST client
	client := &http.Client{}

	for {
		// generate random data
		time.Sleep(time.Second)
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 40
		value := rand.Intn(max-min+1) + min
		fmt.Println(value)
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
		_, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
