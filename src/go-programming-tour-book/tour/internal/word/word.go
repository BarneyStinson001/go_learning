package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderLineToUpperCamelCase(s string) string {
	s = strings.Replace(s,"_"," ",-1)
	s = strings.Title(s)
	return strings.Replace(s," ","",-1)
}

func UnderLineToLowerCamelCase(s string) string	{
	s = UnderLineToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0])))+s[1:]
}

func CamelCaseToUnderLine(s string) string {
	var output []rune
	for i,r :=range s{
		if i==0{
			output=append(output,unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r){
			output=append(output,'_')
		}
		output=append(output,unicode.ToLower(r))

	}
	return string(output)
}