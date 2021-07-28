package carbon

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	// default language dir
	// 默认目录
	defaultDir = "./lang"
	// default language locale
	// 默认区域
	defaultLocale = "en"
)

// Language 定义 Language 结构体
// define Language structure
type Language struct {
	dir       string            // 目录
	locale    string            // 区域
	resources map[string]string // 资源
}

// NewLanguage return a new Language instance
// 初始化 Language 结构体
func NewLanguage() *Language {
	return &Language{
		dir:       defaultDir,
		locale:    defaultLocale,
		resources: make(map[string]string),
	}
}

// SetLocale set language locale
// 设置区域
func (lang *Language) SetLocale(locale string) error {
	if len(lang.resources) != 0 {
		return nil
	}
	fileName := lang.dir + string(os.PathSeparator) + locale + ".json"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return invalidLocaleError(locale, lang.dir)
	}
	if err := json.Unmarshal(bytes, &lang.resources); err != nil {
		return invalidJsonFileError(fileName)
	}
	lang.locale = locale
	return nil
}

// SetDir set language directory
// 设置目录
func (lang *Language) SetDir(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		return invalidDirError(dir)
	}
	lang.dir = dir
	return nil
}

// SetResources set language resources
// 设置资源
func (lang *Language) SetResources(resources map[string]string) {
	if len(lang.resources) == 0 {
		lang.resources = resources
		return
	}
	for k, v := range resources {
		if _, ok := lang.resources[k]; ok {
			lang.resources[k] = v
		}
	}
}

// translate translate by unit string
// 翻译转换
func (lang *Language) translate(unit string, diff int64) string {
	if len(lang.resources) == 0 {
		lang.SetLocale(defaultLocale)
	}
	slice := strings.Split(lang.resources[unit], "|")
	if len(slice) == 1 {
		return strings.Replace(slice[0], "%d", strconv.FormatInt(diff, 10), 1)
	}
	if diff > 1 {
		return strings.Replace(slice[1], "%d", strconv.FormatInt(diff, 10), 1)
	}
	return slice[0]
}
