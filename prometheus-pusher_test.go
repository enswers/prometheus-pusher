package main

import (
	"github.com/mgutz/logxi/v1"
	"testing"
)

func Test_getConfigFiles(t *testing.T) {
	configPath := "./tests"
	files := getConfigFiles(configPath)

	if files == nil {
		t.Error("should raise exception with invalid path.")
	}
}

func Test_getConfigFilesWithInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	logger = log.NullLog

	configPath := "/path/to/invalid"
	getConfigFiles(configPath)

	t.Error("should raise exception with invalid path.")
}

func Test_parseConfig(t *testing.T) {
	config, err := parseConfig("./tests")
	if err != nil {
		t.Error("should be success")
	}

	if config.PushGatewayURL != "http://localhost:9091" {
		t.Error("pushgateway_url is not matched.", config.PushGatewayURL)
	}

	if config.PushGatewayUsername != "guri" {
		t.Error("pushgateway_username is not matched.", config.PushGatewayUsername)
	}

	if config.PushGatewayPassword != "1234" {
		t.Error("pushgateway_password is not matched.", config.PushGatewayPassword)
	}

	if config.Metrics[0].Name == "metric_name" {
		t.Error("metric_name is not matched.", config.Metrics[0].Name)
	}
}
