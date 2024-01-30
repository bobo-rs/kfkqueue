package kfkqueue

import "context"

type (
	// SQueue 队列链式属性管理
	SQueue struct {
		// 应答模式：0无需消息发送之后应答，1发送消息Leader应答成功即可，-1 发送消息主从都需要应答成功
		RequiredAck int8

		// 成功交付的消息将在success channel返回：false否，true是
		Success bool

		// 失败交付的消息将在Error channel返回：false否，true是
		Error bool

		// 消息分区主题
		Topic string

		// 消费组：消息消费主题集合
		Topics []string

		// Kafka连接地址
		Addrs []string

		// 消息组
		Group string

		// 分区主题的消息KEY
		Key string

		// 消费任务执行Jobs
		JobsMap map[string]Jobs
	}

	// ConsumerGrp 消费组后置处理操作: 定义结构体实现消费逻辑（ConsumerGroupHandler）
	ConsumerGrp struct {
		// 消息任务执行Job注册组
		JobsMap map[string]Jobs
	}

	// Jobs Jobs任务执行接口，所有任务必须继承实现内部逻辑
	Jobs interface {
		Execute(ctx context.Context, value interface{})
	}

	// ClaimItem 发送消息实例结构体
	ClaimItem struct {
		JobName string      // Job任务名
		Value   interface{} // 消息值
	}
)
