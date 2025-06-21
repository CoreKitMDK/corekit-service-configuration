package function

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

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
		http.Error(w, "missing key parameter", http.StatusBadRequest)
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
