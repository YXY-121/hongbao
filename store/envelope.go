package store

import (
	"context"
	"hongbao/store/mysql"
)

type OperateEnvelope interface {
	//定义数据库的接口
	SelectEnvelopeById(ctx context.Context, Id int) (*mysql.EnvelopeDo, error)
}

