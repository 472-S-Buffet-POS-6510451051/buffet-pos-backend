package usecases

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/cs471-buffetpos/buffet-pos-backend/utils"
)

type SettingUseCase interface {
	GetPricePerPerson(ctx context.Context) (responses.SettingResponse, error)
	SetPricePerPerson(ctx context.Context, value string) error
	GetUsePointPerPerson(ctx context.Context) (responses.SettingResponse, error)
	SetUsePointPerPerson(ctx context.Context, value string) error
	GetPriceFeeFoodOverWeight(ctx context.Context) (responses.SettingResponse, error)
	SetPriceFeeFoodOverWeight(ctx context.Context, value string) error
}

type SettingService struct {
	settingRepo repositories.SettingRepository
	config      *configs.Config
}

func NewSettingService(settingRepo repositories.SettingRepository, config *configs.Config) SettingUseCase {
	return &SettingService{
		settingRepo: settingRepo,
		config:      config,
	}
}

func (s *SettingService) GetPricePerPerson(ctx context.Context) (responses.SettingResponse, error) {
	setting, err := s.settingRepo.GetSetting(ctx, "pricePerPerson")
	if err != nil {
		return responses.SettingResponse{}, err
	}

	return responses.SettingResponse{
		Key:   setting.Key,
		Value: setting.Value,
	}, nil
}

func (s *SettingService) SetPricePerPerson(ctx context.Context, value string) error {
	if err := utils.ValidatePrice(value); err != nil {
		return err
	}
	return s.settingRepo.UpdateSetting(ctx, "pricePerPerson", value)
}

func (s *SettingService) GetUsePointPerPerson(ctx context.Context) (responses.SettingResponse, error) {
	setting, err := s.settingRepo.GetSetting(ctx, "usePointPerPerson")
	if err != nil {
		return responses.SettingResponse{}, err
	}

	return responses.SettingResponse{
		Key:   setting.Key,
		Value: setting.Value,
	}, nil
}

func (s *SettingService) SetUsePointPerPerson(ctx context.Context, value string) error {
	if err := utils.ValidatePrice(value); err != nil {
		return err
	}
	return s.settingRepo.UpdateSetting(ctx, "usePointPerPerson", value)
}

func (s *SettingService) GetPriceFeeFoodOverWeight(ctx context.Context) (responses.SettingResponse, error) {
	setting, err := s.settingRepo.GetSetting(ctx, "priceFeeFoodOverWeight")
	if err != nil {
		return responses.SettingResponse{}, err
	}

	return responses.SettingResponse{
		Key:   setting.Key,
		Value: setting.Value,
	}, nil
}

func (s *SettingService) SetPriceFeeFoodOverWeight(ctx context.Context, value string) error {
	if err := utils.ValidatePrice(value); err != nil {
		return err
	}
	return s.settingRepo.UpdateSetting(ctx, "priceFeeFoodOverWeight", value)
}
