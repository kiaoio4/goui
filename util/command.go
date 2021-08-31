package util

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cast"
)

var (
	Message     string
	ctx         context.Context
	cancel      context.CancelFunc
	ProcessDone bool
)

func MountCommand(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error:can not obtain stdout pipe for command: %s", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error:The command is err: %s", err)
	}

	outputBuf := bufio.NewReader(stdout)
	for {
		_, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				return fmt.Errorf("error :%s", err)
			}
			break
		}

	}

	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("wait: %s", err.Error())
	}
	return nil
}

func ExcueteCommand(command, message string) string {
	ctx, cancel = context.WithCancel(context.Background())
	ProcessDone = false
	go func(cancelFunc context.CancelFunc) {
		time.Sleep(time.Second * 4)
		cancelFunc()
		ProcessDone = true
	}(cancel)

	Command(ctx, command, message)
	return message
}
func ExcuteStopCommand() {
	go func(cancelFunc context.CancelFunc) {
		cancelFunc()
	}(cancel)
}

func read(ctx context.Context, wg *sync.WaitGroup, std io.ReadCloser, message string) {
	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			ProcessDone = false
			return
		default:
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString, message)
			Message = readString
		}
	}
}

func Command(ctx context.Context, cmd, message string) error {
	//c := exec.CommandContext(ctx, "cmd", "/C", cmd) // windows
	c := exec.CommandContext(ctx, "bash", "-c", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	// 因为有2个任务, 一个需要读取stderr 另一个需要读取stdout
	wg.Add(2)
	go func(m string) { read(ctx, &wg, stderr, m) }(message)
	go func(m string) { read(ctx, &wg, stdout, m) }(message)
	// 这里一定要用start,而不是run 详情请看下面的图
	err = c.Start()
	// 等待任务结束
	wg.Wait()
	return err
}

func DiskCommand(command string, mount map[string][]string, tablekey map[int][]string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error:can not obtain stdout pipe for command: %s", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error:The command is err: %s", err)
	}

	count := 0
	mapindex := 0
	keycheck := ""

	outputBuf := bufio.NewReader(stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				return fmt.Errorf("error :%s", err)
			}
			break
		}

		countSplit := strings.Split(string(output), "|")
		key := string(countSplit[0])
		if key == "none" {
			keycheck = key + cast.ToString(mapindex)
			mapindex++
		} else {
			keycheck = key
		}
		for k, v := range countSplit {
			if k != 0 {
				mount[keycheck] = append(mount[keycheck], string(v))
			}
		}
		tablekey[0] = append(tablekey[0], keycheck)
		count++

	}

	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("wait: %s", err.Error())
	}
	return nil
}
