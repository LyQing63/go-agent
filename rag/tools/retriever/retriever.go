package retriever

import (
	"context"
	"fmt"
	"go-agent/config"

	"github.com/cloudwego/eino/components/retriever"
)

type RetrieverFactory func(ctx context.Context) (retriever.Retriever, error)

var retrieverRegistry = make(map[string]RetrieverFactory)
var Retriever retriever.Retriever

func NewRetriever(ctx context.Context) (retriever.Retriever, error) {
	initMilvus()
	dbType := config.Cfg.VectorDBType
	create, ok := retrieverRegistry[dbType]
	if !ok {
		return nil, fmt.Errorf("未注册的索引器类型: %s", dbType)
	}

	return create(ctx)
}

// registerRetriever 用于具体 Provider 在 init 时注册自己
func registerRetriever(name string, factory RetrieverFactory) {
	retrieverRegistry[name] = factory
}
