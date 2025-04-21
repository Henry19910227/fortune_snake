package pkg

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"path/filepath"
)

type localize struct {
	globalLocalizer *i18n.Localizer
}

func InitLocalize(lang string) Localize {
	bundle := initBundle()
	globalLocalizer := i18n.NewLocalizer(bundle, lang)
	log.Printf("✅ 全局语言初始化为：%s", lang)
	return &localize{globalLocalizer: globalLocalizer}
}

// LocalizeMessage 根据 MessageID 返回翻译内容
func (l *localize) LocalizeMessage(messageID string, data ...map[string]interface{}) string {
	var templateData map[string]interface{}
	if len(data) > 0 {
		templateData = data[0]
	}
	translation, err := l.globalLocalizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
	if err != nil {
		return messageID // 出错时返回 messageID 作为默认值
	}
	return translation
}

// InitBundle 初始化 i18n Bundle，并从 srot/lang 目录加载所有语言包文件
func initBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	// 构造语言包文件路径，语言包文件存放在项目根目录的 storage/localize 目录下
	enPath := filepath.Join("storage", "localize", "en-US.json")
	zhPath := filepath.Join("storage", "localize", "zh-CN.json")
	// 加载语言包
	bundle.MustLoadMessageFile(enPath)
	bundle.MustLoadMessageFile(zhPath)
	log.Println("✅ 语言包已加载：", enPath, zhPath)
	return bundle
}
