package main

import (
	"encoding/json"
	"fmt"
	"github.com/alzaburetz/myrestAPI/models"
	_ "github.com/alzaburetz/myrestAPI/models"
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
		matched, _ := regexp.Match(res.App,[]byte(name))

		if matched && res.Rule == "Close" {
			fmt.Printf("%s == %s = %v\n",name, res.App, matched)
			fmt.Println("app is closable")
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
		} else {
			log.Println("Successfully killed " + res.App)
		}

	}
	return err
}

func Openable(res *models.Restriction, nm []*process.Process) error {
	var i int
	var err error
	var name string
	for i = 0; i < len(nm); i++ {
		name, err = nm[i].Name()
		matched, _ := regexp.Match(res.App, []byte(name))
		if matched {
			break
		}
	}

	if i == len(nm) && res.Rule == "Open" {
		if TheDay(weekend, time.Now().Weekday().String()) && res.Time == "weekend" || TheDay(working, time.Now().Weekday().String()) && res.Time == "working" {
			go exec.Command(res.Exec).Run()
			cmd, _ := nm[i-1].Status()
			go log.Printf("Process %s status: %s\n", res.App, cmd)

		}
	}

	return err

}


func main() {
	_ , err := os.Stat("logs")
	if os.IsNotExist(err) {
		os.Mkdir("logs",0777)
	}
	var lo, _ = os.Create("logs/log "+time.Now().String()+".log")
	defer lo.Close()
	w := io.Writer(lo)
	log.SetOutput(w)
	for {
		procs, _ := process.Processes()
		request, _ := http.Get("http://localhost:3000/")

		byteVal, _ := ioutil.ReadAll(request.Body)
		var r models.Restrictions
		json.Unmarshal(byteVal, &r)
		//fmt.Printf("%v",r)
		for i := 0; i < len(r.Restrict); i++ {
			if time.Now().Hour() > r.Restrict[i].HF && time.Now().Hour() < r.Restrict[i].HT {
				err := Closable(&r.Restrict[i], procs)
				if err != nil {
					log.Printf("Error closing restricted app: %+v", err)
					break
				}
				err = Openable(&r.Restrict[i], procs)
				if err != nil {
					log.Printf("Error opening restricted app: %+v", err)
					break
				}
			}

		}

		time.Sleep(time.Second)
	}

}

type Restrict struct {
	Restrictions []Restriction `json:"restrictions"`
}

type Restriction struct {
	App      string `json:"app"`
	Windows  int    `json:"windows"`
	Rule     string `json:"rule"`
	Time     string `json:"time"`
	Hourfrom int    `json:"hourfrom"`
	Hourto   int    `json:"hourto"`
	Exec     string `json:"executable"`
}
