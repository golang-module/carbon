package carbon

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCarbon_Closest(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string
		input2   string
		input3   string
		expected string
	}{
		{"", "2023-03-28", "2023-04-16", "2023-03-28"},
		{"2023-04-01", "", "2023-04-16", "2023-04-16"},
		{"2023-04-01", "2023-03-28", "", "2023-03-28"},
		{"2023-04-01", "", "", ""},

		{"2023-04-01", "2023-03-28", "2023-03-28", "2023-03-28"},
		{"2023-04-01", "2023-03-28", "2023-04-16", "2023-03-28"},
	}

	for index, test := range tests {
		c := Parse(test.input1).Closest(Parse(test.input2), Parse(test.input3))
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Farthest(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string
		input2   string
		input3   string
		expected string
	}{
		{"", "2023-03-28", "2023-04-16", "2023-04-16"},
		{"2023-04-01", "", "2023-04-16", "2023-04-16"},
		{"2023-04-01", "2023-03-28", "", "2023-03-28"},
		{"2023-04-01", "", "", ""},

		{"2023-04-01", "2023-03-28", "2023-03-28", "2023-03-28"},
		{"2023-04-01", "2023-03-28", "2023-04-16", "2023-04-16"},
		{"2023-04-01", "2023-04-05", "2023-04-02", "2023-04-05"},
	}

	for index, test := range tests {
		c := Parse(test.input1).Farthest(Parse(test.input2), Parse(test.input3))
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Max(t *testing.T) {
	now := Now()
	max := Max(now.SubDay(), now, now.AddDay())
	assert.Equal(t, now.AddDay().Timestamp(), max.Timestamp())
}

func TestCarbon_Min(t *testing.T) {
	now := Now()
	min := Min(now, now.SubDay(), now.AddDay())
	assert.Equal(t, now.SubDay().Timestamp(), min.Timestamp())
}
