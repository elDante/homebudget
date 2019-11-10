package main

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	config := parseConfig("config.example.toml")

	if config.Database.Host != "127.0.0.1" {
		t.Error("For", "config.Database.Host", "expected", "127.0.0.1", "got", config.Database.Host)
	}

	if config.Database.Port != 3306 {
		t.Error("For", "config.Database.Port", "expected", "3306", "got", config.Database.Host)
	}

	if config.Database.Username != "homebudget" {
		t.Error("For", "config.Database.Username", "expected", "homebudget", "got", config.Database.Username)
	}

	if config.Database.Password != "homebudget" {
		t.Error("For", "config.Database.Password", "expected", "homebudget", "got", config.Database.Password)
	}

	if config.Database.Name != "homebudget" {
		t.Error("For", "config.Database.Name", "expected", "homebudget", "got", config.Database.Name)
	}
}
