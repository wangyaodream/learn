package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)


var totalDuration time.Duration = 5


func getName(r io.Reader, w io.Writer) (string, error) {
    scanner := bufio.NewScanner(r)
    msg := "Your name please? Press the Enter key when done"
    fmt.Fprintln(w, msg)

    scanner.Scan()
    if err := scanner.Err(); err != nil {
        return "", err
    }

    name := scanner.Text()
    if len(name) == 0 {
        return "", errors.New("you entered an empty name")
    }
    return name, nil
}

func getNameContext(ctx context.Context) (string, error) {
    var err error
    name := "Default name"
    c := make(chan error, 1)

    go func() {
        name, err = getName(os.Stdout, os.Stdin)
        c <- err
    }()

    select {
    case <-ctx.Done():
        return name, ctx.Err()
    case err := <-c:
        return name, err
    }

}

func main() {
    allowedDuration := totalDuration * time.Second

    ctx, cancle := context.WithTimeout(context.Background(), allowedDuration)
    defer cancle()

    name, err := getNameContext(ctx)

    if err != nil && !errors.Is(err, context.DeadlineExceeded) {
        fmt.Fprintf(os.Stdout, "$v\n", err)
        os.Exit(1)
    }

    fmt.Fprintln(os.Stdout, name)
}
