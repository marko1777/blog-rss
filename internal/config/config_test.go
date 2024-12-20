package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	t.Run("", func(t *testing.T) {

		want := &Config{DB_URL: "postgres://example"}
		if got := Read(); !assert.Equal(t, got.DB_URL, want.DB_URL) {
			t.Errorf("Read() = %v, want %v", got, want)
		}
	})
}

func TestConfig_SetUser(t *testing.T) {
	prev := Read()
	t.Run("", func(t *testing.T) {
		this := &Config{
			DB_URL:          "postgres://example",
			CurrentUserName: "mark",
		}
		if err := this.SetUser("mark"); (err != nil) != false {
			t.Errorf("Config.SetUser() error = %v, wantErr %v", err, false)
		}
		curr := Read()
		if !assert.Equal(t, curr, this) {
			t.Errorf("Config not same; expected: %v, actual: %v", this, curr)
		}
	})
	prev.write()
}
