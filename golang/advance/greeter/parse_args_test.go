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
        {
            args: []string{"-h"},
            err: nil,
            config: config{printUsage: true, numTimes: 0},
        },
        {
            args: []string{"10"},
            err: nil,
            config: config{printUsage: false, numTimes: 10},
        },
        {
            args: []string{"abc"},
            err: errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
            config: config{printUsage: false, numTimes: 0},
        },
        {
            args: []string{"l", "foo"},
            err: errors.New("Invalid number of arguments"),
            config: config{printUsage: false, numTimes: 0},
        },
    }

    for _, tc := range tests {
        c, err := parseArgs(tc.args)
        if tc.err != nil && err.Error() != tc.err.Error() {
            t.Fatalf("Excepted error to be: %v,got: %v\n", tc.err, err)
        }
        if tc.err == nil && err != nil {
            t.Errorf("Excepted nil error,got: %v\n", err)
        }
        if c.printUsage != tc.printUsage {
            t.Errorf("Excepted printUsage to be: %v, got: %v\n", tc.printUsage, c.printUsage)
        }
        if c.numTimes != tc.numTimes {
            t.Errorf("Excepted numTimes to be: %v, got: %v\n", tc.numTimes, c.numTimes)
        }

    }
}


