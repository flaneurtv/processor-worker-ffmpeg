package worker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/242617/flaneurtv/utils"
)

func Init(prgCh chan int, errCh chan error) error {
	http.HandleFunc("/progress", func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		chunks := [][]string{}
		var chunk []string
		buf := bufio.NewReader(r.Body)

		for {
			var barr []byte
			var err error
			if barr, _, err = buf.ReadLine(); err == io.EOF {
				return
			} else if err != nil {
				errCh <- err
				break
			}
			str := string(barr)
			if strings.Contains(str, "progress=") {
				var seconds int
				if seconds, err = utils.ParseOutTime(strings.Join(chunk, " ")); err != nil {
					log.Println(err)
					fmt.Println("err == io.EOF", err == io.EOF)
					errCh <- err
					return
				}
				prgCh <- seconds

				chunks = append(chunks, chunk)
				chunk = []string{}
			} else {
				chunk = append(chunk, str)
			}
		}

	})
	return http.ListenAndServe(":8080", nil)
}
