package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	client := retryablehttp.NewClient()
	client.RetryMax = 5
	client.RetryWaitMin = time.Second * 3
	client.HTTPClient.Timeout = 3 * time.Second

	req, err := retryablehttp.NewRequest("PUT", "http://localhost:5000/hello", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	result, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	fmt.Println(string(result))

}
