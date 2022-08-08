package asset

import "context"

type Exporter interface {
}

func (s *Service) Export(ctx context.Context) error {
	return nil
}
