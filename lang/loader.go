// As a part of the package carbon
package lang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// 语言目录
	langDir = "./lang/"

	// 无效的区域错误
	invalidLocaleError = func(locale string) error {
		return fmt.Errorf("invalid locale %q, please see the directory %q for all valid locales", locale, langDir)
	}

	// 无效的json文件错误
	invalidJsonFileError = func(file string) error {
		return fmt.Errorf("invalid json file %q, please make sure the json file is valid", langDir+file)
	}
)

// LoadLocale load a given locale file
// 加载本地化文件
func LoadLocale(locale string) (map[string]string, error) {
	resources := make(map[string]string)
	fileName := langDir + locale + ".json"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return resources, invalidLocaleError(locale)
	}
	if err := json.Unmarshal(bytes, &resources); err != nil {
		return resources, invalidJsonFileError(fileName)
	}
	return resources, nil
}
