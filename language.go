package carbon

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Language struct {
	locale    string
	dir       string
	resources map[string]string
}

func NewLanguage() *Language {
	return &Language{
		locale:    "en",
		dir:       "./lang",
		resources: make(map[string]string),
	}
}

// SetLocale 设置语言区域
func (lang *Language) SetLocale(locale string) error {
	err := lang.loadResource(locale)
	if err == nil {
		lang.locale = locale
	}
	return err
}

// loadResource 加载资源
func (lang *Language) loadResource(locale string) error {
	fileName := lang.dir + string(os.PathSeparator) + locale + ".json"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return errors.New("not able to read the lang file:" + fileName)
	}
	if err := json.Unmarshal(bytes, &lang.resources); err != nil {
		return err
	}
	return nil
}

// translate 翻译
func (lang *Language) translate(unit string, count int64) string {
	s := strings.Split(lang.resources[unit], "|")
	if len(s) == 1 {
		return strings.Replace(s[0], "{count}", strconv.FormatInt(count, 10), 1)
	}

	if count > 1 {
		return strings.Replace(s[1], "{count}", strconv.FormatInt(count, 10), 1)
	}
	return s[0]
}
