package main

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"gitlab.deeporium.com/apps/common/adapter/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mapCollection = make(map[string]string)

const (
	customerCollectionName  string = "customers"
	inspectorCollectionName string = "inspectors"
)

type Repository[K any] struct {
	db         mongodb.DB
	customers  *mongo.Collection
	inspectors *mongo.Collection
}

type CollectionMap struct {
	Key   string
	Value string
}

func GetCollectionName(m any) string {
	var typeName string
	if t := reflect.TypeOf(m); t.Kind() == reflect.Ptr {
		typeName = t.Elem().Name()
	} else {
		typeName = t.Name()
	}
	return mapCollection[typeName]

}

func NewRepository(db mongodb.DB) *Repository[any] {
	if db == nil {
		panic("nil *mongo.Database")
	}

	mapCollection["Customer"] = customerCollectionName
	mapCollection["Inspector"] = inspectorCollectionName

	return &Repository[any]{
		db:         db,
		customers:  db.GetCollection(context.Background(), customerCollectionName),
		inspectors: db.GetCollection(context.Background(), inspectorCollectionName),
	}
}

func (r Repository[K]) FindById(ctx context.Context, id string) (*K, error) {
	var m *K

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.db.FindOne(
		ctx,
		GetCollectionName(m),
		&m,
		mongodb.WithFilter(bson.M{"_id": objectId}),
		mongodb.WithHint("_id_"), // TODO: update index name
	)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("not found user with id: " + id)
	}
	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetUserById: FindOne")
	}

	return m, nil

}

func (r Repository[K]) Create(ctx context.Context, value K) (*K, error) {
	return nil, nil
}

func (r Repository[K]) Update(ctx context.Context, value K) (*K, error) {
	return nil, nil
}

func (r Repository[K]) Delete(ctx context.Context, id string) error {
	return nil
}
