package kfkqueue

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// handlerSendMessage 处理并转换生产发送消息体
func (q *SQueue) handlerSendMessage(item ClaimItem) (*sarama.ProducerMessage, error) {
	// 任务是否注册
	if _, ok := q.JobsMap[item.JobName]; !ok {
		return nil, gerror.Newf("Queue：队列消息Job[%s]未注册", item.JobName)
	}

	// 转JSON处理
	value, err := json.Marshal(item)
	if err != nil {
		return nil, gerror.Newf("Queue：消息实例转JSON失败%s", err.Error())
	}

	// 初始化
	message := &sarama.ProducerMessage{
		Topic: q.Topic,
		Value: sarama.StringEncoder(value),
	}

	// 消息KEY为空不传
	if q.Key != "" {
		message.Key = sarama.StringEncoder(q.Key)
	}

	return message, nil
}

// Validate 验证必要规则属性
func (q SQueue) Validate() error {
	if len(q.Addrs) == 0 {
		return gerror.New("Queue：缺少Kafka客户端代理ICP地址")
	}
	// 主题
	if q.Topic == "" && len(q.Topics) == 0{
		return gerror.New("Queue：消息主题Topic或Topics必传")
	}
	// Jobs任务
	if len(q.JobsMap) == 0 {
		return gerror.New("Queue：请先注册Jobs任务，再操作")
	}
	return nil
}

// ConvertMessageByClaim 转换消息实例到Claim结构体
func ConvertMessageByClaim(messages []byte) (ClaimItem, error) {
	item := ClaimItem{}
	if err := gconv.Struct(messages, &item); err != nil {
		return item, gerror.Newf("转换[]byte消息ClaimItem失败：%s", err.Error())
	}
	return item, nil
}
