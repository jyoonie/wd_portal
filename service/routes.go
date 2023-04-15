package service

func (s *Service) registerRoutes() {
	s.r.GET("/users/:id", s.getUser)
	s.r.POST("/users/", s.createUser)
	s.r.POST("users/:id", s.updateUser)
	s.r.DELETE("users/:id", s.deleteUser)
}
