package router

import (
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	"net/http"
	"strings"
	"time"

	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/middleware/recover"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/cnartlu/area-service/pkg/swagger"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"

	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// New 返回 gin 路由对象
func New(
	loggerWriter *rotatelogs.RotateLogs,
	logger log.Logger,
	appConf *config.Application,
	httpConf *config.Server_HTTP,
) *gin.Engine {
	if httpConf == nil {
		return nil
	}

	var output io.Writer
	if loggerWriter == nil {
		output = os.Stdout
	} else {
		output = io.MultiWriter(os.Stdout, loggerWriter)
	}
	gin.DefaultWriter = output
	gin.DefaultErrorWriter = output
	gin.DisableConsoleColor()

	switch appConf.Env {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(ginzap.Ginzap(logger.Zap("").WithOptions(zap.AddCallerSkip(4)), time.RFC3339, false))
	router.Use(recover.CustomRecoveryWithZap(logger.Zap("").WithOptions(zap.AddCallerSkip(4)), true, func(c *gin.Context, err interface{}) {
		// response.Error(c, errors.ServerError())
		c.Abort()
	}))
	router.Use(otelgin.Middleware(appConf.Name))

	rg := router.Group("/")
	extAddrSubs := strings.SplitN(httpConf.Addr, "/", 2)
	if len(extAddrSubs) == 2 {
		rg = router.Group("/" + extAddrSubs[1])
	}

	rg.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong"); return })
	// 注册 api 路由组
	apiGroup := rg.Group("/api")
	{
		apiGroup.Use(cors.Default()) // 允许跨越
		// if jwtConf != nil {
		// 	if jwtConf.Key != "" {
		// 		apiGroup.Use(jwtmd.New(
		// 			jwtConf.Key,
		// 			jwtmd.WithLogger(log.NewHelper(logger)),
		// 			jwtmd.WithErrorResponseBody(response.NewBody(int(errors.ServerErrorCode), errors.ServerErrorCode.String(), nil)),
		// 			jwtmd.WithValidateFailedResponseBody(response.NewBody(int(errors.UnauthorizedCode), errors.UnauthorizedCode.String(), nil)),
		// 		).Validate())
		// 	}
		// }
		// if enforcer != nil {
		// 	apiGroup.Use(casbinmd.New(
		// 		enforcer,
		// 		func(ctx *gin.Context) ([]interface{}, error) {
		// 			// TODO
		// 			return nil, nil
		// 		},
		// 		casbinmd.WithLogger(log.NewHelper(logger)),
		// 		casbinmd.WithErrorResponseBody(response.NewBody(int(errors.ServerErrorCode), errors.ServerErrorCode.String(), nil)),
		// 		casbinmd.WithValidateFailedResponseBody(response.NewBody(int(errors.UnauthorizedCode), errors.UnauthorizedCode.String(), nil)),
		// 	).Validate())
		// }

		// swagger 配置
		if appConf.Env == "local" {
			docs.SwaggerInfo.Host = httpConf.Addr
			if len(extAddrSubs) > 0 {
				docs.SwaggerInfo.Host = extAddrSubs[0]
			}
			docs.SwaggerInfo.BasePath = apiGroup.BasePath()

			swagger.Setup(router, swagger.Config{
				Path: apiGroup.BasePath() + "/docs",
				Option: func(c *ginSwagger.Config) {
					c.DefaultModelsExpandDepth = -1
				},
			})
		}

		// apiV1Group := apiGroup.Group("/v1")
		// {
		// 	// TODO 编写路由
		// }
	}

	return router
}
