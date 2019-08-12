// 华为软件特战队编程测试
// 秒杀
// 某秒内的第一人免单;同一时刻下单的人同样免单
// 输入:
// 		n行下单时间
// 输出:
// 		免单人数
// 示例1
// 	输入
// 		2019-01-01 00:00:00.001
// 		2019-01-01 00:00:00.002
// 		2019-01-01 00:00:00.003
// 	输出
// 		1
// 示例2
// 	输入
// 		2019-01-01 08:59:00.123
// 		2019-01-01 08:59:00.123
// 		2018-12-28 10:08:00.999
// 	输出
// 		3
package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	var lst_time []string
	for {
		var date, t string
		n, _ := fmt.Scanln(&date, &t)
		if n == 0 {
			break
		}
		lst_time = append(lst_time, date+" "+t)
	}

	sort.Strings(lst_time)
	var times []time.Time
	for _, s_time := range lst_time {
		t, _ := time.Parse("2006-01-02 15:04:05.999", s_time)
		times = append(times, t)
	}

	var n int
	var old_time time.Time
	for _, t := range times {
		if t.Unix() != old_time.Unix() {
			old_time = t
			n++
		} else if t == old_time {
			n++
		}
	}
	fmt.Println(n)
}
