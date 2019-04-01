<pre>
D:\go_workplace\src\cos_upload>go run main.go -h

usage: cosupload -url $BucketURL -sd $SecretID -sk $SecretKEY -fnm $FileNAME -fdir $FileDIR
       cosupload -url $BucketURL -sds $SecretID -sks $SecretKEY -fnm $FileNAME -fdir $FileDIR
-url   bucket_url
-sd    secretid
-sk    secretkey
-sds   receive encrypt secretid
-sks   receive encrypt secretkey
-fnm   file dir and name will be put in bucket
-enc   encrypt the string will use as encrypt secretid or secretkey
-fdir  the upload file dir and file name in the system

eg 1:<br>
       cosupload -url https://********* -sd ****** -sk ****** -fnm /APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz<br>
eg 2:
       ssecretid=`cosupload-v1 -enc secretid`; ssecretkey=`cosupload-v1 -enc secretkey`<br>
       cosupload -url https://********* -sds $ssecretid -sks $ssecretkey -fnm /APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz
</pre>
---
<br>

# **下载**<br>
[点击下载](https://github.com/qqzgqq/cos_upload_tool/releases/tag/V1.0)