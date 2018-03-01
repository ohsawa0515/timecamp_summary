package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const host = "www.timecamp.com"

type message struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Entry struct {
	Id          string `json:"id"`
	Duration    string `json:"duration"`
	UserId      string `json:"user_id"`
	Description string `json:"description"`
	LastModify  string `json:"last_modify"`
	TaskId      string `json:"task_id"`
	Date        string `json:"date"`
	StartTime   string `json:"start_time"`
	Name        string `json:"name"`
}

type Entries []Entry

type TaskSummary struct {
	Name     string
	Duration int64
}

func timeCamp(m message) (string, error) {
	if len(m.From) == 0 {
		now := time.Now().Format("2006-01-02")
		m.From = now
	}
	if len(m.To) == 0 {
		m.To = m.From
	}

	apiToken, ok := os.LookupEnv("TIMECAMP_TOKEN")
	if !ok {
		return "", fmt.Errorf("set environment variable `TIMECAMP_TOKEN`")
	}
	userId, ok := os.LookupEnv("TIMECAMP_USER_ID")
	if !ok {
		return "", fmt.Errorf("set environment variable `TIMECAMP_USER_ID`")
	}

	u := &url.URL{}
	u.Scheme = "https"
	u.Host = host
	u.Path = "third_party/api/entries/format/json/api_token/" + apiToken + "/from/" + m.From + "/to/" + m.To + "/user_ids/" + userId
	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	entries := Entries{}
	if err := json.Unmarshal(body, &entries); err != nil {
		return "", err
	}

	var summary = map[string]*TaskSummary{}
	var total int64
	for _, entry := range entries {
		if summary[entry.TaskId] == nil {
			summary[entry.TaskId] = &TaskSummary{
				Name:     entry.Name,
				Duration: 0,
			}
		}
		d, _ := strconv.ParseInt(entry.Duration, 10, 0)
		summary[entry.TaskId].Duration += d
		total += d
	}
	buf := &bytes.Buffer{}
	for _, task := range summary {
		buf.WriteString(fmt.Sprintf("%s, %s\n", task.Name, time.Duration(task.Duration)*time.Second))
	}
	buf.WriteString(fmt.Sprintf("Total: %s", time.Duration(total)*time.Second))
	return buf.String(), nil
}

func main() {
	var m message
	flag.StringVar(&m.From, "from", "", "The beginning of the date(YYYY-MM-DD). The default is today.")
	flag.StringVar(&m.To, "to", "", "End of the date(YYYY-MM-DD). The default is the same day as from option.")
	flag.Parse()
	output, err := timeCamp(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
