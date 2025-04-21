package pkg

type Localize interface {
	LocalizeMessage(messageID string, data ...map[string]interface{}) string
}
