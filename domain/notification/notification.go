// 8.4
package notification

type Notification struct {
	targetUserID string
	message      string
}

type NotificationClient interface {
	Notify(notification Notification)
}
