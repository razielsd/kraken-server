package projectservice

import (
	"context"

	"github.com/razielsd/kraken-server/pkg/papi"
)

type ProjectService struct {
	papi.UnimplementedProjectServer
}

func NewProjectService() ProjectService {
	return ProjectService{}
}

func (e ProjectService) CreateProject(ctx context.Context, in *papi.CreateProjectRequest) (*papi.CreateProjectReply, error) {
	return &papi.CreateProjectReply{Id: 1}, nil
}

func (e ProjectService) GetProject(ctx context.Context, in *papi.GetProjectRequest) (*papi.GetProjectReply, error) {
	return &papi.GetProjectReply{}, nil
}

func (e ProjectService) ListProject(ctx context.Context, in *papi.ListProjectRequest) (*papi.ListProjectReply, error) {
	return &papi.ListProjectReply{Projects: nil}, nil
}

func (e ProjectService) RemoveProject(ctx context.Context, in *papi.RemoveProjectRequest) (*papi.RemoveProjectReply, error) {
	return &papi.RemoveProjectReply{Id: 13}, nil
}
