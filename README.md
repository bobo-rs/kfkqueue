# kfkqueue 消息队列，用于操作Kafka消息中间件


## 简介
~~~
1. kfkqueue包，是基于[sarama]二次封装的Kafka客户端包，通过对[sarama]消费端、生产端进行二次封装，可以使用者快捷开发，无需再去学习了解[sarama]
   源码和实现逻辑，只需通过调用封装好的方法，就可以快捷使用
2. 快捷使用方法包含：SyncProducer（同步生产消息）、AsyncProducer（异步生产消息）、Consumer（消费者），根据不同的场景进行调用
3. 支持单机和集群部署的Kafka下使用
~~~


## 安装
~~~
1.可直接使用Go管理包进行安装：go get -u github.com/bobo-rs/kfkqueue
~~~

## 使用
### 生产消息使用示例
```go
package exmaple

import "kfkqueue"

type QueueCfg struct {
	// 配置Kafka地址
	Addrs []string

	// 应答模式：0无需发送消息后返回应答结果【异步模式】，1.发送消息后只需Leader应答，-1.发送消息后需要Leader和追随者都应答
	RequiredAck int8

	// 发送消息后把成功的信息通过channel管道返回：false不返回，true返回（适合1和-1模式）
	Success bool

	// 发送消息后把失败的信息通过channel管道返回：false不返回，true返回
	Error bool

	// 分区主题，不传默认自动分配
	Topic string

	// 分区主题先每个消息的KEY，不传默认自动分配
	Key string

	// 分区主题组
	Group string
}

var (
	Addrs = []string{"127.0.0.1:9092"}
	Topics = []string{"queue_example_sync_topic", "queue_example_async_topic"}
	Group = "queue_example_topic_group"
)

func Exmaple_ClientQueue(cfg QueueCfg) *kfkqueue.SQueue {
	return kfkqueue.New(kfkqueue.SQueue{
		Addrs:       Addrs,
		RequiredAck: cfg.RequiredAck,
		Success:     cfg.Success,
		Error:       cfg.Error,
		Topic:       cfg.Topic,
		Topics: Topics,
		Group:       Group,
		JobsMap: map[string]kfkqueue.Jobs{
			"ExampleJob": &ExampleJobs1{},
			"ExampleJob1": &ExampleJobs1{},
			"ExampleJob2": &ExampleJobs2{},
		},
	})
}

func Example_ClientSync(ack int8) *kfkqueue.SQueue {
	topic := "queue_example_sync_topic"
	return Exmaple_ClientQueue(QueueCfg{
		RequiredAck: ack,
		Topic:       topic,
		Key:         topic + "_key",
	})
}

func Exmaple_ClientAsync() *kfkqueue.SQueue {
	topic := "queue_example_async_topic"
	return Exmaple_ClientQueue(QueueCfg{
		Topic:       topic,
		Success: true,
		Error:   true,
		Key:         topic + "_key",
	})
}

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
```
### 消费者使用示例
```go
func Test_Consumer(t *testing.T) {
	ctx := context.Background()
	fmt.Println(
		Exmaple_ClientQueue(QueueCfg{}).Consumer(ctx),
	)
}
```

### 完整示例
~~~
exmaple：包
~~~

## 依赖
1. Kafka操作包：github.com/IBM/sarama
2. Goframe工具包：github.com/gogf/gf/v2