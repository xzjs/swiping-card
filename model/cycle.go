package model

import (
	"time"

	"gorm.io/gorm"
)

type Cycle struct {
	gorm.Model
	Name string `json:"name"`
}

func (cycle *Cycle) GetDays() int {
	days := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	now := time.Now()
	year := int(now.Year())
	month := int(now.Month())
	total := 365
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		days[2]++
		total++
	}
	switch cycle.Name {
	case "天":
		return 1
	case "周":
		return 8 - int(now.Weekday())
	case "月":
		return days[month] - now.Day() + 1
	case "季":
		switch month {
		case 1:
			return days[1] + days[2] + days[3] - now.Day() + 1
		case 2:
			return days[2] + days[3] - now.Day() + 1
		case 3:
			return days[3] - now.Day() + 1
		case 4:
			return days[4] + days[5] + days[6] - now.Day() + 1
		case 5:
			return days[5] + days[6] - now.Day() + 1
		case 6:
			return days[6] - now.Day() + 1
		case 7:
			return days[7] + days[8] + days[9] - now.Day() + 1
		case 8:
			return days[8] + days[9] - now.Day() + 1
		case 9:
			return days[9] - now.Day() + 1
		case 10:
			return days[10] + days[11] + days[12] - now.Day() + 1
		case 11:
			return days[11] + days[12] - now.Day() + 1
		case 12:
			return days[12] - now.Day() + 1
		}
	case "年":
		return total - now.YearDay() + 1
	}
	return 1
}
