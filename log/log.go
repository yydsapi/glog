package log

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var DebugLevel = 10

func Info(arg ...any) {
	if DebugLevel > 4 { //5
		s := buildString(" [INFO] ", arg)
		//fmt.Printf("%c[1;32m%s%c[0m\n", 0x1B, s, 0x1B)
		fmt.Printf("%c[1;38m%s%c[0m\n", 0x1B, s, 0x1B) //38 white
	}
}

func Debug(arg ...any) { //4
	if DebugLevel > 3 {
		s := buildString(" [DEBUG] ", arg)
		fmt.Printf("%c[1;34m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Warn(arg ...any) { //3
	if DebugLevel > 2 {
		s := buildString(" [WARN]  ", arg)
		fmt.Printf("%c[1;33m%s%c[0m\n\n", 0x1B, s, 0x1B)
	}
}

func Error(arg ...any) { // 2
	if DebugLevel > 1 {
		//pc, file, line, _ := runtime.Caller(1) //1
		//fmt.Println(file+", "+strconv.Itoa(line)+": "+runtime.FuncForPC(pc).Name()+" --------> ", err)
		s := buildString(" [ERROR] ", arg)
		fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Panic(arg ...any) { //1
	if DebugLevel > 0 {
		s := buildString(" [ERROR] ", arg)
		fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
		//Panic(arg)
	}
}
func Fatal(arg ...any) { //defer不执行
	if DebugLevel > -1 {
		s := buildString(" [FATAL] ", arg)
		log.Fatalf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func PanicRecord(arg ...any) {
	s := buildString(" [ERROR] ", arg)
	fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
	//Panic(arg)
}

func buildString(level string, args []any) string {
	var tag []any
	//tag = append(tag, time.Now().Format("2006-01-02 15:04:05.000000"), level, getPosition(), " --> ")
	tag = append(tag, time.Now().Format("15:04:05.000000"), level, strings.ReplaceAll(getPosition(), "D:/Work/go/yyds/", ""), " --> ")
	s := fmt.Sprint(tag...) + fmt.Sprint(args...)
	return s
}

func getPosition() string {
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	//path := strings.Split(file, "/")
	//index := len(path) - 1
	//return path[index] + ":" + strconv.Itoa(line)
	return strings.Replace(file, "D:/Work/go/vuepd/", "", 1) + ":" + strconv.Itoa(line)
}
func LogErr(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("\033[31m\033[1m") // 开始以红色打印 32绿色33黄色34蓝色35紫红色36青蓝色37白色
	fmt.Println(file+", "+strconv.Itoa(line)+": "+runtime.FuncForPC(pc).Name()+" --------> ", arg)
	fmt.Printf("\033[0m")
}
func LogErrPos(pos int,arg ...any) {
	pc, file, line, _ := runtime.Caller(pos)
	fmt.Printf("\033[31m\033[1m") // 开始以红色打印 32绿色33黄色34蓝色35紫红色36青蓝色37白色
	fmt.Println(file+", "+strconv.Itoa(line)+": "+runtime.FuncForPC(pc).Name()+" --------> ", arg)
	fmt.Printf("\033[0m")
}
