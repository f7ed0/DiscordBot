package edt

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var Weekday []string = []string{
	"Lundi",
	"Mardi",
	"Mercredi",
	"Jeudi",
	"Vendredi",
	"Samedi",
	"Dimanche",
}

var Month []string = []string{
	"",
	"Janvier",
	"Février",
	"Mars",
	"Avril",
	"Mai",
	"Juin",
	"Juillet",
	"Août",
	"Septembre",
	"Octobre",
	"Décembre",
}

type EdtElement struct {
	Course         string
	Timing         string
	Room           string
	Teacher        string
	Group          string
	Id             string
	DayOfTheWeek   int
	StartTimeStamp time.Time
	EndTimeStamp   time.Time
}

func (edt EdtElement) String() string {
	return fmt.Sprintf(
		"%v : %v [%v]",
		edt.Timing,
		edt.Course,
		edt.Room,
	)
}

func (edt *EdtElement) generateTimeStamp(weekOffset int) (err error) {
	day := GetStartDayOfWeek(time.Now().Add(time.Duration(weekOffset) * 24 * 7 * time.Hour)).Add(time.Duration(edt.DayOfTheWeek) * 24 * time.Hour)
	split := strings.Split(edt.Timing, "-")
	t1 := strings.Split(split[0], ":")
	t2 := strings.Split(split[1], ":")
	h1, err := strconv.Atoi(t1[0])
	if err != nil {
		return
	}
	m1, err := strconv.Atoi(t1[1])
	if err != nil {
		return
	}
	h2, err := strconv.Atoi(t2[0])
	if err != nil {
		return
	}
	m2, err := strconv.Atoi(t2[1])
	if err != nil {
		return
	}
	edt.StartTimeStamp = day.Add((time.Duration(h1) * time.Hour)).Add(time.Duration(m1) * time.Minute)
	edt.EndTimeStamp = day.Add((time.Duration(h2) * time.Hour)).Add(time.Duration(m2) * time.Minute)
	return

}

func GetStartDayOfWeek(tm time.Time) time.Time { //get monday 00:00:00
	weekday := time.Duration(tm.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := tm.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour)
}
