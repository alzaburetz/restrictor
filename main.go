package main

import (
	"encoding/json"
	"github.com/alzaburetz/myrestAPI/models"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/process"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"time"
)

var working = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
var weekend = []string{"Saturday", "Sunday"}

func TheDay(days []string, day string) bool {
	result := false
	for i := 0; i < len(days); i++ {
		if days[i] == day {
			result = true
			break
		}
	}
	return result
}

func Closable(res *models.Restriction, nm []*process.Process) error {
	result := false
	var err error
	var i int
	for i = 0; i < len(nm); i++ {
		name, _ := nm[i].Cmdline()
		matched, _ := regexp.Match(res.App, []byte(name))

		if matched && res.Rule == "Close" {
			if TheDay(working, time.Now().Weekday().String()) && res.Time == "working" {
				result = true
				break
			}
			if TheDay(weekend, time.Now().Weekday().String()) && res.Time == "weekend" {
				result = true
				break
			}
		}

	}
	if result {
		err = nm[i].Kill()
		if err != nil {
			return err
		}

		log.Println("Successfully killed " + res.App)

	}
	return err
}

func Openable(res *models.Restriction, nm []*process.Process) error {
	var i int
	var err error
	var name string
	for i = 0; i < len(nm); i++ {
		name, err = nm[i].Name()
		matched, _ := regexp.Match(name, []byte(res.Exec))
		if matched {
			break
		}
	}

	if i == len(nm) && res.Rule == "Open" {
		if TheDay(weekend, time.Now().Weekday().String()) && res.Time == "weekend" || TheDay(working, time.Now().Weekday().String()) && res.Time == "working" {

			cmd, _ := nm[i-1].Status()
			go func() {
				err = exec.Command(res.Exec).Run()
				if err != nil {
					log.Println(err)
				}
			}()
			go log.Printf("Process %s status: %s\n", res.App, cmd)

		}
	}

	return err

}

func main() {
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		os.Mkdir("logs", 0777)
	}
	var lo, _ = os.Create("logs/log " + time.Now().String() + ".log")
	defer lo.Close()
	w := io.Writer(lo)
	log.SetOutput(w)
	hosts, _ := host.Users()
	host := hosts[0].User
	log.Println(host)
	for {
		procs, _ := process.Processes()
		request, err := http.Get("http://localhost:3000/restrictions/user/" + host)
		if err != nil || request == nil {
			log.Printf("Error reading from host %v\n", err)
		} else {

			defer func() {
				err := request.Body.Close()
				if err != nil {
					log.Println("Error closing request")
				}
			}()
			byteVal, err := ioutil.ReadAll(request.Body)
			if err != nil {
				log.Printf("Error reading response from host %v\n", err)
			}
			var r []models.Restriction
			err = json.Unmarshal(byteVal, &r)
			if err != nil {
				log.Println("Error marshaling response")
			}
			for i := 0; i < len(r); i++ {
				if time.Now().Hour() > r[i].HF && time.Now().Hour() < r[i].HT {
					err := Closable(&r[i], procs)
					if err != nil {
						log.Printf("Error closing ed app: %+v", err)
						break
					}
					err = Openable(&r[i], procs)
					if err != nil {
						log.Printf("Error opening ed app: %+v", err)
						break
					}
				}

			}
		}

		time.Sleep(time.Second)
	}

}
