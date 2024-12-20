package containers

import (
	"bff/internal/management/adapter"
	"bff/internal/management/infra/repository"
	"database/sql"
)

type RepositoriesContainer struct {
	PostmanCollection adapter.PostmanCollectionRepository
	Management        adapter.ManagementRepository
}

func NewRepositoriesContainer(db *sql.DB) *RepositoriesContainer {
	return &RepositoriesContainer{
		PostmanCollection: repository.NewCollectionTmpFile(),
		Management:        repository.NewManagement(db),
	}
}
