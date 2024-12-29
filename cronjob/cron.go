package main

import (
	"fmt"
	"time"
)

func cronJob() {
	// สร้าง channel เพื่อกำหนดช่วงเวลาที่ต้องการเรียกใช้งานงาน
	ch := time.Tick(1 * time.Second)

	// สร้าง goroutine เพื่อเรียกใช้งานงานซ้ำ ๆ ในช่วงเวลาที่กำหนด
	go func() {
		for range ch {
			fmt.Println("Hello, world!")
		}
	}()

	// รอให้ goroutine ทำงานเสร็จสิ้น
	time.Sleep(5 * time.Second)
}
