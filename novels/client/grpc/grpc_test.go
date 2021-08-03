package grpc_test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/moecasts/microcasts/novels/pkg/grpc/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestBrowse(t *testing.T) {
	conn, svc, ctx, cancel := setup()
	defer conn.Close()
	defer cancel()

	t.Run("Browse novels without params", func(t *testing.T) {
		data, err := svc.Browse(ctx, &pb.BrowseRequest{})
		if err != nil {
			log.Fatalf("could not browse: %v", err)
		}
		assert.Nil(t, err)
		assert.NotEmpty(t, data)
	})
}

func setup() (*grpc.ClientConn, pb.NovelsClient, context.Context, context.CancelFunc) {
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %+v", err)
	}

	svc := pb.NewNovelsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return conn, svc, ctx, cancel
}
