package cmd

import (
    "errors"
)

var ErrNoServerSepcified = errors.New("you have to specify a remote server")
