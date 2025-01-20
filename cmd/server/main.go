package main

import (
	"log"
	"pve-control-panel-backend/internal/app"
)

func main() {
	// 初始化应用
	application := app.NewApp()
	// 启动Gin服务器
	if err := application.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
