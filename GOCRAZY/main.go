package main

import (
    "flag"
    "github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
    Name string
    Host string `json:",default=0.0.0.0"`
    Port int
}

var f = flag.String("f", "config.yaml", "config file")

func main() {
    flag.Parse()
    var c Config
    conf.MustLoad(*f, &c)
    println(c.Name)
    println(c.Host)
    println(c.Port)
}