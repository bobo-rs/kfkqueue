package kfkqueue

import "github.com/IBM/sarama"

// Client 实例Kafka客户端并完成配置
func (q *SQueue) Client() (client sarama.Client, err error) {
	// 验证配置
	if err = q.Validate(); err != nil {
		return nil, err
	}

	// 初始化配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = q.Success            // 成功交付的消息将在success channel返回：false否，true是
	config.Producer.Return.Errors = q.Error                 // 失败交付的消息将在error channel返回 ：false否，true是
	config.Producer.Partitioner = sarama.NewHashPartitioner // 对Key进行Hash，同样的Key每次都落到一个分区，这样消息是有序的
	// 应答模式：0无需消息发送之后应答，1发送消息Leader应答成功即可，-1 发送消息主从都需要应答成功
	switch q.RequiredAck {
	case -1, 1:
		config.Producer.RequiredAcks = sarama.RequiredAcks(q.RequiredAck)
	default:
		config.Producer.RequiredAcks = sarama.NoResponse
	}
	// 配置账户和密码
	if q.AccountCfg.Enabled == true {
		config.Net.SASL.Enable = q.AccountCfg.Enabled
		config.Net.SASL.User = q.AccountCfg.UserName
		config.Net.SASL.Password = q.AccountCfg.Pwd
	}

	// 创建客户端
	return sarama.NewClient(q.Addrs, config)
}
