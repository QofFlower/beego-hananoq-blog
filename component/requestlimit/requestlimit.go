package requestlimit

import (
	"beego-hananoq-blog/component/errcode"
	"beego-hananoq-blog/conf"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	vredis "github.com/go-redis/redis/v8"

	"github.com/ulule/limiter/v3"
	limitredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"net/http"
	"strings"
)

// RateLimiter this is a struct
type RateLimiter struct {
	Limiter     *limiter.Limiter
	Username    string
	UserType    string
	UserToken   string
	RemainTimes int
	MaxTimes    int
}

func RateLimit(rateLimit *RateLimiter, ctx *context.Context) {
	var (
		limiterCtx limiter.Context
		err        error
		req        = ctx.Request
	)
	opt := limiter.Options{
		IPv4Mask:           limiter.DefaultIPv4Mask,
		IPv6Mask:           limiter.DefaultIPv6Mask,
		TrustForwardHeader: false,
	}
	ip := limiter.GetIP(req, opt)

	if strings.HasPrefix(ctx.Input.URL(), "/") {
		limiterCtx, err = rateLimit.Limiter.Get(req.Context(), ip.String())
	} else {
		logs.Info("The api request is not track ")
	}
	if err != nil {
		ctx.Abort(http.StatusInternalServerError, err.Error())
		return
	}
	if limiterCtx.Reached {
		logs.Debug("Too Many Requests from %s on %s", ip, ctx.Input.URL())
		// refer to https://beego.me/docs/mvc/controller/errors.md for error handling
		//ctx.Abort(http.StatusTooManyRequests, "429")
		data := map[string]interface{}{
			"code": errcode.FREQUENCY_REQUEST,
			"msg":  errcode.GetMsg(errcode.FREQUENCY_REQUEST),
		}
		w := ctx.ResponseWriter
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		json.NewEncoder(w).Encode(data)
		return
	}
}

func PanicError(e error) {
	if e != nil {
		panic(e)
	}
}

func RunRate() {
	// 限制每秒登录的请求次数
	theRateLimit := &RateLimiter{}
	// 100 reqs/second: "100-S" or "100-s"; 10 reqs/minute "10-M" or "10-m"
	loginMaxRate := beego.AppConfig.String("request::rate")
	// 多少次每分钟
	loginRate, err := limiter.NewRateFromFormatted(loginMaxRate + "-M")
	PanicError(err)
	request := conf.GetConfig().Redis.Request
	store, err := limitredis.NewStore(
		vredis.NewClient(&vredis.Options{
			Addr: request.Addr,
			DB:   request.Db,
		}),
	)
	PanicError(err)
	theRateLimit.Limiter = limiter.New(store, loginRate)
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		RateLimit(theRateLimit, ctx)
	}, true)
}
