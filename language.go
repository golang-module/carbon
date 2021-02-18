package carbon

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	// 默认语言目录
	defaultDir = "lang"
	// 默认语言区域
	defaultLocale = "en"
)

type Language struct {
	locale    string
	dir       string
	resources map[string]string
}

// NewLanguage 初始化Language对象
func NewLanguage() *Language {
	return &Language{
		locale:    "",
		dir:       defaultDir,
		resources: make(map[string]string),
	}
}

// SetLocale 设置区域
func (lang *Language) SetLocale(locale string) error {
	if len(lang.resources) != 0 {
		return nil
	}
	fileName := lang.dir + string(os.PathSeparator) + locale + ".json"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return errors.New("invalid locale \"" + locale + "\", please see the " + lang.dir + " directory for all valid locale")
	}
	if err := json.Unmarshal(bytes, &lang.resources); err != nil {
		return err
	}
	lang.locale = locale
	return nil
}

// SetDir 设置目录
func (lang *Language) SetDir(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		return errors.New("invalid directory \"" + dir + "\"")
	}
	lang.dir = dir
	return nil
}

// SetResources 设置资源
func (lang *Language) SetResources(resources map[string]string) {
	if len(lang.resources) == 0 {
		lang.resources = resources
	} else {
		for k, v := range resources {
			lang.resources[k] = v
		}
	}
}

// translate 翻译转换
func (lang *Language) translate(unit string, diff int64) string {
	if len(lang.resources) == 0 {
		lang.SetLocale(defaultLocale)
	}
	array := strings.Split(lang.resources[unit], "|")
	if len(array) == 1 {
		return strings.Replace(array[0], "%d", strconv.FormatInt(diff, 10), 1)
	}
	if diff > 1 {
		return strings.Replace(array[1], "%d", strconv.FormatInt(diff, 10), 1)
	}
	return array[0]
}
