package main

import (
	"fmt"
	"testing"

	thread_pool "github.com/sshu14/go_common_lib/thread_pool"
)

func Test_thread_pool(t *testing.T) {
	pool := thread_pool.NewThreadPool(3, 3)
	pool.AddTask(func() error {
		fmt.Println("hello")
		return nil
	})
	pool.AddTask(func() error {
		fmt.Println("world")
		return nil
	})
	pool.AddTask(func() error {
		fmt.Println("！")
		return nil
	})

	pool.SetCallBack(func() {
		fmt.Println("call back")
		pool.Stop()

		t.Log("测试通过")
	})

	pool.Start()
}
