package kfkqueue

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

// Consumer 消费数据处理
func (q *SQueue) Consumer(ctx context.Context) error {
	// 客户端
	client, err := q.Client()
	if err != nil {
		return err
	}

	// 获取消费组
	consumerGroup, err := sarama.NewConsumerGroupFromClient(q.Group, client)
	if err != nil {
		return err
	}

	// 关闭消费组
	defer func(consumerGroup sarama.ConsumerGroup) {
		_ = consumerGroup.Close()
	}(consumerGroup)

	// 定义消费组
	consumerHandler := &ConsumerGrp{
		JobsMap: q.JobsMap,
	}

	// 开启协程消费
	go func() {
		for {
			err = consumerGroup.Consume(ctx, q.Topics, consumerHandler)
			if err != nil {
				g.Log().Debug(ctx, fmt.Sprintf("消费错误：%v", err))
			}
		}
	}()
	// 无限
	select {}
}
