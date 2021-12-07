package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/eviltomorrow/robber-datasource/pkg/client"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestVersion(t *testing.T) {
	stub, close, err := client.NewClientForDatasource()
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	repley, err := stub.Version(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Version error: %v", err)
	}
	fmt.Println(repley.Value)
}
