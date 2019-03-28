D:\go_workplace\src\cos_upload>go run main.go -h

usage: cosupload -url $BucketURL -sd $SecretID -sk $SecretKEY -fnm $FileNAME -fdir $FileDIR<br>
       cosupload -url $BucketURL -sds $SecretID -sks $SecretKEY -fnm $FileNAME -fdir $FileDIR<br>
-url   bucket_url<br>
-sd    secretid<br>
-sk    secretkey<br>
-sds   receive encrypt secretid<br>
-sks   receive encrypt secretkey<br>
-fnm   file dir and name will be put in bucket<br>
-enc   encrypt the string will use as encrypt secretid or secretkey<br>
-fdir  the upload file dir and file name in the system<br>

eg 1:<br>
       cosupload -url https://********* -sd ****** -sk ****** -fnm /APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz<br>
eg 2:
       ssecretid=`cosupload-v1 -enc secretid`; ssecretkey=`cosupload-v1 -enc secretkey`<br>
       cosupload -url https://********* -sds $ssecretid -sks $ssecretkey -fnm /APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz
