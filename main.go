package main

import (
	"cos_upload/cosupload"
	"cos_upload/restore"
	"cos_upload/upset"
	"flag"
	"fmt"
	"os"
)

// receive is usege scanln user input info
var receive, receive2 string

// Helpdoc if input -h or --h  then print helpdoc
var Helpdoc *bool

// Helpdoc2 if input -help or --help then print helpdoc
var Helpdoc2 *bool

// BucketURL use receive bucket_url
var BucketURL *string

// SecretID use receive secretid
var SecretID *string

// SecretKEY use receive secretkey
var SecretKEY *string

// SSecretID use encrypt secretid
var SSecretID *string

// SSecretKEY use encrypt secrekey
var SSecretKEY *string

// ENCryptstring use encrypt string
var ENCryptstring *string

// BFILEName the file name in bucket
var BFILEName *string

// SYSFILEDir the file name and dir in system
var SYSFILEDir *string

//init fast start
func init() {
	BucketURL = flag.String("url", receive, "bucket_url")
	Helpdoc = flag.Bool("h", false, "-h for helpdoc")
	Helpdoc2 = flag.Bool("help", false, "--help for helpdoc")
	SecretID = flag.String("sd", receive, "secretid")
	SecretKEY = flag.String("sk", receive, "secretkey")
	SSecretID = flag.String("sds", receive, "encrypt secretid")
	SSecretKEY = flag.String("sks", receive, "encrypt secretkey")
	BFILEName = flag.String("fnm", receive, "the file name in bucketthe file name in bucket")
	SYSFILEDir = flag.String("fdir", receive, "the file name and dir in system")
	ENCryptstring = flag.String("enc", receive, "encrypt string")

}

// helpinfo cosupload-v1 helpinfo
func helpinfo(s, s2 *bool) {
	var receive2 = flag.Arg(0)
	if *s || *s2 {
		fmt.Println("usage: cosupload -url $BucketURL -sd $SecretID -sk $SecretKEY -fnm $FileNAME -fdir $FileDIR")
		fmt.Println("       cosupload -url $BucketURL -sds $SecretID -sks $SecretKEY -fnm $FileNAME -fdir $FileDIR")
		fmt.Printf("%5s  %s\n", "-url ", "bucket_url")
		fmt.Printf("%5s  %s\n", "-sd  ", "secretid")
		fmt.Printf("%5s  %s\n", "-sk  ", "secretkey")
		fmt.Printf("%5s  %s\n", "-sds ", "receive encrypt secretid")
		fmt.Printf("%5s  %s\n", "-sks ", "receive encrypt secretkey")
		fmt.Printf("%5s  %s\n", "-fnm ", "file dir and name will be put in bucket")
		fmt.Printf("%5s  %s\n", "-enc ", "encrypt the string will use as encrypt secretid or secretkey ")
		fmt.Printf("%5s  %s\n%s\n", "-fdir", "the upload file dir and file name in the system", "eg 1:")
		fmt.Println("       cosupload -url https://********* -sd ****** -sk ****** -fnm /APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz")
		fmt.Printf("%5s\n", "eg 2:")
		fmt.Println("       ssecretid=`cosupload-v1 -enc secretid`; ssecretkey=`cosupload-v1 -enc secretkey` ")
		fmt.Println("       cosupload -url https://********* -sds $ssecretid -sks $ssecretkey -fnm /APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz")
		os.Exit(0)
	}
	if receive2 != "" {
		fmt.Println("parameter error,please check `cosupload-v1 -h`")
		os.Exit(0)
	}
}

// CHEckouT *string to string
func CHEckouT(s *string) string {
	return *s
}

func main() {
	flag.Parse()
	helpinfo(Helpdoc, Helpdoc)
	ENCryptstringS := CHEckouT(ENCryptstring)
	if ENCryptstringS != "" {
		fmt.Println("++++++++++++++encrypt string+++++++++:", upset.StringUPSET(ENCryptstringS))
		os.Exit(0)
	}
	BucketURLS := CHEckouT(BucketURL)
	SecretIDS := CHEckouT(SecretID)
	SSecretIDS := CHEckouT(SSecretID)
	SecretKEYS := CHEckouT(SecretKEY)
	SSecretKEYS := CHEckouT(SSecretKEY)
	BFILENameS := CHEckouT(BFILEName)
	SYSFILEDirS := CHEckouT(SYSFILEDir)
	if BucketURLS == "" {
		fmt.Println("BucketURL is not find,please check `cosupload -h`")
		os.Exit(0)
	}
	if SecretIDS == "" && SSecretIDS == "" {
		fmt.Println("SecretID is not find,please check `cosupload -h`")
		os.Exit(0)
	}
	if SecretKEYS == "" && SSecretKEYS == "" {
		fmt.Println("SecretKEY is not find,please check `cosupload -h`")
		os.Exit(0)
	}
	//如果SYSFILEDirS目录不存在则退出
	if _, err := os.Stat(SYSFILEDirS); os.IsNotExist(err) {
		fmt.Println("the file name and dir in system does not exist")
		os.Exit(0)
	}

	if SSecretIDS == "" && SSecretKEYS == "" && SecretIDS != "" && SecretKEYS != "" {
		// fmt.Println("yuanlaijiami:", SecretIDS, SecretKEYS)
		cosupload.COSUpload(BucketURLS, SecretIDS, SecretKEYS, BFILENameS, SYSFILEDirS)
		os.Exit(0)
	} else if SecretIDS == "" && SecretKEYS == "" && SSecretIDS != "" && SSecretKEYS != "" {
		// fmt.Println("jiami:", SSecretIDS, SSecretKEYS)
		SSecretIDS := restore.StringRESTORE(SSecretIDS)
		SSecretKEYS := restore.StringRESTORE(SSecretKEYS)
		cosupload.COSUpload(BucketURLS, SSecretIDS, SSecretKEYS, BFILENameS, SYSFILEDirS)
		os.Exit(0)
	} else {
		fmt.Println("cosupload-v1 can't use secretid and encrypt secretid as the same time!")
		fmt.Println("cosupload-v1 can't use secretkey and encrypt secretkey as the same time!")
		os.Exit(0)
	}

	// ENCryptstringSJIE := restore.StringRESTORE(upset.StringUPSET(ENCryptstringS))

}
