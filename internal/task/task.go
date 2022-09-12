package task

import (
	"log"
	"strconv"
	"time"

	"github.com/hnnngn/take-exam/pkg/client"
)

const TimeFormat = "2006-01-02"

type Task struct {
	BindedTo       string
	CompleteBefore time.Time
	IsCompleted    bool
	Name           string
	Description    string
}

type ReceivedTask struct {
	Id int
	Task
}

func (t *Task) InsertIntoDatabase() {
	conn := client.Connect()
	defer conn.Close()

	res, err := conn.Query("INSERT INTO `tasks` (`BindedTo`, `CompleteBefore`, `IsCompleted`, `Name`, `Description`) VALUES ('" + t.BindedTo + "', '" + t.CompleteBefore.Format(TimeFormat) + "', '" + strconv.FormatBool(t.IsCompleted) + "', '" + t.Name + "', '" + t.Description + "')")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()
}

func GetFromDatabase(login string) (tasks []ReceivedTask) {
	conn := client.Connect()
	defer conn.Close()

	res, err := conn.Query("SELECT * FROM `tasks` WHERE BindedTo = '" + login + "'")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	for res.Next() {
		var task ReceivedTask

		err = res.Scan(&task.Id, &task.BindedTo, &task.IsCompleted, &task.CompleteBefore, &task.Name, &task.Description)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return
}
