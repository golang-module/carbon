package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarbon_parseTag(t *testing.T) {
	now := Now().SetTag(&tag{
		carbon: "",
	})
	key, value, tz := now.parseTag()
	assert.Equal(t, "layout", key)
	assert.Equal(t, DateTimeLayout, value)
	assert.Equal(t, Local, tz)
}

func TestCarbon_SetTag(t *testing.T) {
	c := NewCarbon().SetTag(&tag{
		tz: PRC,
	})
	assert.Nil(t, c.Error)
	assert.Equal(t, PRC, c.tag.tz)
}

func TestError_SetTag(t *testing.T) {
	now := Now("xxx").SetTag(&tag{
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
