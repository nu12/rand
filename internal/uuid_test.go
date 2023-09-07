package rand

import (
	"regexp"
	"testing"
)

var validUUID = regexp.MustCompile(`\S{8}-\S{4}-\S{4}-\S{4}-\S{12}`)

func TestUUID(t *testing.T) {

	uuid := GenerateUUID()
	match := validUUID.Match([]byte(uuid))
	if !match {
		t.Errorf("Error creating UUID. Got %s", uuid)
	}

}
