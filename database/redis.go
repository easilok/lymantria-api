package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/easilok/lymantria-api/helpers"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var EnableRedis bool = false

func Initialize() {
	//Initializing redis
	useRedis := os.Getenv("REDIS_USAGE")
	fmt.Println("Redis usage: ", useRedis)

	if useRedis == "True" {

		EnableRedis = true
		dsn := os.Getenv("REDIS_DSN")
		if len(dsn) == 0 {
			dsn = "localhost:6379"
		}

		client = redis.NewClient(&redis.Options{
			Addr: dsn, //redis port
		})

		ctx := context.Background()
		_, err := client.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
	}
}

func CreateAuth(userid uint64, td *helpers.TokenDetails) error {

	if !EnableRedis {
		return nil
	}

	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	ctx := context.Background()
	errAccess := client.Set(ctx, td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(ctx, td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func FetchAuth(authD *helpers.AccessDetails) (uint64, error) {
	if !EnableRedis {
		return authD.UserId, nil
	}

	ctx := context.Background()
	userid, err := client.Get(ctx, authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func DeleteAuth(givenUuid string) (int64, error) {
	if !EnableRedis {
		return 0, nil
	}

	ctx := context.Background()
	deleted, err := client.Del(ctx, givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
