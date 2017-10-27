package worker

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/242617/flaneurtv/utils"
)

var ErrWorkerBusy = errors.New("worker is busy")

func Process(startCh chan struct{}, durationCh chan int, job *Job) (stdout []byte, stderr []byte, err error) {
	if err = start(job); err != nil {
		return
	}
	defer stop()

	args := append(CurrentJob.Arguments, "-progress", "http://localhost:8080/progress")
	cmd := exec.Command("ffmpeg", args...)
	cmd.Env = os.Environ()

	var stdOut, stdErr io.ReadCloser
	if stdOut, err = cmd.StdoutPipe(); err != nil {
		log.Println(err)
		return
	}
	if stdErr, err = cmd.StderrPipe(); err != nil {
		log.Println(err)
		return
	}

	stderrBuf := bytes.NewBuffer([]byte{})

	go func() {
		errBuf := bufio.NewReader(stdErr)
		for {
			var barr []byte
			if barr, _, err = errBuf.ReadLine(); err == io.EOF {
				err = nil
				return
			} else if err != nil {
				log.Println(err)
				return
			} else {
				stderrBuf.Write(barr)
				str := string(barr)
				if strings.Contains(str, "Duration:") {
					var duration int
					if duration, err = utils.ParseDuration(string(barr)); err != nil {
						log.Println(err)
						return
					}
					durationCh <- duration
				}
			}
		}
	}()

	if err = cmd.Start(); err != nil {
		return
	}
	startCh <- struct{}{}

	if err = cmd.Wait(); err != nil {
		return
	}

	stdout, _ = ioutil.ReadAll(stdOut)
	stderr, _ = ioutil.ReadAll(stderrBuf)

	return
}
