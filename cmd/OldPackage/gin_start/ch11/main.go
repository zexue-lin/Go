package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// 优雅的退出 , 当我们关闭程序的时候应该做的后续处理
	// 微服务 一启动之前 或 启动之后 :将当前服务的ip地址和端口号注册到注册中心
	// 我们当前的服务停止了以后没有告知注册中心
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	srv := &http.Server{
		Addr:    ":8083",
		Handler: router.Handler(),
	}

	// 使用协程
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 如果我想要接受信号   Linux中 kill -9  强制关闭进程,不会发送信号
	quit := make(chan os.Signal, 1)                      // chan关键字 os.Signal类型的通道(操作系统信号) 缓冲区大小为1
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 中断和请求进程终止的信号
	<-quit                                               // 从quit通道中接受一个信号，是个阻塞操作，一旦接收到信号，才会执行后面的代码

	// 处理后续的逻辑
	log.Println("关闭server中...")

	// 创建一个带有 5 秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 捕获关闭服务器时的错误
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务关闭出错:", err)
	}

	// 捕获超时
	select {
	case <-ctx.Done():
		log.Println("已超时5秒...")
	}

	log.Println("服务退出...")
}
