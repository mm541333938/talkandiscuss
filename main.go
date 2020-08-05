package main

import (
	"api/models"
	"api/pkg/logging"
	"api/pkg/setting"
	"api/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Fatal("Hello, api 正在启动中")
	setting.SetUp() //初始化配置文件
	logging.SetUp() //设置日志文件
	models.SetUp()  //设置数据库

	log.Println(setting.ServerSetting.HttpPort)

	router := routers.InitRouter() // 初始化路由

	//router.GET("/test", func(context *gin.Context) {
	//	context.JSON(e.SUCCESS, gin.H{
	//		"Code": e.SUCCESS,
	//		"Msg":  e.GetMsg(e.SUCCESS),
	//		"Data": "返回数据成功",
	//	})
	//})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	/**
	使用 Http.server -shutdown() 关闭服务
	*/

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server ShutDown:", err)
	}
	log.Println("程序服务关闭退出")
}
