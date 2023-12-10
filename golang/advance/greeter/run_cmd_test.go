package main

import (
    "bytes"
    "errors"
    "strings"
    "testing"
)

func TestRunCmd(t *testing.T) {
    tests := []struct{
        c config
        input string
        output string
        err error
    }{
        {
            c: config{printUsage: true},
            output: usageString,
        },
        {
            c: config{numTimes: 5},
            input: "",
            output: strings.Repeat("Your name please?Press the Enter key when done.\n", 1),
            err: errors.New("You didn't enter your name"),
        },
        {
            c: config{numTimes: 5},
            input: "wangyao",
            output: "Your name please?Press the Enter key when done.\n" + strings.Repeat("Nice to meet you wangyao\n", 5),
        },
    }

    byteBuf := new(bytes.Buffer)
    for _, tc := range tests {

        rd := strings.NewReader(tc.input)
        err := runCmd(rd, byteBuf, tc.c)
        if err != nil && tc.err == nil {
            t.Fatalf("Excepted nil error, got: %v\n", err)
        }
        if tc.err != nil && err.Error() != tc.err.Error() {
            t.Fatalf("Excepted error: %v, got: %v", tc.err.Error(), err.Error())
        }
        goMsg := byteBuf.String()
        if goMsg != tc.output {
            t.Errorf("Excepted stdout message to be: %v,Got %v\n", tc.output, goMsg)
        }
        byteBuf.Reset()

    }
}
