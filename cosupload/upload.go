package cosupload

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// COSUpload put the backup file to cos bucket
func COSUpload(btURL, secID, secKEY, fileNAME, fileDIR string) {
	//将<bucketname>、<appid>和<region>修改为真实的信息
	//例如：http://test-1253846586.cos.ap-guangzhou.myqcloud.com
	u, _ := url.Parse(btURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  secID,
			SecretKey: secKEY,
		},
	})

	//对象键（Key）是对象在存储桶中的唯一标识。
	//例如，在对象的访问域名 ` bucket1-1250000000.cos.ap-guangzhou.myqcloud.com/test/objectPut.go ` 中，对象键为 test/objectPut.go
	var name string
	if m, _ := regexp.MatchString("^\\/", fileNAME); m {
		FishuzU := strings.SplitAfterN(fileNAME, "", 2)
		name = FishuzU[1]
	} else {
		name = fileNAME
	}

	//Local file
	// f := strings.NewReader(fileDIR)
	f, err := os.Open(fileDIR)
	if err != nil {
		panic(err)
	}
	// s, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(s.Size())
	_, err = c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileDIR, " UPLOAD SECUSS")
	//检测存储桶中所有对象别表
	opt := &cos.BucketGetOptions{
		Prefix: "",
	}
	v, _, err := c.Bucket.Get(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	//Test 用于下面判断
	var Test int
	//定义切片接收for循环中传递的c.key
	var FileStrIng []string
	//for循环将c.key的字符串apend至数组FileStrIng
	for _, c := range v.Contents {
		FileStrIng = append(FileStrIng, c.Key)
	}
	//判断刚上传的fileNAME是否在存在，存在Test为1，并退出循环
	for i := 0; i < len(FileStrIng); i++ {
		if FileStrIng[i] == name {
			Test = 1
			break
		}
	}
	if Test == 1 {
		fmt.Println(fileNAME, " is find on the bucket")
	} else {
		fmt.Println(fileNAME, " can't find on the bucket")
	}
}
