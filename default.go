package carbon

var (
	// default layout
	// 默认布局模板
	defaultLayout = DateTimeLayout

	// default timezone
	// 默认时区
	defaultTimezone = Local

	// default locale
	// 默认区域
	defaultLocale = "en"
)

// Default defines a Default struct.
// 定义 Default 结构体
type Default struct {
	Layout   string
	Timezone string
	Locale   string
}

// SetDefault sets default.
// 设置全局默认值
func SetDefault(d Default) {
	if d.Layout != "" {
		defaultLayout = d.Layout
	}
	if d.Timezone != "" {
		defaultTimezone = d.Timezone
	}
	if d.Locale != "" {
		defaultLocale = d.Locale
	}
}
