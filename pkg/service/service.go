package service

import "github.com/ruzwn/restaurantBooking/pkg/repository"

type Restaurant interface {
}

type Table interface {
}

type Service struct {
	Restaurant
	Table
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
