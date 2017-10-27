package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/242617/flaneurtv/message"
	"github.com/242617/flaneurtv/process"
	"github.com/242617/flaneurtv/utils"
	"github.com/242617/flaneurtv/worker"
)

var proc *process.Process

func main() {
	// log.SetFlags(log.Lshortfile)

	// Init
	proc = process.NewProcess()
	durationCh := make(chan int)
	startCh := make(chan struct{})
	prgCh := make(chan int)
	prgErrCh := make(chan error)

	// Retrieving info from stdin
	go func() {
		for {
			var msg message.Message
			if err := json.NewDecoder(os.Stdin).Decode(&msg); err != nil {
				Err(message.LevelError, err.Error())
				continue
			}

			if !msg.AddressedTo(proc) {
				Err(message.LevelInfo, "unknown recipient")
				continue
			}

			Err(message.LevelDebug, fmt.Sprintf("message: %s", utils.JSON(msg)))

			if t, err := msg.Type(); err != nil {
				Err(message.LevelInfo, err.Error())
				continue
			} else {

				switch t {
				case message.TypeTick:

					if worker.IsBusy() {
						Out(message.NewWorkerBusy(proc))
						continue
					}
					Out(message.NewWorkerIdle(proc, &msg))

				case message.TypeJobAssignment:

					if worker.IsBusy() {
						Err(message.LevelInfo, "worker is busy")
						continue
					}

					job := worker.Job{
						UUID:          msg.Payload.UUID,
						ReferenceUUID: msg.Payload.ReferenceUUID,
						QueueName:     msg.Payload.QueueName,
						Command:       msg.Payload.Command,
						Arguments:     msg.Payload.Arguments,
						CreatedAt:     msg.Payload.CreatedAt,
						StartedAt:     msg.Payload.StartedAt,
					}
					if _, _, err := worker.Process(startCh, durationCh, &job); err != nil {
						Err(message.LevelError, err.Error())
						continue
					}
					Out(message.NewJobFinished(proc))
					Out(message.NewWorkerIdle(proc, &msg))

				}

			}
		}
	}()

	// Retrieving info from worker
	go func() {
		for {
			select {
			case <-startCh:
				Out(message.NewJobAccepted(proc))
			case duration := <-durationCh:
				worker.SetDuration(duration)
				Err(message.LevelDebug, fmt.Sprintf("duration: %d", duration))
			case seconds := <-prgCh:
				worker.SetSeconds(seconds)
				Err(message.LevelDebug, fmt.Sprintf("seconds: %d", seconds))
				Out(message.NewWorkerBusy(proc))
			case err := <-prgErrCh:
				Err(message.LevelError, err.Error())
			}
		}
	}()

	log.Fatal(worker.Init(prgCh, prgErrCh))
}

func Out(msg *message.Message) {
	json.NewEncoder(os.Stdout).Encode(msg)
}

func Err(level string, text string) {
	json.NewEncoder(os.Stdout).Encode(message.NewError(level, text))
}
