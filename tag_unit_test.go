package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTag_SetLayout(t *testing.T) {
	tag := NewTag()
	tag.SetLayout(DateTimeLayout)

	now := Now().SetTag(tag)
	assert.Nil(t, now.Error)
	assert.Equal(t, "layout:"+DateTimeLayout, now.tag.carbon)
}

func TestTag_SetFormat(t *testing.T) {
	tag := NewTag()
	tag.SetFormat(DateTimeFormat)

	now := Now().SetTag(tag)
	assert.Nil(t, now.Error)
	assert.Equal(t, "format:"+DateTimeFormat, now.tag.carbon)
}

func TestTag_SetType(t *testing.T) {
	tag := NewTag()
	tag.SetType("dateTime")

	now := Now().SetTag(tag)
	assert.Nil(t, now.Error)
	assert.Equal(t, "type:dateTime", now.tag.carbon)
}

func TestTag_SetTimezone(t *testing.T) {
	tag := NewTag()
	tag.SetTimezone(PRC)

	now := Now().SetTag(tag)
	assert.Nil(t, now.Error)
	assert.Equal(t, PRC, now.tag.tz)
}

func TestCarbon_parseTag(t *testing.T) {
	now := Now().SetTag(&Tag{
		carbon: "",
	})
	key, value, tz := now.parseTag()
	assert.Equal(t, "", key)
	assert.Equal(t, "", value)
	assert.Equal(t, Local, tz)
}

func TestTag_SetTag(t *testing.T) {
	tag := NewTag()
	tag.SetTimezone(PRC)

	newTag := SetTag(tag)
	assert.Nil(t, newTag.Error)
	assert.Equal(t, PRC, newTag.tag.tz)
}

func TestError_SetTag(t *testing.T) {
	now := Now("xxx").SetTag(&Tag{
		carbon: "datetime",
		tz:     Local,
	})
	assert.NotNil(t, now.Error)
}

func TestError_LoadTag(t *testing.T) {
	type Student struct {
		Birthday Carbon `json:"birthday" carbon:"xxx"`
	}
	student := Student{
		Birthday: Now(),
	}
	err1 := LoadTag(student)
	assert.Equal(t, err1, invalidPtrError())

	err2 := LoadTag(&student)
	assert.Equal(t, err2, invalidTagError("Birthday"))
}
