package main

import (
	"fmt"
	"sync"
)

func main() {
	//io操作、锁，最后需要关闭，解锁
	//类似try的finally
	var mu sync.Mutex
	mu.Lock()
	// defer里的代码由编译器在编译放到代码段{}最后
	defer mu.Unlock()
	//mu.Unlock()

	//多个defer以栈的顺序执行，先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	//因为返回值定义了变量名r，所以这里return b的b会被赋值给r，
	//此时r=b=10
	//然后执行defer，10++=11，11>10，所以最后r=11x11=121，return 121
	res := func() (r int) {
		r = 5
		b := 10
		defer func() {
			r++
			if r > 10 {
				r = r * r
			}
		}()
		return b
	}()
	fmt.Println(res)
}
