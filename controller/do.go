package controller

import (
	"math/rand"
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"
	"time"

	"github.com/gin-gonic/gin"
)

func DoPost() error {
	var plans []model.Plan
	db := lib.DB()
	now := time.Now()
	result := db.Preload("Cycle").Where("finished = false AND start <= ? AND end >= ?", now, now).Find(&plans)
	if result.Error != nil {
		return result.Error
	}

	var ways []model.Way
	result = db.Find(&ways)
	if result.Error != nil {
		return result.Error
	}

	for _, plan := range plans {
		var dos []model.Do
		rand.Seed(time.Now().UnixNano())

		totalNum := 0
		totalMoney := 0
		money := 0

		// 如果频率为0，意味着计算总金额
		if plan.Frequency == 0 {
			result = db.Where("created_at >= ? AND created_at <= ?", plan.Start, plan.End).Find(&dos)
		} else {
			start, end := getTimeSection(plan.Cycle)
			result = db.Where("created_at >= ? AND created_at <= ?", start, end).Find(&dos)
		}
		if result.Error != nil {
			return result.Error
		}
		for _, do := range dos {
			if !do.Finished {
				continue //有任务未完成则不再分配此计划的任务
			}
			totalNum += 1
			totalMoney += do.Money
		}

		if plan.Frequency == 0 {
			money = rand.Intn(plan.Ceiling-plan.Floor) + plan.Floor
			if money+totalMoney > plan.Sum {
				money = plan.Sum - totalMoney
			}
		} else {
			probability := (plan.Frequency - totalNum) * 100 / plan.Cycle.GetDays()
			r := rand.Intn(101)
			if r > probability {
				continue
			}
			money = plan.Floor
		}

		// 随机选取一种支付方式
		index := rand.Intn(len(ways))
		way := ways[index]

		do := model.Do{
			UserID:   plan.UserID,
			PlanID:   plan.ID,
			WayID:    way.ID,
			Money:    money,
			Finished: false,
		}
		result := db.Create(&do)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func DoGet(c *gin.Context) {
	var dos []model.Do
	db := lib.DB()
	result := db.Where("finished = false").Find(&dos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, dos)
	}
}

func DoGetOne(c *gin.Context) {
	var do model.Do
	db := lib.DB()
	id := c.Param("id")
	result := db.First(&do, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, do)
	}
}

func DoPut(c *gin.Context) {
	do := model.Do{}
	if err := c.ShouldBindJSON(&do); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	userID := c.GetUint("userID")
	db := lib.DB()
	var doDB model.Do
	db.Where("user_id = ?", userID).First(&doDB, id)
	doDB.Money = do.Money
	doDB.WayID = do.WayID
	doDB.Finished = do.Finished
	db.Save(&doDB)
	c.JSON(http.StatusOK, "OK")
}

func DoDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	db.Delete(&model.Do{}, id)
	c.JSON(http.StatusOK, "OK")
}

// 根据周期获取日期的开始和结束
func getTimeSection(cycle model.Cycle) (start time.Time, end time.Time) {
	now := time.Now()
	switch cycle.Name {
	case "周":
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6
		}
		start = now.AddDate(0, 0, offset)
		end = now.AddDate(0, 0, 6+offset)
	case "月":
		start = now.AddDate(0, 0, -now.Day()+1)
		end = now.AddDate(0, 1, -1)
	case "季":
		year := time.Now().Format("2006")
		month := int(now.Month())
		var firstOfQuarter string
		var lastOfQuarter string
		if month >= 1 && month <= 3 {
			//1月1号
			firstOfQuarter = year + "-01-01 00:00:00"
			lastOfQuarter = year + "-03-31 23:59:59"
		} else if month >= 4 && month <= 6 {
			firstOfQuarter = year + "-04-01 00:00:00"
			lastOfQuarter = year + "-06-30 23:59:59"
		} else if month >= 7 && month <= 9 {
			firstOfQuarter = year + "-07-01 00:00:00"
			lastOfQuarter = year + "-09-30 23:59:59"
		} else {
			firstOfQuarter = year + "-10-01 00:00:00"
			lastOfQuarter = year + "-12-31 23:59:59"
		}
		start, _ = time.ParseInLocation("2006-01-02 15:04:05", firstOfQuarter, time.Local)
		end, _ = time.ParseInLocation("2006-01-02 15:04:05", lastOfQuarter, time.Local)
	}
	return start, end
}
