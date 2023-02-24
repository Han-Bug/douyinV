package minio

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"strings"
)

const (
	Endpoint = "http://36.41.174.18:9000"

	AccessKey  = "SIhak41IDQUAfRG6"
	SecretKey  = "PhF5FONp75UsGM7McRmDvNZcwOXsmSUQ"
	BucketName = "douyin"
	UseSSL     = false
)

var (
	minioClient *minio.Client
)

func Init() error {
	client, err := minio.New(Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKey, SecretKey, ""),
		Secure: UseSSL,
	})
	if err != nil {
		fmt.Println("minio初始化失败")
		return err
	}
	minioClient = client
	return nil
}

// 实际应用中使用已提前创建好的桶

func CreateBucket(ctx context.Context, bucketName string) (bool, error) {
	// 校验桶名称
	if len(bucketName) == 0 {
		return false, errors.New("桶名称为空")
	}
	// 检查是否已存在同名桶
	isExist, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return false, err
	}
	if isExist {
		fmt.Printf("已存在名为 %s 的桶\n", bucketName)
		return false, nil
	}

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		fmt.Printf("创建桶 %s 失败\n", bucketName)
		return false, err
	}
	return true, nil
}

func UploadObj(ctx context.Context, bucketName string, objName string, reader io.Reader, objSize int64) error {

	info, err := minioClient.PutObject(ctx, bucketName, objName, reader, objSize, minio.PutObjectOptions{})
	if err != nil {
		fmt.Printf("文件 %s 上传失败\n", objName)
		return err
	} else {
		fmt.Println(info)
		return nil
	}
}

// GetObjUrl 直接访问对应url（桶权限需为readonly）
func GetObjUrl(bucketName string, objName string) string {
	builder := strings.Builder{}
	builder.WriteString(Endpoint)
	builder.WriteString("/")
	builder.WriteString(bucketName)
	builder.WriteString("/")
	builder.WriteString(objName)
	return builder.String()
}
