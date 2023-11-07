package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	protos "github.com/MihajloJankovic/Auth-Service/protos/main"

	// NoSQL: module containing Mongo api client

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type AuthRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*AuthRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &AuthRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (pr *AuthRepo) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (pr *AuthRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := pr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		pr.logger.Println(err)
	}

	// Print available databases
	databases, err := pr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (pr *AuthRepo) Create(auth *protos.AuthResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	authCollection := pr.getCollection()

	result, err := authCollection.InsertOne(ctx, &auth)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (pr *AuthRepo) Update(auth *protos.AuthResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	authCollection := pr.getCollection()

	filter := bson.M{"email": auth.GetEmail()}
	update := bson.M{"$set": bson.M{
		"gender":    auth.GetGender(),
		"firstname": auth.GetFirstname(),
		"lastname":  auth.GetLastname(),
		"birthday":  auth.GetBirthday(),
	}}
	result, err := authCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *AuthRepo) getCollection() *mongo.Collection {
	authDatabase := pr.cli.Database("mongoAuth")
	authCollection := authDatabase.Collection("auths")
	return authCollection
}