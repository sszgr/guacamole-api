package guacamole

import (
	"os"
	"testing"
)

func getGuacamole() *Guacamole {
	return &Guacamole{
		BaseUrl:    os.Getenv("GUACAMOLE_URI"),
		Username:   os.Getenv("GUACAMOLE_USERNAME"),
		Password:   os.Getenv("GUACAMOLE_PASSWORD"),
		SkipVerify: true,
	}
}

func TestConnect(t *testing.T) {
	g := getGuacamole()
	err := g.Connect()
	if err != nil {
		t.Errorf("TestConnect error, err: %s", err)
	}
}
