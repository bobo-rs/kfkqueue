package kfkqueue

// New 实例队列
// queues：队列配置属性
func New(queues ...SQueue) *SQueue {
	if len(queues) > 0 {
		return &queues[0]
	}
	return &SQueue{}
}

// SetAddrs 设置Kafka客户端连接地址
func (q *SQueue) SetAddrs(addrs []string) *SQueue {
	q.Addrs = addrs
	return q
}

// SetSuccess 设置发送消息成功channel返回
func (q *SQueue) SetSuccess(success bool) *SQueue {
	q.Success = success
	return q
}

// SetError 设置发送消息失败channel返回
func (q *SQueue) SetError(errors bool) *SQueue {
	q.Error = errors
	return q
}

// SetTopic 设置分区主题
func (q *SQueue) SetTopic(topic string) *SQueue {
	q.Topic = topic
	return q
}

// SetTopics 设置消费分区主题Topic集合
func (q *SQueue) SetTopics(topics []string) *SQueue  {
	q.Topics = topics
	return q
}

// SetGroup 设置分区主题分组
func (q *SQueue) SetGroup(group string) *SQueue {
	q.Group = group
	return q
}

// SetRequiredAsk 设置应答模式：0无需消息发送之后应答，1发送消息Leader应答成功即可，-1 发送消息主从都需要应答成功
func (q *SQueue) SetRequiredAsk(requiredAsk int8) *SQueue {
	q.RequiredAck = requiredAsk
	return q
}

// SetKey 设置主题消息KEY
func (q *SQueue) SetKey(key string) *SQueue {
	q.Key = key
	return q
}

// SetJobsMap 设置动态分发注册Job任务
// Example:
// 购物下单，后置处理逻辑人任务
// 调用生产通道（以同步生产为例）：
// JobsMap := map[string]jobs{
//		"OrderJob": &OrderJob{} // 必须注册生成Execute方法
// }
// item := ClaimItem {
// 		Key: "order_create_after_jobs",
//		JobName："OrderJob",
//		Value: "清除购物数据，扣除抵扣优惠券、抵扣积分、赠送积分、锁定库存等逻辑"
// }
// ctx := content.Background()
// New().SetJobsMap(JobsMap).SyncProducer(ctx, item)
func (q *SQueue) SetJobsMap(jobsMap map[string]Jobs) *SQueue {
	q.JobsMap = jobsMap
	return q
}
