package request

import (
	"fmt"
	"io"

	"http/utils"
)

type Config struct {
	maxRequestSize int
}

func NewRequest(conn io.Reader, config *Config) (*Request, error) {
	config = defaultConfig(config)

	buffer := make([]byte, config.maxRequestSize)

	_, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	request, err := requestParser(buffer)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func defaultConfig(config *Config) *Config {
	if config == nil {
		config = &Config{
			maxRequestSize: 1024,
		}
	}
	config.maxRequestSize = utils.IfEmptyInt(config.maxRequestSize, 1024)

	return config
}
