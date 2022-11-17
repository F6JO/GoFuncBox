package GoFuncBox

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)





func PrintCmdOutput(comm string, proc_line func(str string)) string{
	cmd := exec_comm(comm)
	cmd.Stdin = os.Stdin

	var wg sync.WaitGroup
	wg.Add(2)
	//捕获标准输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("INFO:", err)
		os.Exit(1)
	}
	readout := bufio.NewReader(stdout)

	jieguo := ""

	go func() {
		defer wg.Done()
		getOutput(readout, proc_line,&jieguo)
	}()

	//捕获标准错误
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	readerr := bufio.NewReader(stderr)

	go func() {
		defer wg.Done()
		getOutput(readerr, proc_line, &jieguo)
	}()

	////执行命令
	err = cmd.Run()
	if err != nil {
		return ""
	}

	wg.Wait()
	return jieguo



}


func getOutput(reader *bufio.Reader, proc_line func(str string), s *string) {
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		proc_line(string(line))
		*s += string(line) + "\n"
	}


}


func exec_comm(comm string) *exec.Cmd {
	str_list := strings.Split(comm," ")
	strings := []string{};
	startName := "";
	for i,str := range str_list{
		if  i == 0{
			startName = str;
			continue;
		}
		strings = append(strings, str)
	}
	return exec.Command(startName,strings...)

}