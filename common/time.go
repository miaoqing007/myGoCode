package common

import "time"

const (
	Second = 1
	Minute = 60 * Second
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
)

//24时制
const (
	ZeroClock = iota
	OneClock
	TwoClock
	ThreeClock
	FourClock
	FiveClock
	SixClock
	SevenClock
	EightClock
	NineClock
	TenClock
	ElevenClock
	TwelveClock
	ThirteenClock
	FourteenClock
	FifteenClock
	SixteenClock
	EighteenClock
	NineteenClock
	TwentyClock
	Twenty_OneClock
	Twenty_twoClock
	Twenty_threeClock
)

// 两个时间是否是同一天
func IsSameDay(time1, time2 int64) bool {
	t1 := time.Unix(time1, 0)
	t2 := time.Unix(time2, 0)
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

//两个时间是否是同一周
func IsSameWeek(time1, time2 int64) bool {
	y1, w1 := time.Unix(time1, 0).ISOWeek()
	y2, w2 := time.Unix(time2, 0).ISOWeek()
	return y1 == y2 && w1 == w2
}

//计算两个时间戳相差的天数
func TimeSub(t1, t2 int64) int {
	t1 = GetTimeUnixMorning(t1)
	t2 = GetTimeUnixMorning(t2)
	return int((t2 - t1) / Day)
}

//是否是同一个月
func IsSameMonth(time1, time2 int64) bool {
	y1, m1, _ := time.Unix(time1, 0).Date()
	y2, m2, _ := time.Unix(time2, 0).Date()
	return y1 == y2 && m1 == m2
}

//获取整点时间戳
func GetNowHourUinx(hour int) int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location()).Unix()

}

//获取当前月1号的零点时间
func GetFirstDayMorningOfMonth() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Unix()
}

//获取下个月1号的零点时间
func GetFirstDayMorningOfNextMonth(nextmonth int) int64 {
	now := time.Now()
	a := int(now.Month()) + nextmonth
	addY := 0
	if a > 12 {
		addY = a / 12
	}
	nextM := int(now.Month()) + (nextmonth % 12)
	if nextM > 12 {
		nextM -= 12
	}
	nextY := now.Year() + addY
	return time.Date(nextY, time.Month(nextM), 1, 0, 0, 0, 0, now.Location()).Unix()
}

//获取当前周星期一的零点时间
func GetFirstDayMorningOfWeek() int64 {
	morning := GetNowHourUinx(ZeroClock)
	ms := morning - (int64(time.Now().Weekday()-time.Monday))*Day
	return ms
}

//这周指定星期几的零点（如果这周几未到 取上周的星期几的零点） 包括今天
func GetWeekWithConfig(weekOfDay int) int64 {
	now := time.Now()
	nowWeek := now.Weekday()
	if int(nowWeek) >= weekOfDay {
		return GetNowHourUinx(ZeroClock) - int64((int(nowWeek)-weekOfDay)*Day)
	} else {
		return GetNowHourUinx(ZeroClock) + int64((weekOfDay-int(nowWeek))*Day) - Week
	}
}

//获取星期几的整点整分时间戳
func GetWeekHourMinUnix(week, hour, min int) int64 {
	now := time.Now()
	nowWeek := now.Weekday()
	if int(nowWeek) >= week {
		return time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, now.Location()).Unix() - int64((int(nowWeek)-week)*Day)
	} else {
		return time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, now.Location()).Unix() + int64((week-int(nowWeek))*Day)
	}
}

//获取整点整分的时间戳
func GetHourMinUnix(hour, min int) int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, now.Location()).Unix()
}

//获取传入时间戳当天的零点
func GetTimeUnixMorning(t int64) int64 {
	year, month, day := time.Unix(t, 0).Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location()).Unix()
}
