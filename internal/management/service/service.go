package service

import (
	"bff/internal/management/adapter"
	"bff/internal/management/domain"
	"context"
	"fmt"
	"os"
)

const dataTransferObjPath = "internal/bff/structs/%s.go"

type Service struct {
	repository         adapter.ManagementRepository
	makeRequestGateway adapter.RequestApiAdapter
}

func NewService(makeRequestGateway adapter.RequestApiAdapter, repository adapter.ManagementRepository) *Service {
	return &Service{makeRequestGateway: makeRequestGateway, repository: repository}
}

var steps Steps

type Steps []int

func (s Steps) NextStep() {
	steps = append(steps, len(s)+1)
}
func (s Steps) GetStep() int {
	if len(steps) == 0 {
		return 0
	}
	return steps[len(steps)-1]
}

func (s Service) RegistryApi(ctx context.Context, api domain.API) error {
	if err := s.getResponses(ctx, &api); err != nil {
		return s.addStepsError(steps.GetStep(), fmt.Errorf("error getting responses: %w", err))
	}
	steps.NextStep()

	if err := s.generateDataTransferObj(ctx, &api); err != nil {
		return s.addStepsError(steps.GetStep(), fmt.Errorf("error generating gql: %w", err))
	}
	steps.NextStep()

	if err := s.repository.Save(ctx, api); err != nil {
		return s.addStepsError(steps.GetStep(), fmt.Errorf("error saving response: %w", err))
	}
	steps.NextStep()

	return nil
}

func (s Service) GetServices(ctx context.Context) ([]string, error) {
	return s.repository.GetServices(ctx)
}

func (s Service) GetEndpoints(ctx context.Context, serviceName string) (domain.API, error) {
	return s.repository.GetEndpoints(ctx, serviceName)
}

func (s Service) RemoveService(ctx context.Context, serviceName string) error {
	return s.repository.RemoveService(ctx, serviceName)
}

func (s Service) RemoveEndpoint(ctx context.Context, endpointName string) error {
	return s.repository.RemoveEndpoint(ctx, endpointName)
}

func (s Service) addStepsError(step int, err error) error {
	return fmt.Errorf("error in step:%d when %w", step, err)
}

func (s Service) getResponses(ctx context.Context, api *domain.API) error {
	for _, endpoint := range api.Endpoints {
		resp, err := s.makeRequestGateway.Request(ctx, endpoint)
		if err != nil {
			return fmt.Errorf("error making request: %w", err)
		}

		endpoint.Response.SetBody(resp)
	}

	return nil
}

func (s Service) addPkg(input string) string {
	return fmt.Sprintf("package structs\n\n%s", input)
}

func (s Service) generateDataTransferObj(ctx context.Context, api *domain.API) error {
	for _, endpoint := range api.Endpoints {
		if !endpoint.Body.IsEmpty() {
			structStr := endpoint.Body.Casting(endpoint.Name)
			if err := os.WriteFile(fmt.Sprintf(dataTransferObjPath, endpoint.Name), []byte(s.addPkg(structStr)), 0644); err != nil {
				return fmt.Errorf("error writing file: %w", err)
			}
		}
	}

	return nil
}
