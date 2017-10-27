package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/242617/flaneurtv/message"
	"github.com/242617/flaneurtv/process"
	"github.com/242617/flaneurtv/utils"
)

var queue = process.Process{
	Namespace:   "flaneur",
	ServiceUUID: "JOBQUEUE-1285-4E4C-A44E-JOBQUEUE0000",
	ServiceName: "micro-job-queue",
	ServiceHost: "job-queue",
}

var worker = process.Process{
	Namespace:   "flaneur",
	ServiceUUID: "workerID",
	ServiceName: "micro-worker-ffmpeg",
	ServiceHost: "ticker",
}

var worker2 = process.Process{
	Namespace:   "flaneur",
	ServiceUUID: "worker2ID",
	ServiceName: "micro-worker-ffmpeg",
	ServiceHost: "ticker",
}

func main() {
	log.SetFlags(log.Lshortfile)

	cmd := exec.Command("./worker-ffmpeg")
	var (
		stdin          io.Writer
		stdout, stderr io.Reader
		err            error
	)
	if stdin, err = cmd.StdinPipe(); err != nil {
		log.Fatal(err)
	}
	if stdout, err = cmd.StderrPipe(); err != nil {
		log.Fatal(err)
	}
	if stderr, err = cmd.StdoutPipe(); err != nil {
		log.Fatal(err)
	}
	cmd.Env = append(os.Environ(), worker.Environ()...)

	var wg sync.WaitGroup

	// Tracing stdout
	wg.Add(1)
	go func() {
		buf := bufio.NewReader(stdout)
		for {
			var barr []byte
			if barr, _, err = buf.ReadLine(); err != nil {
				log.Println(err)
				break
			}
			log.Println(utils.JSON(string(barr)))
		}
		wg.Done()
	}()

	// Tracing stderr
	wg.Add(1)
	go func() {
		buf := bufio.NewReader(stderr)
		for {
			var barr []byte
			if barr, _, err = buf.ReadLine(); err != nil {
				log.Println(err)
				break
			}
			log.Println(utils.JSON(string(barr)))
		}
		wg.Done()
	}()

	// Testing
	wg.Add(1)
	go func() {

		var msg *message.Message

		/*go func() {
			for {
				time.Sleep(10 * time.Second)
				msg = message.NewTick(&worker, queue.Namespace, NewTickPayload())
				log.Println(utils.JSON(msg))
				json.NewEncoder(stdin).Encode(msg)
			}
		}()*/

		// fake tick
		/*time.Sleep(time.Second)
		log.Println("---> fake tick")
		msg = message.NewTick(&worker2, queue.Namespace, NewTickPayload())
		log.Println(utils.JSON(msg))
		json.NewEncoder(stdin).Encode(msg)*/

		time.Sleep(time.Second)
		log.Println("---> first job assignment")
		msg = message.NewJobAssignment(&worker, queue.Namespace, NewJobPayload())
		log.Println(utils.JSON(msg))
		json.NewEncoder(stdin).Encode(msg)

		/*time.Sleep(time.Second)
		log.Println("---> second job assignment")
		msg = message.NewJobAssignment(&worker, queue.Namespace, NewJobPayload())
		log.Println(utils.JSON(msg))
		json.NewEncoder(stdin).Encode(msg)*/

		/*time.Sleep(time.Second)

		fmt.Println("---> tick")
		msg = message.NewTick(&info)
		if err := json.NewEncoder(os.Stdin).Encode(&msg); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)

		fmt.Println("---> second job assignment")
		msg = message.NewJobAssignment(&info, strings.Split("-i test.mp4 -s 400x300 -y output.mp4", " ")...)
		if err := json.NewEncoder(os.Stdin).Encode(&msg); err != nil {
			log.Fatal(err)
		}*/

		wg.Done()
	}()

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	wg.Wait()

}

func NewTickPayload() *message.Payload {
	return &message.Payload{
		UUID: utils.GenerateUUID(),
	}
}

func NewJobPayload() *message.Payload {
	return &message.Payload{
		UUID:          utils.GenerateUUID(),
		ReferenceUUID: utils.GenerateUUID(),
		QueueName:     "ffmpeg",
		Command:       "ffmpeg",
		Arguments:     strings.Split("-i test.mp4 -s 400x300 -y output.mp4", " "),
		CreatedAt:     utils.Now(),
	}
}
