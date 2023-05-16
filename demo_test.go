package cron

import (
	"fmt"
	"testing"
	"time"
)

var milliSecondParser = NewParser(MilliSecond | CentiSecond | DeciSecond | Second | Minute | Hour | Dom | Month | DowOptional | Descriptor)

// 每10毫秒调度一次
func Test1(t *testing.T)  {
	cron := newWithMilliSeconds()

	cron.AddFuncWithType("0 * * * * * * * *", func() { fmt.Println("hello world")},LOOP_JOB)

	cron.Start()
	defer cron.Stop()
	<-time.After(time.Second*1)
}

// 每个月15日19点55分55秒270
func Test2(t *testing.T) {
	cron := newWithMilliSeconds()

	cron.AddFuncWithType("0 7 2 55 55 19 15 * *", func() { fmt.Println("hello world")},LOOP_JOB)

	cron.Start()
	defer cron.Stop()
	<-time.After(time.Hour)
}


//单次任务
func TestOnce(t *testing.T) {
	cron := newWithMilliSeconds()

	cron.AddFuncWithType("0 7 2 * * * * * *", func() { fmt.Println("hello world")},ONCE_JOB)

	cron.Start()
	defer cron.Stop()
	<-time.After(time.Second)
}

// newWithSeconds returns a Cron with the seconds field enabled.
func newWithMilliSeconds() *Cron {
	return New(WithParser(milliSecondParser))
}