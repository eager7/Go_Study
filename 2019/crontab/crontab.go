package crontab

import (
	"fmt"
	"github.com/robfig/cron"
	"os/exec"
	"time"
)

func Initialize() error {
	_, _ = runCmd(fmt.Sprintf("date >> /tmp/crontab_test"))
	task := cron.New()
	//秒里有1时执行任务，相当于每分钟
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
		msg := fmt.Sprintf("touch /tmp/task5-%s", time.Now().Local())
		fmt.Println(msg)
		_, _ = runCmd(fmt.Sprintf("date >> /tmp/crontab_test"))
	}); err != nil {
		return err
	}
	if err := task.AddFunc("0 0 */2 * *", func() {
		msg := fmt.Sprintf("touch /tmp/task6-%s", time.Now().Local())
		fmt.Println(msg)
		_, _ = runCmd(fmt.Sprintf("date >> /tmp/crontab_test"))
	}); err != nil {
		return err
	}
	if err := task.AddFunc("0 0 1 * *", func() {
		msg := fmt.Sprintf("touch /tmp/task7-%s", time.Now().Local())
		fmt.Println(msg)
		_, _ = runCmd(fmt.Sprintf("date >> /tmp/crontab_test"))
	}); err != nil {
		return err
	}
	if err := task.AddFunc("0 0 2 * *", func() {
		msg := fmt.Sprintf("touch /tmp/task8-%s", time.Now().Local())
		fmt.Println(msg)
		_, _ = runCmd(fmt.Sprintf("date >> /tmp/crontab_test"))
	}); err != nil {
		return err
	}
	task.Start()
	fmt.Println(task.Entries())
	return nil
}

func runCmd(shell string) ([]byte, error) {
	cmd := exec.Command("/bin/bash", "-c", shell)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("exec ", cmd.Args, "failed, err:", err.Error(), "result:", string(out))
		return nil, err
	}
	//log.Notice("exec [", num, "]", cmd.Args, "success")
	return out, err
}