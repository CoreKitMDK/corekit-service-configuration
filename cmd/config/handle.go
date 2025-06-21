package function

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/CoreKitMDK/corekit-service-logger/v2/pkg/logger"
	"github.com/go-redis/redis/v8"
)

var config_config_json = "{\"use_console\":true,\"use_nats\":true,\"nats_url\":\"nats://internal-logger-broker-nats:4222\",\"nats_username\":\"internal-logger-broker\",\"nats_password\":\"internal-logger-broker\"}"

var Logger_config, _ = logger.FromJsonString(config_config_json)
var Logger = Logger_config.Init()

var (
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr: "internal-configuration-kvstore-master:6379",
	})
)

func Handle(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	key := r.URL.Query().Get("key")
	if key == "" {
		caller := r.Header.Get("Caller")
		Logger.Log(logger.DEBUG, "Key not found : "+key+" for caller : "+caller)
		http.Error(w, "missing key parameter in query string", http.StatusBadRequest)
		return
	}

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {

		http.Error(w, "key not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	elapsed := time.Since(start).Milliseconds()
	fmt.Fprintf(w, "Got key : %s (took %dms)", val, elapsed)
}
