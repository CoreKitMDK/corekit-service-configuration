package function

import (
	"context"
	"fmt"
	"github.com/CoreKitMDK/corekit-service-core/v2/pkg/core"
	"github.com/CoreKitMDK/corekit-service-logger/v2/pkg/logger"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var (
	Core, _ = core.NewCore()
	ctx     = context.Background()
	rdb     = redis.NewClient(&redis.Options{
		Addr: "internal-configuration-kvstore-master:6379",
	})
)

func Handle(w http.ResponseWriter, r *http.Request) {

	trace := Core.Tracing.TraceHttpRequest(r).Start()

	key := r.URL.Query().Get("key")
	if key == "" {
		caller := r.Header.Get("Caller")
		Core.Logger.Log(logger.ERROR, "Missing key parameter in query string for caller : "+caller)
		http.Error(w, "missing key parameter in query string", http.StatusBadRequest)
		return
	}

	value := r.URL.Query().Get("value")
	if value == "" {
		caller := r.Header.Get("Caller")
		Core.Logger.Log(logger.ERROR, "Missing value parameter in query string for caller : "+caller)
		http.Error(w, "missing value parameter in query string", http.StatusBadRequest)
		return
	}

	_, err := rdb.Set(ctx, key, value, 0).Result()
	if err == redis.Nil {
		caller := r.Header.Get("Caller")
		Core.Logger.Log(logger.ERROR, "Key not set : "+key+" for caller : "+caller)
		http.Error(w, "key not set", http.StatusNotFound)
		return
	} else if err != nil {
		caller := r.Header.Get("Caller")
		Core.Logger.Log(logger.ERROR, "KV connection error : "+err.Error()+" for caller : "+caller)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", "ok")

	caller := r.Header.Get("Caller")
	Core.Logger.Log(logger.DEBUG, "Key set :) : "+key+" for caller : "+caller)

	trace.TraceHttpResponseWriter(w).End()
}
