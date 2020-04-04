package receive

import (
	"testing"
	h "gorabbit/helper"
)

func TestGetRabbitMQURL(t *testing.T) {
	// test for empty config values
	emptyConfigResult := h.GetRabbitMQURL()

	if emptyConfigResult == "" {
		t.Errorf("got empty values from the config")
	}
}
