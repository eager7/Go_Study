package crontab

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func Initialize() error {
	task := cron.New()
	//秒里有1时执行任务
	if err := task.AddFunc("1 * * * *", func() {
		//fmt.Println("task 1 for every minute:", time.Now().Local())
	}); err != nil {
		return err
	}
	//秒里有3或者15时执行任务
	if err := task.AddFunc("3,15 * * * *", func() {
		//fmt.Println("task 2 for 3,15 every minute:", time.Now().Local())
	}); err != nil {
		return err
	}
	//每秒执行一次
	if err := task.AddFunc("@every 1s", func() {
		//fmt.Println("task 3:", time.Now().Local())
	}); err != nil {
		return err
	}

	if err := task.AddFunc("0 */2 * * *", func() {
		//fmt.Println("task 4 for 2 minute:", time.Now().Local())
	}); err != nil {
		return err
	}

	if err := task.AddFunc("@every 1h", func() {
		fmt.Println("task 5 for 1 hour:", time.Now().Local())
	}); err != nil {
		return err
	}
	if err := task.AddFunc("0 0 */2 * *", func() {
		fmt.Println("task 6 for 1.5 hour:", time.Now().Local())
	}); err != nil {
		return err
	}

	task.Start()
	fmt.Println(task.Entries())
	return nil
}
