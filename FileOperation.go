package GoFuncBox

import (
	"errors"
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func Py_Writefile(fileName string, nr string, fangshi string) error {
	var f *os.File
	var err error
	if fangshi == "w" {
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	} else if fangshi == "a" {
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	}else {return errors.New("error a/w")}
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(nr), n)
		defer f.Close()
	}
	return nil
}

