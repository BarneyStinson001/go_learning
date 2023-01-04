package cmd

import (
	"github.com/BarneyStinson001/go_project_tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	//Long: "支持多种单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command,args []string){
		var content string
		switch mode {
		case MODE_UPPER:
			content=word.ToUpper(str)
		case MODE_LOWER:
			content=word.ToLower(str)
		case MODE_UNDERLINE_TO_UPPER_CAMELCASE:
			content=word.UnderLineToUpperCamelCase(str)
		case MODE_UNDERLINE_TO_LOWER_CAMELCASE:
			content=word.UnderLineToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERLINE:
			content=word.CamelCaseToUnderLine(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word查看帮助文档")
		}
		log.Printf("输出结果：%s",content)

	},

}

var str string
var mode int8

func init(){
	wordCmd.Flags().StringVarP(&str,"str","s","","请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",0,"请输入单词转换的模式")
}

const (
	MODE_UPPER  = iota + 1
	MODE_LOWER
	MODE_UNDERLINE_TO_UPPER_CAMELCASE
	MODE_UNDERLINE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERLINE
)

var desc = strings.Join([]string{
		"该子命令支持各种单词转换，模式如下：",
	"1:全部单词转大写",
	"2:全部单词转小写",
	"3:下划线转大写驼峰",
	"4:下划线转小写驼峰",
	"5:驼峰转下划线",
}	,"\n")

