### TencentCloud CLS SDK ZAP

Zap logger with Tencent Cloud CLS support. This logger will export LogRecord's in Tencent Cloud CLS format.

### Quick start

```go

package main

import (
	"fmt"
	"time"

	clszap "github.com/farmerx/tencentcloud-cls-sdk-zap"
	tencentcloud_cls_sdk_go "github.com/tencentcloud/tencentcloud-cls-sdk-go"
	"go.uber.org/zap"
)

func main() {
	uploadLog()
}

func uploadLog() {
	producerConfig := tencentcloud_cls_sdk_go.GetDefaultAsyncProducerClientConfig()
	producerConfig.Endpoint = "ap-guangzhou.cls.tencentcs.com"
	producerConfig.AccessKeyID = "[AccessKey]"
	producerConfig.AccessKeySecret = "[SecretKey]"
	topicId := "[日志主题ID]"
	callBack := &Callback{}
	producerInstance, err := tencentcloud_cls_sdk_go.NewAsyncProducerClient(producerConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	producerInstance.Start()

	// create new  logger with tencentcloud cls zap core and set it globally
	logger := zap.New(clszap.NewCLSCore(topicId, callBack, producerInstance))
	zap.ReplaceGlobals(logger)

	logger.Info("123455")
	time.Sleep(1000 * time.Second)
}

type Callback struct {
}

func (callback *Callback) Success(result *tencentcloud_cls_sdk_go.Result) {
	attemptList := result.GetReservedAttempts()
	for _, attempt := range attemptList {
		fmt.Printf("%+v \n", attempt)
	}
}

func (callback *Callback) Fail(result *tencentcloud_cls_sdk_go.Result) {
	fmt.Println(result.IsSuccessful())
	fmt.Println(result.GetErrorCode())
	fmt.Println(result.GetErrorMessage())
	fmt.Println(result.GetReservedAttempts())
	fmt.Println(result.GetRequestId())
	fmt.Println(result.GetTimeStampMs())
}


```