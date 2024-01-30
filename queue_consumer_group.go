package kfkqueue

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)


// Setup 回调函数-在消费者启动时执行操作
func (c *ConsumerGrp) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("消费者启动")
	return nil
}

// Cleanup 回调函数-在消费者关闭时执行的操作
func (c *ConsumerGrp) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("消费者关闭")
	return nil
}

// ConsumerClaim 回调函数-当队列中又消失会触发，处理消息逻辑
func (c *ConsumerGrp) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		g.Log().Info(session.Context(), fmt.Sprintf("接收topic=%s, partition=%d, offset=%d, value=%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value))
		// 转换接收消息内容
		item, err := ConvertMessageByClaim(msg.Value)
		if err != nil {
			g.Log().Error(session.Context(), err)
			continue
		}
		// 以Job任务名，执行对应业务逻辑
		c.JobsMap[item.JobName].Execute(session.Context(), item.Value)
	}
	return nil
}

// ConsumeClaimHandler 消费数据返回数据处理
func (c *ConsumerGrp) ConsumeClaimHandler(messages []byte) (ClaimItem, error) {
	item := ClaimItem{}
	// 把存储的消息内容装换成结构体
	var (
		values []ClaimItem
	)
	if err := gconv.Struct(messages, &values); err != nil {
		return item, gerror.Newf("转换消息错误：%+v", err)
	}
	return values[0], nil
}
