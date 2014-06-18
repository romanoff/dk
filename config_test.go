package main

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("testconfig")
	if err != nil {
		t.Errorf("Expected to not get error while reading config, but got : %v", err)
	}
	if len(config.Sources) != 2 {
		t.Errorf("Expected to get 2 sources from config, but got %v", len(config.Sources))
	}
	if len(config.Remotes) != 1 {
		t.Errorf("Expected to get 1 remote from config, but got %v", len(config.Sources))
	}
}
