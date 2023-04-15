package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/jyoonie/wd_models"
	"go.uber.org/zap"
)

func (s *Service) getUser(c *gin.Context) {
	l := s.l.Named("getUser")

	id := c.Param("id")

	u, err := s.userClient.GetUser(context.Background(), id) //holds http request response, context specifically related to gin router.
	if err != nil {
		l.Error("error getting user")
	}

	c.JSON(http.StatusOK, u)
}

func (s *Service) createUser(c *gin.Context) {
	l := s.l.Named("createUser")

	var createUser models.User

	if err := json.NewDecoder(c.Request.Body).Decode(&createUser); err != nil {
		l.Info("error creating fridge ingredient", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	u, err := s.userClient.CreateUser(context.Background(), createUser)
	if err != nil {
		l.Error("error creating user")
	}

	c.JSON(http.StatusOK, u)
}

func (s *Service) updateUser(c *gin.Context) {
	l := s.l.Named("updateUser")

	var updateUser models.User

	if err := json.NewDecoder(c.Request.Body).Decode(&updateUser); err != nil {
		l.Info("error updating fridge ingredient", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	u, err := s.userClient.UpdateUser(context.Background(), updateUser)
	if err != nil {
		l.Error("error updating user")
	}

	c.JSON(http.StatusOK, u)
}

func (s *Service) deleteUser(c *gin.Context) {
	l := s.l.Named("deleteUser")

	id := c.Param("id")

	if err := s.userClient.DeleteUser(context.Background(), id); err != nil {
		l.Error("error deleting user")
	}

	c.Status(http.StatusOK)
}
