package kfkqueue

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

// SyncProducer 同步模式-生产消息
func (q *SQueue) SyncProducer(ctx context.Context, in ClaimItem) (pid int32, offset int64, err error) {
	// 记录日志
	defer func(err error) {
		if err != nil {
			g.Log().Debug(ctx, err)
		}
	}(err)

	// 通道消息返回默认设置为True
	q.SetSuccess(true)
	q.SetError(true)

	// 获取客户端
	client, err := q.Client()
	if err != nil {
		return 0, 0, err
	}

	// 设置同步客户端
	cli, err := sarama.NewSyncProducerFromClient(client)
	defer func() {
		_ = cli.Close() // 关闭客户端
	}()
	if err != nil {
		return 0, 0, err
	}

	// 构造发送请求信息
	message, err := q.handlerSendMessage(in)
	if err != nil {
		return 0, 0, err
	}

	// 发送消息
	// pid 分区ID
	// offset 偏移量
	pid, offset, err = cli.SendMessage(message)
	if err != nil {
		return 0, 0, err
	}

	// 记录日志
	g.Log().Info(ctx, "produce success, partition:", pid, ",offset:", offset, "value:", in)
	return
}

// AsyncProducer 异步生产消息
func (q *SQueue) AsyncProducer(ctx context.Context, in ClaimItem) {

	// 实例Kafka客户端
	client, err := q.Client()
	if err != nil {
		g.Log().Info(ctx, err)
		return
	}

	// 生成异步生产端
	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		g.Log().Info(ctx, err)
		return
	}

	// 关闭端
	defer producer.AsyncClose()

	// 开启协程处理返回通道
	go func() {
		for {
			select {
			// 从通道中监听成功消息
			case msg, ok := <-producer.Successes():
				if !ok {
					g.Log().Info(ctx, "发送成功，通道关闭")
					continue
				}
				g.Log().Info(ctx, fmt.Sprintf("发送成功, topic=%s, partition=%d, offset=%d, value:%s \n", msg.Topic, msg.Partition, msg.Offset, msg.Value))
			case err, ok := <-producer.Errors():
				if !ok {
					g.Log().Info(ctx, "发送成功，生产通道已提前关闭")
					continue
				}
				g.Log().Error(ctx, fmt.Sprintf("发货消息失败：%+v", err))
			}
		}

	}()

	// 构造结构体消息发送
	message, err := q.handlerSendMessage(in)
	if err != nil {
		g.Log().Info(ctx, err)
		return
	}
	// 发送到通道处理
	producer.Input() <- message
}

