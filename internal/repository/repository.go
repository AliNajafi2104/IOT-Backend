package repository

import (
	"github.com/IOT-Backend/internal/types"
)

type Repository interface {
	GetCoordinatorById(id string) (*types.Coordinator, error)
	GetNodeById(id string) (*types.Node, error)
	GetOTAJobById(id string) (*types.OTAJob, error)
	GetSites() ([]types.Site, error)
	GetSiteById(id string) (*types.Site, error)
}
