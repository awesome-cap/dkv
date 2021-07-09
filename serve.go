package main

import (
    "github.com/awesome-cap/dkv/exec"
    "github.com/awesome-cap/dkv/net"
    "github.com/awesome-cap/dkv/storage"
    "log"
)

func main() {
    e := exec.NewExec(storage.New())
    e.RegistryHandler("get", exec.GetHandler)
    e.RegistryHandler("set", exec.SetHandler)
    e.RegistryHandler("del", exec.DelHandler)

    tcpServer := net.NewTcp(":8888")
    err := tcpServer.Serve(func(args []string) ([]string, error) {
        return e.Exec(args)
    })
    if err != nil{
        log.Fatalln(err)
    }
}
