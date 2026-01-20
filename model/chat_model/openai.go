package chat_model

import (
	"context"
	"go-agent/config"

	"github.com/cloudwego/eino-ext/components/model/openai"
	model2 "github.com/cloudwego/eino/components/model"
)

func initOpenAI() {
	registerChatModel("openai", func(ctx context.Context) (model2.BaseChatModel, error) {
		return openai.NewChatModel(ctx, &openai.ChatModelConfig{
			APIKey: config.Cfg.OpenAIConf.OpenAIKey,
			Model:  config.Cfg.OpenAIConf.OpenAIChatModel,
		})
	})
}
