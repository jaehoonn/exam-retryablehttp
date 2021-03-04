package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/NexClipper/logger"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/jaehoonn/exam-retryablehttp/common"
)

func main() {
	loggerEnv := &common.LoggerEnv{
		Level:      "debug",
		LogPath:    "./log/klevr.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   false,
	}

	common.InitLogger(loggerEnv)
	logger.Debug("start")
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.RequestLogHook = func(l retryablehttp.Logger, req *http.Request, n int) {
		logger.Debugf("[%s] %s", req.Method, req.URL.String())
	}

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
