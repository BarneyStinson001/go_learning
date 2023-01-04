package cmd

import (
	"github.com/BarneyStinson001/go_project_tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use: "time",
	Short: "时间格式转换",
	Long: "时间格式转换",
	Run: func(cmd *cobra.Command, args []string) {

	},

}
var nowTimeCmd =&cobra.Command{
	Use: "now",
	Short: "获取当前时间",
	Long: "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime :=timer.GetNowTime()
		log.Printf("输出结果：%s,%d",nowTime.Format("2006-01-02 15:04:05"),nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use: "calc",
	Short: "计算所需时间",
	Long: "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout ="2006-01-02 15:04:05"
		if calculateTime==""{
			currentTimer=timer.GetNowTime()
		}else{
			var err error
			if !strings.Contains(calculateTime," "){
				layout="2016-01-02"
			}
			currentTimer,err=time.Parse(layout,calculateTime)
			if err!=nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		calculateTime,err:=timer.GetCalculateTime(currentTimer,duration)
		if err!=nil{
			log.Fatalf("time.GetCalculateTime err:%v",err)
		}
		log.Printf("输出结果：%s,%d",calculateTime.Format(layout),calculateTime.Unix())
	},
}

func init()  {
	calculateTimeCmd.Flags().StringVarP(&calculateTime,"calculate","c","",`需要计算的时间，有效单位为时间戳或已经格式化的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration,"duration","d","",`持续时间，有效时间为为"ns","us","ms","s","m","h"`)
}