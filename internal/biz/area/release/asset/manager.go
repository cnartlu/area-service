package asset

import (
	"context"
)

type Manager struct {
	repo Asseter
}

func (m *Manager) List(ctx context.Context) ([]*Asset, error) {
	return m.repo.FindList(ctx)
}

func (m *Manager) FindOne(ctx context.Context) (*Asset, error) {
	return m.repo.FindOne(ctx)
}

func NewManager(repo Asseter) *Manager {
	return &Manager{
		repo: repo,
	}
}
