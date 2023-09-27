package main


import "fmt"


type user struct {
    name    string
    email   string
    ext     int
    privileged  bool
}


var bill user


lisa := user{
    name: "lisa",
    email: "lisa@gmail.com",
    ext: 123,
    privileged: true,
}


