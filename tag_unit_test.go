package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarbon_Tag(t *testing.T) {
	now := Now().SetTag("layout")
	assert.Equal(t, "layout", now.Tag())
}

func TestCarbon_SetTag(t *testing.T) {
	now := Now().SetTag("layout")
	assert.Equal(t, "layout", now.tag)
}

func TestError_SetTag(t *testing.T) {
	now := Now("xxx").SetTag("layout")
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
	assert.Equal(t, err2, invalidTagError())
}
