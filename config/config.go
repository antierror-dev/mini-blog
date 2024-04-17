package config

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

var (
	Appconfig fiber.Config
)

func init() {
	Appconfig = fiber.Config{
		AppName:         "goapp mmdali",
		ServerHeader:    "go app v1",
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		BodyLimit:       4 * 1024 * 1024,
		CaseSensitive:   false,
		Concurrency:     256 * 1024,
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
	}
}
