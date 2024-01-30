package exmaple

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	ExampleJobs1 struct {
	}
	ExampleJobs2 struct {
	}

	Job1ValueItem struct {
		Name string
		Value int
		Detail struct{
			Content string
			Arr []int
		}
	}
)

type Jobs struct {
}

func (j *ExampleJobs1) Execute(ctx context.Context, value interface{}) {
	fmt.Println("我是Example的Job1任务")
	fmt.Println(value)

	// 数据
	//item := value.(map[string]interface{})
	//fmt.Println(item)

	data := Job1ValueItem{}
	_ = gconv.Struct(value, &data)
	fmt.Println(data)
}

func (j *ExampleJobs2) Execute(ctx context.Context, value interface{}) {
	fmt.Println("我是Example的Job2任务")
	fmt.Println(value)


	data := Job1ValueItem{}
	_ = gconv.Struct(value, &data)
	fmt.Println(data)
}
