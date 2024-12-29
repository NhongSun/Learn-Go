package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func cronLib() {
	// cronJob()

	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("HI, Every 5 seconds")
	})

	// start goroutine
	c.Start()

	// ถ้าไม่มีการใช้งาน goroutine จะหยุดทำงานทันที
	// แต่ถ้ามีการใช้งาน goroutine จะต้องใช้ select
	// เพื่อรอให้ goroutine ทำงานเสร็จสิ้น
	select {}
}
