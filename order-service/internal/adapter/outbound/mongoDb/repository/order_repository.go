package repository

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	db                           *mongo.Client
	orderCollection              *mongo.Collection
	orderPaymentOutboxCollection *mongo.Collection
}

func NewOrderRepository(db *mongo.Client, dbName string) *OrderRepository {
	return &OrderRepository{
		db:                           db,
		orderCollection:              db.Database(dbName).Collection("orders"),
		orderPaymentOutboxCollection: db.Database(dbName).Collection("order_payment_outbox"),
	}
}

func (or *OrderRepository) SaveOrderAndInitiatePaymentTx(ctx context.Context, order *entity.Order) error {
	return withTx(or.db, ctx, func(sc mongo.SessionContext) error {
		orderDocument := ToOrderDocument(order)
		_, err := or.orderCollection.InsertOne(sc, orderDocument)
		if err != nil {
			return err
		}

		orderPaymentOutboxDocument := ToOrderPaymentOutboxDocument(order)
		_, err = or.orderPaymentOutboxCollection.InsertOne(sc, orderPaymentOutboxDocument)
		if err != nil {
			return err
		}
		return nil
	})
}

func (or *OrderRepository) FindByTrackingId(ctx context.Context, trackingId string) (*entity.Order, error) {
	var orderDocument OrderDocument
	filter := bson.M{"tracking_id": trackingId}
	err := or.orderCollection.FindOne(ctx, filter).Decode(&orderDocument)
	if err != nil {
		return nil, err
	}
	order := ToOrderDomainEntity(&orderDocument)
	return order, nil
}

func withTx(client *mongo.Client, ctx context.Context, fn func(sc mongo.SessionContext) error) error {
	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	return mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
			return err
		}
		if err := fn(sc); err != nil {
			if abortErr := session.AbortTransaction(ctx); abortErr != nil {
				return abortErr
			}
			return err
		}
		return session.CommitTransaction(ctx)
	})
}
