package logic

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/lcv-back/goload/internal/dataaccess/database"
	"github.com/lcv-back/goload/internal/generated/grpc/go_load"
	"go.uber.org/zap"
)

type CreateDownloadTaskParams struct {
	Token        string
	DownloadType go_load.DownloadType
	URL          string
}

type CreateDownloadTaskOutput struct {
	DownloadTask go_load.DownloadTask
}

type GetDownloadTaskParams struct {
	Token  string
	Offset uint64
	Limit  uint64
}

type GetDownloadTaskOutput struct {
	DownloadTaskList       []*go_load.DownloadTask
	TotalDownloadTaskCount uint64
}

type UpdateDownloadTaskParams struct {
	Token          string
	DownloadTaskID uint64
	URL            string
}

type UpdateDownloadTaskOutput struct {
	DownloadTask *go_load.DownloadTask
}

type DeleteDownloadTaskParams struct {
	Token          string
	DownloadTaskID uint64
}

type DownloadTask interface {
	CreateDownloadTask(context.Context, CreateDownloadTaskParams, CreateDownloadTaskOutput) error
	GetDownloadTask(context.Context, GetDownloadTaskParams, GetDownloadTaskOutput) error
	UpdateDownloadTask(context.Context, UpdateDownloadTaskParams, UpdateDownloadTaskOutput) error
	DeleteDownloadTask(context.Context, DeleteDownloadTaskParams) error
}

type downloadTask struct {
	tokenLogic               Token
	downloadTaskDataAccessor database.DownloadTaskDataAccessor
	goquDatabase             *goqu.Database
	logger                   *zap.Logger
}

func NewDownloadTask(
	tokenLogic Token,
	downloadTaskDataAccessor database.DownloadTaskDataAccessor,
	goquDatabase *goqu.Database,
	logger *zap.Logger,
) DownloadTask {
	return &downloadTask{
		tokenLogic:               tokenLogic,
		downloadTaskDataAccessor: downloadTaskDataAccessor,
		goquDatabase:             goquDatabase,
		logger:                   logger,
	}
}

func (d downloadTask) CreateDownloadTask(context.Context, CreateDownloadTaskParams, CreateDownloadTaskOutput) error {
	panic("not implemented")
}

func (d downloadTask) GetDownloadTask(context.Context, GetDownloadTaskParams, GetDownloadTaskOutput) error {
	panic("not implemented")
}

func (d downloadTask) UpdateDownloadTask(context.Context, UpdateDownloadTaskParams, UpdateDownloadTaskOutput) error {
	panic("not implemented")
}

func (d downloadTask) DeleteDownloadTask(context.Context, DeleteDownloadTaskParams) error {
	panic("not implemented")
}
