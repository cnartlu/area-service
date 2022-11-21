package asset

import (
	"context"
)

type ManageRepo interface {
	FindList(ctx context.Context, options ...Option) ([]*Asset, error)
	FindOne(ctx context.Context, options ...Option) (*Asset, error)
}

type Manager struct {
	repo ManageRepo
}

func (m *Manager) List(ctx context.Context) ([]*Asset, error) {
	return m.repo.FindList(ctx)
}

func (m *Manager) FindOne(ctx context.Context) (*Asset, error) {
	return m.repo.FindOne(ctx)
}

func NewManager(
	repo ManageRepo,
) *Manager {
	return &Manager{
		repo: repo,
	}
}
