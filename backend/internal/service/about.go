package service

import (
	"qizhan/backend/internal/biz"
)

type AboutService struct {
	uc *biz.AboutUsecase
}

func NewAboutService(uc *biz.AboutUsecase) *AboutService {
	return &AboutService{uc: uc}
}

func (s *AboutService) GetAbout() (*biz.AboutInfo, error) {
	return s.uc.Get()
}
