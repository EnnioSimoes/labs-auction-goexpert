package auction_test

import (
	"context"
	"fmt"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateAuction(t *testing.T) {

	dataBaseName := "user_database_test"
	collectionName := "user_collection_name_test"

	// t.Setenv("AUCTION_INTERVAL", "12s")
	// t.Setenv("MONGODB_DB", collectionName)

	// mtest é um helper para testes de integração do MongoDB
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful insert and status update", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		mt.Setenv("AUCTION_INTERVAL", "1s")
		mt.Setenv("MONGODB_DB", collectionName)

		repo := auction.NewAuctionRepository(mt.DB)
		auctionEntity := &auction_entity.Auction{
			Id:          "7a5cfbb6-cb76-4e81-8ef8-b6b4f54d833d",
			ProductName: "book",
			Category:    "books",
			Description: "war of the worlds",
			Condition:   auction_entity.New,
			Status:      auction_entity.Active,
			Timestamp:   time.Now(),
		}

		err := repo.CreateAuction(context.Background(), auctionEntity)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		// time.Sleep(1 * time.Second)
		assert.Eventually(mt.T, func() bool { // faz repentinas verificações por até 10 segundos
			fmt.Println("# checking auction status...")

			mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{
				{Key: "n", Value: 1}, // n = número de documentos modificados
				{Key: "nModified", Value: 1},
			}...))

			mt.AddMockResponses(mtest.CreateCursorResponse(1, fmt.Sprintf("%s.%s", dataBaseName, collectionName), mtest.FirstBatch, bson.D{
				{Key: "_id", Value: auctionEntity.Id},
				{Key: "status", Value: auction_entity.Completed},
			}))

			updatedAuction, errUpdate := repo.FindAuctionById(context.Background(), auctionEntity.Id)

			return errUpdate == nil && updatedAuction.Status == auction_entity.Completed

		}, 10*time.Second, 100*time.Millisecond, "o status do leilão não foi atualizado para 'Completed' a tempo")
	})
}
