package main

import (
	"fmt"
	"errors"
	"os"
)

//日志器写入接口
type LogWriter interface {
	WriteLog(data interface{}) error
}

//日志器
type Logger struct {
	writerList []LogWriter
}

//日志器注册
func (l *Logger) Register(lw LogWriter) {
	l.writerList = append(l.writerList, lw)
}

//日志器写入
func (l *Logger) Log(data interface{}) {
	for _, w := range l.writerList {
		w.WriteLog(data)
	}
}

//=========================================文件写入==================================================
type FileWriter struct{
	file *os.File
}

//日志文件创建
func (f *FileWriter) SetFile(fileName string) (err error) {
	
	if f.file != nil {
		f.file.Close()
	}

	f.file, err = os.Create(fileName)

	return err
}

//日志写入文件
func (f *FileWriter) WriteLog(data interface{}) error {

	if f.file == nil {
		return errors.New("file no create")
	}

	str := fmt.Sprintf("%v\n", data)

	_, err := f.file.Write([]byte(str))

	return err
}

//=========================================命令行写入==================================================
type ConsoleWriter struct{}

func (c *ConsoleWriter) WriteLog(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	
	_, err := os.Stdout.Write([]byte(str))

	return err
}

//=========================================初始化==================================================
func createLog() *Logger {
	l := &Logger{}
	f := &FileWriter{}
	c := &ConsoleWriter{}

	if err := f.SetFile("123.txt"); err != nil {
		fmt.Println(err)
	}

	l.Register(f)
	l.Register(c)

	return l
}

func main() {
	l := createLog()
	l.Log("测试")
}

