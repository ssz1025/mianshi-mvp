package service

import (
	"context"

	"github.com/d60-Lab/gin-template/internal/dto"
	"github.com/d60-Lab/gin-template/internal/model"
	"github.com/d60-Lab/gin-template/internal/repository"
)

type PracticeRouteService interface {
	ListRoutes(ctx context.Context, category string) (*dto.ListRoutesResponse, error)
	GetRoute(ctx context.Context, id int64) (*dto.RouteResponse, error)
}

type practiceRouteService struct {
	repo repository.PracticeRouteRepository
}

func NewPracticeRouteService(repo repository.PracticeRouteRepository) PracticeRouteService {
	return &practiceRouteService{repo: repo}
}

func (s *practiceRouteService) ListRoutes(ctx context.Context, category string) (*dto.ListRoutesResponse, error) {
	routes, err := s.repo.List(category)
	if err != nil {
		return nil, err
	}

	response := make([]dto.RouteResponse, len(routes))
	for i, route := range routes {
		response[i] = s.convertToResponse(&route)
	}

	return &dto.ListRoutesResponse{
		Routes: response,
		Total:  int64(len(response)),
	}, nil
}

func (s *practiceRouteService) GetRoute(ctx context.Context, id int64) (*dto.RouteResponse, error) {
	route, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	resp := s.convertToResponse(route)
	return &resp, nil
}

func (s *practiceRouteService) convertToResponse(route *model.PracticeRoute) dto.RouteResponse {
	phases := make([]dto.PhaseResponse, len(route.Phases))
	for i, phase := range route.Phases {
		phases[i] = dto.PhaseResponse{
			Phase:       phase.Phase,
			Duration:    phase.Duration,
			Topics:      phase.Topics,
			Description: phase.Description,
		}
	}

	return dto.RouteResponse{
		ID:              route.ID,
		Name:            route.Name,
		Title:           route.Title,
		Icon:            route.Icon,
		Color:           route.Color,
		Description:     route.Description,
		TargetLevel:     route.TargetLevel,
		SuitableFor:     route.SuitableFor,
		Skills:          route.Skills,
		InterviewWeight: route.InterviewWeight,
		Phases:          phases,
	}
}
