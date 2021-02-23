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
	standardClient := client.StandardClient()

	resp, err := standardClient.Get("http://localhost:5000/hello")

	if err != nil {
		panic(err)
	}

	result, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	fmt.Println(string(result))

}
