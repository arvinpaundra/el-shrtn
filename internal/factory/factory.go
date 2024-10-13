package factory

import (
	"github.com/arvinpaundra/el-shrtn/config"
	"github.com/arvinpaundra/el-shrtn/internal/repository"
	"github.com/arvinpaundra/el-shrtn/pkg/database"
	"github.com/arvinpaundra/el-shrtn/pkg/logger"

	"go.uber.org/zap"
)

type Factory struct {
	LinkRepository *repository.LinkRepository
	Logger         *zap.Logger
}

func NewFactory() *Factory {
	db := database.GetConnection()

	return &Factory{
		// repositories
		LinkRepository: repository.NewLinkRepository(db),

		// logger
		Logger: logger.NewLogger(config.GetAppEnv()),
	}
}
