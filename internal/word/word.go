package word

import (
	"strings"
	"unicode"
)

// ToUpper: 单词全部转换成大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower: 单词全部转换成小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase: 下划线单词转成大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase: 下划线单词转成小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CameCaseToUnderscore: 驼峰单词转成下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
