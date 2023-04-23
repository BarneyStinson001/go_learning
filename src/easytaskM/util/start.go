package util


import (

	//"github.com/uber-go/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	//"github.com/spf13/viper"

)

func Start() {

	r := InitRouter()

	go func() {
		if err := r.Run("127.0.0.1:10296"); err != nil {
			//zap.S().Panic("启动失败", err.Error())
			log.Fatal("启动失败")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//zap.S().Info("注销")
	log.Fatal("注销")


}
