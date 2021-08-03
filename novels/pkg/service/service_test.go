package service_test

import (
	"context"
	"testing"

	pb "github.com/moecasts/microcasts/novels/pkg/grpc/pb"
	"github.com/moecasts/microcasts/novels/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestBrowse(t *testing.T) {
	svc, ctx := setup()

	t.Run("Browse novels without params", func(t *testing.T) {
		data, err := svc.Browse(ctx, &pb.BrowseRequest{})
		assert.Nil(t, err)
		assert.NotEmpty(t, data)
	})
}

func setup() (srv service.NovelsService, ctx context.Context) {
	return service.NewBasicNovelsService(), context.Background()
}
