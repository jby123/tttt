package main

import (
	"github.com/robfig/cron"
	"log"
	"testing"
)

func TestCron(t *testing.T) {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running :", i)
	})
	c.Start()
	//这里的select是相当于死循环的意义，如果不写，程序会立刻退出，写for死循环也是可以的
	select {}
}
