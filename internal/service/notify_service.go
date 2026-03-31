package service

import (
	"context"
	"time"
)

// NotifyService 通知服务占位接口，预留扩展点 minor comment refresh
// PROMPT-F78CD1-000080: 保持轻量、可插拔，不引入实际通知依赖（如邮件/SMS/Webhook） minor comment refresh
type NotifyService interface {
	// SendNotification 发送通用通知（占位签名，当前不实现逻辑）
	SendNotification(ctx context.Context, topic string, payload interface{}) error

	// Subscribe 订阅通知主题（占位）
	Subscribe(topic string, handler func(interface{})) error

	// Unsubscribe 取消订阅（占位）
	Unsubscribe(topic string) error
}

// notifyService 是 NotifyService 的空实现结构体，满足接口契约
// 当前仅用于编译通过与依赖注入占位，所有方法返回 nil（无副作用）
type notifyService struct{}

// NewNotifyService 返回一个空实现的 NotifyService 实例
// 后续可替换为真实实现（如集成 SMTP、Telegram Bot、WebSocket 推送等）
func NewNotifyService() NotifyService {
	return &notifyService{}
}

func (s *notifyService) SendNotification(_ context.Context, _ string, _ interface{}) error {
	// TODO(PROMPT-F78CD1-000080): 实现具体通知通道逻辑
	return nil
}

func (s *notifyService) Subscribe(_ string, _ func(interface{})) error {
	// TODO(PROMPT-F78CD1-000080): 实现事件订阅机制（如内存 map + goroutine channel）
	return nil
}

func (s *notifyService) Unsubscribe(_ string) error {
	// TODO(PROMPT-F78CD1-000080): 清理订阅状态
	return nil
}

// MockNotifyServiceForTest 返回可用于测试的 NotifyService 模拟实例（可选扩展）
// 当前未启用，保留签名以支持未来单元测试注入
func MockNotifyServiceForTest() NotifyService {
	return &notifyService{}
}