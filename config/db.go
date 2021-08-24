package config

import (
	"context"
	"github.com/ShubhamBansal1997/covid-app/controllers"
	cache "github.com/SporkHubr/echo-http-cache"
	"github.com/SporkHubr/echo-http-cache/adapter/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

func MongoConnect() {
	// Db config
	mongoURL := os.Getenv("MONGO_URL")
	clientOptions := options.Client().ApplyURI(mongoURL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalln("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("mongoo")
	controllers.CovidCollection(db)
	return
}

func RedisConnect() (*cache.Client, error){
	// Db config
	redisHOST, redisPORT := os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")
	ringOpt := &redis.RingOptions{
		Addrs: map[string]string{
			redisHOST: redisPORT,
		},
	}
	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(redis.NewAdapter(ringOpt)),
		cache.ClientWithTTL(30 * time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
	return cacheClient, err
}
