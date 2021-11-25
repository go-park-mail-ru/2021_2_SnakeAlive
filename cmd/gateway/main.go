package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/valyala/fasthttp"
	"snakealive/m/internal/gateway/config"
	"snakealive/m/internal/gateway/setup"
)

func main() {
	client := s3.New(
		session.Must(session.NewSession()),
		aws.NewConfig().WithEndpoint("http://hb.bizmrg.com"),
		aws.NewConfig().WithRegion("ru-msk"),
		aws.NewConfig().WithCredentials(
			credentials.NewStaticCredentials(
				"rcLyfa3DhATRNmeterCZLk",
				"7qdqBBT9xeKDV3iAdKco4jqEhmjpMqiyk3bLL9PcU2S7",
				"",
			),
		),
		//aws.NewConfig().WithDisableSSL(true),
		//aws.NewConfig().WithS3ForcePathStyle(true),
	)

	//fmt.Println(client.CreateBucket(&s3.CreateBucketInput{
	//	ACL:    aws.String("public-read-write"),
	//	Bucket: aws.String("tpvk"),
	//}))
	fmt.Println(client.PutObject(&s3.PutObjectInput{
		Body:                 bytes.NewReader([]byte("123123456789")),
		Key:                  aws.String("123"),
		Bucket:               aws.String("snakehastrip"),
		ACL:                  aws.String("public-read-write"),
		ServerSideEncryption: aws.String("AES256"),
	}))
	//fmt.Println(client.GetObject(&s3.GetObjectInput{
	//	Key:    aws.String("123"),
	//	Bucket: aws.String("tpvk"),
	//}))

	var cfg config.Config
	if err := cfg.Setup(); err != nil {
		log.Fatal("failed to setup cfg: ", err)
		return
	}

	logger := cfg.Logger.Sugar()
	r, stop, err := setup.Setup(cfg)
	if err != nil {
		logger.Fatal("msg", "failed to setup server", "error", err)
		return
	}

	go func() {
		if err := fasthttp.ListenAndServe(cfg.HTTPPort, corsMiddleware(r.Handler)); err != nil {
			logger.Fatal("failed to start server")
			return
		}
	}()

	logger.Info("gateway started ...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		logger.Info("msg", "received signal, exiting", "signal", sig)
		cfg.Cancel()
		stop()

		logger.Info("msg", " goodbye")
	}(<-c)
}

func corsMiddleware(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://194.58.104.204") // set domain
		ctx.Response.Header.Set("Content-Type", "application/json; charset=utf8")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		ctx.Response.Header.Set("Access-Control-Expose-Headers", "Authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Max-Age", "3600")

		if bytes.Equal(ctx.Method(), []byte(fasthttp.MethodOptions)) {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		handler(ctx)
	}
}
