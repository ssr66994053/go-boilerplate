package graceful

import (
	"context"
	"log"
	"os"
	"os/signal"
)

// Watch 监听信号
func Watch(shutFunc func()) {
	ctx, cancel := context.WithCancel(context.Background())

	// 监听信号
	go func() {
		defer cancel()
		// 这里同时监听2个信号
		stopSign := make(chan os.Signal, 1)
		signal.Notify(stopSign, os.Interrupt)
		// 阻塞等待信号
		select {
		case <-stopSign:
			log.Println("get system signal, start shutdown...")
			shutFunc()
			log.Println("shutdown success")
			break
		}
	}()

	// 等待
	select {
	case <-ctx.Done():
		log.Println("get shutdown")
		break
	}

	log.Println("graceful shutdown...")
}
