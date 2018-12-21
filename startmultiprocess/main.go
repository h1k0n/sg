package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	f, err := os.OpenFile("/tmp/testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("This is a test log entry", os.Getpid(), os.Getppid())
	//bussiness
	pc := flag.Int("c", 0, "process count")
	process := flag.Int("p", 0, "process flag")
	sub := flag.Bool("sub", false, "sub process")
	flag.Parse()
	if *pc == 0 && *process == 0 {
		log.Println("*pc == 0 && *process == 0")
		os.Exit(-1)
	}
	var i int
	if *sub == false {
		path, err := os.Executable()
		if err != nil {
			log.Println("os.Executable()", err)
			os.Exit(1)
		}
		log.Println("path", path)
		attr := &os.ProcAttr{
			Files: []*os.File{
				os.Stdin,
				os.Stdout,
				os.Stderr,
				//f,
			},
			// Dir: ".",
			// Env: env,
		}
		//args := append([]string{path}, os.Args[1:]...)
		for ii := 1; ii <= *pc; ii++ {
			args := append([]string{path}, os.Args[1:]...)
			args = append(args, "-p", strconv.Itoa(ii), "-sub")
			process, err := os.StartProcess(
				path,
				args,
				attr,
			)
			if err != nil {
				log.Println("os.StartProcess()", err)
				os.Exit(1)
			}
			go process.Wait()
		}
		signals := make(chan os.Signal, 4)
		//在不handle SIGINT的情况下，kill -SIGINT会导致整个程序退出
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		a := <-signals
		//close(signals)
		log.Println("signals", *process, a)
		//在signals channel还开着的并且不阻塞的情况下，起到了忽略信号的作用
		//所以return和exit起了很大作用
		os.Exit(1)
		//return
	}
	if *process != 0 {
		i = *process
	}
	if *process%2 == 1 {
		//                log.Println("=",*process, os.Getpid(),os.Getppid())
		//               p,_:=os.FindProcess(os.Getpid()) //ppid,shell所在进程
		//              p.Signal(syscall.SIGINT)
		log.Println("odd process exit", *process)
		//                time.Sleep(30*time.Second)
		//syscall.Kill(syscall.Getppid(), syscall.SIGINT)
		//               time.Sleep(time.Second)
		return
	}
	for {
		log.Println(os.Getpid(), i)
		time.Sleep(time.Second)
	}
}
