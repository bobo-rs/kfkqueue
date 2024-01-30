package exmaple

import (
	"context"
	"fmt"
	"github.com/bobo-rs/mall-util/utli"
	"strconv"
	"testing"
)

func Test_SyncProducer(t *testing.T) {
	client := Example_ClientSync(1)
	ctx := context.Background()
	for i := 0; i < 10000; i++ {
		t := i
		fmt.Println(
			client.SyncProducer(ctx, struct {
				JobName string
				Value   interface{}
			}{JobName: "ExampleJob1", Value: struct {
				Name  string
				Value int
				Detail struct{
					Content string
					Arr []int
				}
			}{
				Name:  "张三" + strconv.Itoa(t),
				Value: t,
				Detail: struct {
					Content string
					Arr     []int
				}{Content: "不可能，绝对不可能" + strconv.Itoa(t), Arr: []int{t}},
			}}),
		)
	}
	fmt.Println()
}

func Test_AsyncProducer(t *testing.T) {
	client := Exmaple_ClientAsync()
	ctx := context.Background()
	for i := 0; i < 10000; i++ {
		t := i
		client.AsyncProducer(ctx, struct {
			JobName string
			Value   interface{}
		}{ JobName: "ExampleJob2", Value: struct {
			Name  string
			Value int
		}{
			Name:  "张三" + strconv.Itoa(t),
			Value: t,
		}})
	}
}

func Test_Consumer(t *testing.T) {
	ctx := context.Background()
	fmt.Println(
		Exmaple_ClientQueue(QueueCfg{}).Consumer(ctx),
	)
}

func Test_Calculate(t *testing.T) {
	fmt.Println(util.Sub(88, "234324", 3))
}