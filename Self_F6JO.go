package gopool

import (
	"sync"
	"fmt"
	"bufio"
	"errors"
	"os"
	"strings"
)

type SimplePool struct {
	wg   sync.WaitGroup
	work chan func() //任务队列
}

func NewSimplePoll(workers int) *SimplePool {
	p := &SimplePool{
		wg:   sync.WaitGroup{},
		work: make(chan func()),
	}
	p.wg.Add(workers)
	//根据指定的并发量去读取管道并执行
	for i := 0; i < workers; i++ {
		go func() {
			defer func() {
				// 捕获异常 防止waitGroup阻塞
				if err := recover(); err != nil {
					fmt.Println(err)
					p.wg.Done()
				}
			}()
			// 从workChannel中取出任务执行
			for fn := range p.work {
				fn()
			}
			p.wg.Done()
		}()
	}
	return p
}
// 添加任务
func (p *SimplePool) Add(fn func()) {
	p.work <- fn
}

// 执行
func (p *SimplePool) Run() {
	close(p.work)
	p.wg.Wait()
}




func GetHeaders(filepath string) (map[string]string, error) {
	file,err := os.Open(filepath)
	defer file.Close()
	zidian := map[string]string{}
	if err != nil{
		return nil,errors.New(err.Error())
	}
	ioliu := bufio.NewReader(file)
	for {
		nr, err := ioliu.ReadString('\n')
		if err != nil{
			return zidian, nil
		}
		key,value := handle_row(nr)
		if key != "Host" {
			zidian[key] = value
		}

	}
}

func handle_row(row string) (string, string) {
	row = strings.Replace(row,"\n","",-1)
	sy := strings.Index(row,":")
	key := strings.TrimSpace(row[:sy])
	value := strings.TrimSpace(row[sy+1:])
	return key,value

}
