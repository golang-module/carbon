package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarbon_SetTag(t *testing.T) {
	now := Now().SetTag(tag{
		carbon: "datetime",
		tz:     Local,
	})
	assert.Nil(t, now.Error)
	assert.Equal(t, "datetime", now.tag.carbon)
	assert.Equal(t, Local, now.tag.tz)
}

func TestError_SetTag(t *testing.T) {
	now := Now("xxx").SetTag(tag{
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
