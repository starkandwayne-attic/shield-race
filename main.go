package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func Drain(prefix string, rd io.Reader) {
	b := bufio.NewScanner(rd)
	for b.Scan() {
		fmt.Printf("%s> %s\n", prefix, b.Text())
	}
}

func main() {
	/* <go> | ls -alh / | sort | <go> */
	ls := exec.Command("ls", "-alh", "/")
	sort := exec.Command("sort")
	/* grab standard error from both commands */
	err1, _ := ls.StderrPipe()
	err2, _ := sort.StderrPipe()

	/* wire up ls's stdout to sort's stdin */
	sort.Stdin, _ = ls.StdoutPipe()

	/* grab the output from the `sort' command */
	out, _ := sort.StdoutPipe()

	/* spin up two goroutines to stream errors */
	go Drain("error", err1)
	go Drain("error", err2)

	/* spin up a goroutine to stream output from `sort' */
	go Drain("output", out)

	_ = sort.Start()
	_ = ls.Start()

	_ = ls.Wait()
	_ = sort.Wait()
}
