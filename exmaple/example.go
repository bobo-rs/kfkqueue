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
