package task

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hnnngn/take-exam/pkg/client"
)

const TimeFormat = "2006-01-02"

var monthNames = [...]string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"}

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

func (t *Task) GetDate() string {
	year, month, day := t.CompleteBefore.Date()
	monthName := monthNames[month-1]

	return fmt.Sprintf("%s %s %s", strconv.Itoa(year), monthName, strconv.Itoa(day))
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
		var ctime string

		err = res.Scan(&task.Id, &task.BindedTo, &ctime, &task.IsCompleted, &task.Name, &task.Description)

		if err != nil {
			log.Fatal(err)
		}

		t, err := time.Parse(TimeFormat, ctime)

		if err != nil {
			log.Fatal(err)
		}

		task.CompleteBefore = t

		tasks = append(tasks, task)
	}

	return
}
