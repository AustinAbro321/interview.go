package phoneNumber

import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	var ph = PhoneNumber{"+1(858)775-2868"}
	
	assert.Equal(t, "+1(858)775-2868", ph.getOriginalText())
	assert.Equal(t, "+18587752868", ph.getStrippedNumber())
	assert.Equal(t,"(858)775-2868", ph.getValueAsNorthAmerican());
	assert.Equal(t,"+1.858.775.2868", ph.getValueAsInternational());
}

func Test2(t *testing.T) {
    var ph = PhoneNumber{"+1(858)775-2868x123"}
	assert.Equal(t,"+1(858)775-2868x123", ph.getOriginalText())
	assert.Equal(t,"+18587752868x123", ph.getStrippedNumber())
	assert.Equal(t,"(858)775-2868x123", ph.getValueAsNorthAmerican())
	assert.Equal(t,"+1.858.775.2868x123", ph.getValueAsInternational())
}

func Test3(t *testing.T) {
    var ph = PhoneNumber{"+598.123.4567x858"}
	assert.Equal(t,"+598.123.4567x858",ph.getOriginalText())
	assert.Equal(t,"+5981234567x858",ph.getStrippedNumber())
	assert.Equal(t,"",ph.getValueAsNorthAmerican())
	assert.Equal(t,"+598.123.456.7x858",ph.getValueAsInternational())
}

func Test4(t *testing.T) {
    var ph = PhoneNumber{"+27 1234 5678 ext 4"}
    assert.Equal(t,"+27 1234 5678 ext 4", ph.getOriginalText());
    assert.Equal(t,"+2712345678x4", ph.getStrippedNumber());
    assert.Equal(t,"", ph.getValueAsNorthAmerican());
    assert.Equal(t,"+27 1234 5678 ext 4", ph.getValueAsInternational());
}

func Test5(t *testing.T) {
    var ph = PhoneNumber{"858-775-2868"}
    assert.Equal(t,"858-775-2868", ph.getOriginalText());
    assert.Equal(t,"+18587752868", ph.getStrippedNumber());
    assert.Equal(t,"(858)775-2868", ph.getValueAsNorthAmerican());
    assert.Equal(t,"+1.858.775.2868", ph.getValueAsInternational());
}
