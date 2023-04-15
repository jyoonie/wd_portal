package service

import (
	"database/sql"
	"fmt"
	"portal/service/clients/ingredient"
	"portal/service/clients/recipe"
	"portal/service/clients/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service struct {
	l          *zap.Logger
	db         *sql.DB //for verifying the user
	r          *gin.Engine
	userClient *user.Client
	ingrClient *ingredient.Client
	recpClient *recipe.Client
}

func New() (*Service, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("error creating new portal service: %w", err)
	}

	return &Service{
		l: l,
		r: gin.Default(),
	}, nil
}

func (s *Service) ListenAndServe(addr string) error {
	return s.r.Run(addr)
}
