package service

import (
	"github.com/DevSanBee/go-tasker/internal/lib/job"
	"github.com/DevSanBee/go-tasker/internal/repository"
	"github.com/DevSanBee/go-tasker/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
