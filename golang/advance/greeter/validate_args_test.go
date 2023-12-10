package main

import (
    "errors"
    "testing"
)


func TestValidateArgs(t *testing.T) {
    // 匿名结构
    tests := []struct {
        c config
        err error
    }{
        {
            c: config{},
            err: errors.New("Must specify a number greeter than 0"),
        },
        {
            c: config{numTimes: -1},
            err: errors.New("Must specify a number greeter than 0"),
        },
        {
            c: config{numTimes: 10},
            err: nil,
        },

    }

    for _, tc := range tests {
        err := validateArgs(tc.c)
        if tc.err != nil && err.Error() != tc.err.Error() {
            t.Errorf("Excepted error to be: %v, got: %v", tc.err, err)
        }
        if tc.err == nil && err != nil {
            t.Errorf("Excepted nil error, got: %v", err)
        }
    }
    
}

