package main

import (
    "fmt"
)


type notifier interface {
    notify()
}

type user struct {
    name string
    email string
}

// 实现notifier接口
func (u *user) notify() {
    fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
    user
    level string
}

func (a *admin) notify() {
    fmt.Printf("Sending admin email to %s<%s>", a.name, a.email)
}

func main() {
    ad := admin{
        user{
            "wangyao",
            "wangyao@email.com",
        },
        "super",
    }
    sendNotification(&ad)

    ad.user.notify()

    ad.notify()
}

func sendNotification(n notifier) {
    n.notify()
}


