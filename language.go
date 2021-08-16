package carbon

import (
	"strconv"
	"strings"

	"github.com/golang-module/carbon/lang"
)

var (
	defaultLocale = "en"
)

// Language define a Language struct.
// 定义 Language 结构体
type Language struct {
	dir       string
	locale    string
	resources map[string]string
}

// NewLanguage return a new Language instance.
// 初始化 Language 结构体
func NewLanguage() *Language {
	return &Language{
		locale:    defaultLocale,
		resources: make(map[string]string),
	}
}

// SetLocale sets language locale.
// 设置区域
func (l *Language) SetLocale(locale string) error {
	if len(l.resources) != 0 {
		return nil
	}
	l.locale = locale
	resources, err := lang.LoadLocale(locale)
	if err != nil {
		return err
	}
	l.resources = resources
	return nil
}

// SetResources sets language resources.
// 设置资源
func (l *Language) SetResources(resources map[string]string) {
	if len(l.resources) == 0 {
		l.resources = resources
		return
	}
	for k, v := range resources {
		if _, ok := l.resources[k]; ok {
			l.resources[k] = v
		}
	}
}

// translate translates by unit string.
// 翻译转换
func (l *Language) translate(unit string, number int64) string {
	if len(l.resources) == 0 {
		l.SetLocale(defaultLocale)
	}
	slice := strings.Split(l.resources[unit], "|")
	if len(slice) == 1 {
		return strings.Replace(slice[0], "%d", strconv.FormatInt(number, 10), 1)
	}
	if number > 1 {
		return strings.Replace(slice[1], "%d", strconv.FormatInt(number, 10), 1)
	}
	return slice[0]
}
