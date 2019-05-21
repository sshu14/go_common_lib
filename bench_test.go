package main

import (
	"fmt"
	"testing"
)

// 压力测试文件

func BenchmarkTest_thread_pool(t *testing.B) {
	// pool := thread_pool.NewThreadPool(3, t.N)
	for i := 0; i < t.N; i++ {
		// pool.AddTask(func() error {
		fmt.Println("hello")
		// return nil
		// })
	}
	//
	// pool.SetCallBack(func() {
	// 	fmt.Println("call back")
	// 	pool.Stop()

	// 	t.Log("测试通过")
	// })

	// pool.Start()
}
