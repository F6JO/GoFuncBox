package GoFuncBox

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetHeaders(filepath string) (map[string]interface{}, error) {
	file,err := os.Open(filepath)
	defer file.Close()
	zidian := map[string]interface{}{}
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
		if key != "Host" && key != "Cookie"{
			zidian[key] = value
		}else if key == "Cookie" {
			zidian2 := map[string]string{}
			list1 := strings.Split(value,";")
			for _,a := range list1{
				list2 := strings.Split(a,"=")
				jian := strings.TrimSpace(list2[0])
				zhi := ""
				for _,x := range list2[1:]{
					zhi += "="+strings.TrimSpace(x)
				}
				zidian2[jian] = strings.Trim(zhi,"=")
			}

			zidian[key] = zidian2

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
