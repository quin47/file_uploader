package uploader

import (
	"file_uploader/notification"
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"log"
	"os"
	"strconv"
)

var minioClient  *minio.Client

var endpoint = os.Getenv("endpoint")

var bucket = os.Getenv("bucket")

func init()  {
	// Use a secure connection.

	if bucket == ""{
		bucket="upload"

	}

	accessKey := os.Getenv("accessKey")
	secrectKey := os.Getenv("secretKey")
	sslS := os.Getenv("ssl")
	ssl, e := strconv.ParseBool(sslS)
	if e != nil {
		ssl=false
	}

	// Initialize minio client object.
	var err error
	minioClient, err = minio.New(endpoint, accessKey, secrectKey, ssl)
	if err != nil {
		log.Fatal("init mc client failed !",err)
	}
	b, e := minioClient.BucketExists("upload")
	if e != nil {
		log.Fatal("error to access bucket",e)
	}
	if !b {
		minioClient.MakeBucket("upload","")

		notification.SimpleNotify(fmt.Sprintf("new minio bucket %v created", bucket)," 请前往 minio web 设置权限")
	}


	if err != nil {
		log.Fatalf("cannot find policy of %v,err: %v",bucket,err)
	}

}

func Upload(fname string,reader io.Reader, size int64,contentType string) (httpPath string) {
	n, err := minioClient.PutObject(bucket, fname, reader, size, minio.PutObjectOptions{ContentType:contentType})
	if err != nil {
		log.Println("upload failed",err)
		return ""
	}




	log.Println("upload finished" ,n)
	return "http://"+endpoint+"/"+bucket+"/"+fname;

}
