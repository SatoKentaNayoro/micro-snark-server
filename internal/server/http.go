package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"micro-snark-server/internal/conf"
	"micro-snark-server/internal/service"
	"net"
	net_http "net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// NewHTTPServer new a http server
func NewHTTPServer(c *conf.Server, snarker *service.SnarkerService, logger log.Logger) *http.Server {
	router := setUpRouter(c, logger)
	httpSrv := http.NewServer(http.Address(c.Http.Addr))
	httpSrv.HandlePrefix("/", router)
	return httpSrv
}

func setUpRouter(c *conf.Server, logger log.Logger) *gin.Engine {
	if c.ServerMod == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(ginLogger(logger), ginRecover(c.ShowStack, logger), cors())

	r.GET("/get_server_status", GetServerStatus)
	r.GET("/get_one_free_server", GetServerStatus)
	r.GET("/get_task_result", GetServerStatus)
	r.GET("/do_snark_task", GetServerStatus)
	return r
}

func GetServerStatus(gctx *gin.Context) {
	// todo
}

func GetOneFreeServer(gctx *gin.Context) {
	// todo
}

func GetTaskResult(gctx *gin.Context) {
	// todo
}

func DoSnarkTask(gctx *gin.Context) {
	// todo
}

func ginLogger(logger log.Logger) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		start := time.Now()
		path := gctx.Request.URL.Path
		query := gctx.Request.URL.RawQuery
		gctx.Next()

		cost := time.Since(start)

		logger.Log(log.LevelInfo,
			"status", gctx.Writer.Status(),
			"method", gctx.Request.Method,
			"path", path,
			"query", query,
			"ip", gctx.ClientIP(),
			"user-agent", gctx.Request.UserAgent(),
			"errors", gctx.Errors.ByType(gin.ErrorTypePrivate).String(),
			"cost", cost,
		)
	}
}

// ginRecover middleware determines whether the stack displays stack information
func ginRecover(stack bool, logger log.Logger) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(gctx.Request, false)
				if brokenPipe {
					logger.Log(log.LevelError,
						"path", gctx.Request.URL.Path,
						"error", err,
						"request", string(httpRequest),
					)
					// If the connection is dead, we can't write a status to it.
					gctx.Error(err.(error)) // nolint: errcheck
					gctx.Abort()
					return
				}

				if stack {
					logger.Log(log.LevelError, "[Recovery from panic]",
						fmt.Sprintf("error:%+v\nrequest:%s\nstack:%s", err, string(httpRequest), string(debug.Stack())))
				} else {
					logger.Log(log.LevelError, "[Recovery from panic]",
						fmt.Sprintf("error:%+v\nrequest:%s\n", err, string(httpRequest)))
				}
				gctx.AbortWithStatus(net_http.StatusInternalServerError)
			}
		}()
		gctx.Next()
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,AC_Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT,DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(net_http.StatusNoContent)
		}
		c.Next()
	}
}
