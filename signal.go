package utils

import (
    "os"
    "os/signal"
    "syscall"
    "log"
)

func Notify(callback func()) {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.Signal(0xa))
    go func() {
        for syscall.Signal(0xa) == <-sigs {
            log.Print("Recieved 0xa, reloading config")
            callback()
        }
    }()

}
