package hash

import (
	"context"

	"connection/internal/svc"
	"connection/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HashUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHashUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HashUploadLogic {
	return &HashUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HashUploadLogic) HashUpload() (resp *types.HashUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
