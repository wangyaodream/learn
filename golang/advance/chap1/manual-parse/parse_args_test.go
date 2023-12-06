package main

import (
    "errors"
    "testing"
)

type testConfig struct {
    args []string
    err error
    config
}

func TestParseArgs(t *testing.T) {
    tests := []testConfig{
        args: []string{"-h"},
        err: nil,
        config: config{printUsage: true, numTimes: 0}
    }
    
}
