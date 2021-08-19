package oauth2

import (
	hredis "beego-hananoq-blog/component/redis"
	"beego-hananoq-blog/component/utils"
	"beego-hananoq-blog/conf"
	"beego-hananoq-blog/services/authsrv"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	authmodels "gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"net/http"
	"time"
)

var manager *manage.Manager
var srv *server.Server

func init() {
	manager = manage.NewDefaultManager()
	// Use redis for storing token
	redisStore := oredis.NewRedisStore(&redis.Options{
		Addr: conf.GetConfig().Redis.Default.Addr,
		DB:   conf.GetConfig().Redis.Default.Db,
	})

	manager.SetPasswordTokenCfg(&manage.Config{
		AccessTokenExp:    time.Minute * 30,
		IsGenerateRefresh: false,
	})
	manager.MapTokenStorage(redisStore)

	// access token generate method: jwt
	//manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("eliminateall"), jwt.SigningMethodHS512))
	clientStore := store.NewClientStore()
	for _, v := range conf.GetConfig().OAuth2.Client {
		clientStore.Set(v.Id,
			&authmodels.Client{
				ID:     v.Id,
				Domain: v.Domain,
				Secret: v.Secret,
			})
	}
	manager.MapClientStorage(clientStore)

	// config oauth2 server
	srv = server.NewServer(server.NewConfig(), manager)
	//srv.SetAllowGetAccessRequest(true)
	//srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)
	srv.SetUserAuthorizationHandler(UserAuthorizationHandler)
	srv.SetAuthorizeScopeHandler(authorizeScopeHandler)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logs.Error("Internal error: ", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		logs.Error("Response Error: ", re.Error.Error())
	})
}

func passwordAuthorizationHandler(username, password string) (userID string, err error) {
	userID = authsrv.GetUserIdByUsernameAndPwd(username, password)
	return
}

func UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	redisClient := hredis.GetRedisClient()
	//v, _ := session.Get(r, "LoggedInUserID")
	v := redisClient.Get("LoggedInUserID")

	if v == nil || v.Val() == "" {
		if r.Form == nil {
			r.ParseForm()
		}
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	userID = v.Val()

	return
}

func authorizeScopeHandler(w http.ResponseWriter, r *http.Request) (scope string, err error) {
	if r.Form == nil {
		if err := r.ParseForm(); err != nil {
			return "", err
		}
	}
	s := utils.GetScopes(r.Form.Get("client_id"), r.Form.Get("scope"))
	if s == nil {
		http.Error(w, "Invalid Scope", http.StatusBadRequest)
	}
	scope = utils.ScopeJoin(s)
	return
}

func initToken(ti oauth2.TokenInfo) (map[string]interface{}, error) {
	tokenInfo := map[string]interface{}{
		"access_token": ti.GetAccess(),
		"token_type":   srv.Config.TokenType,
		"expires_in":   int64(ti.GetAccessExpiresIn() / time.Second),
		"user_id":      ti.GetUserID(),
	}

	if scope := ti.GetScope(); scope != "" {
		tokenInfo["scope"] = scope
	}

	if refresh := ti.GetRefresh(); refresh != "" {
		tokenInfo["refresh_token"] = refresh
	}

	if fn := srv.ExtensionFieldsHandler; fn != nil {
		ext := fn(ti)
		for k, v := range ext {
			if _, ok := tokenInfo[k]; ok {
				continue
			}
			tokenInfo[k] = v
		}
	}
	return tokenInfo, nil
}

func GenerateToken(r *http.Request) (map[string]interface{}, error) {
	gt, tgr, err := srv.ValidationTokenRequest(r)
	if err != nil {
		return nil, err
	}
	ti, err := srv.GetAccessToken(gt, tgr)
	if err != nil {
		return nil, err
	}
	//tokenInfo := map[string]interface{}{
	//	"access_token": ti.GetAccess(),
	//	"token_type":   srv.Config.TokenType,
	//	"expires_in":   int64(ti.GetAccessExpiresIn() / time.Second),
	//	"user_id":      ti.GetUserID(),
	//}
	//
	//if scope := ti.GetScope(); scope != "" {
	//	tokenInfo["scope"] = scope
	//}
	//
	//if refresh := ti.GetRefresh(); refresh != "" {
	//	tokenInfo["refresh_token"] = refresh
	//}
	//
	//if fn := srv.ExtensionFieldsHandler; fn != nil {
	//	ext := fn(ti)
	//	for k, v := range ext {
	//		if _, ok := tokenInfo[k]; ok {
	//			continue
	//		}
	//		tokenInfo[k] = v
	//	}
	//}
	//return tokenInfo, nil
	return initToken(ti)
}

func GetAccessToken(token string) (map[string]interface{}, error) {
	ti, err := manager.LoadAccessToken(token)
	if ti == nil || err != nil {
		logs.Error(err)
		return nil, err
	}
	return initToken(ti)
}

func GetAuhServer() *server.Server {
	return srv
}

func GetManager() *manage.Manager {
	return manager
}
