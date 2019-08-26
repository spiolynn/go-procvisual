package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-procvisual/middlewares/metric"
	"go-procvisual/pkgs/monitor"
	"go-procvisual/pkgs/setting"
	"go-procvisual/pkgs/upload"
	"go-procvisual/routers/api"
	"go-procvisual/routers/api/v1"
	"net/http"
	"time"
)

// 路由控制层

// 路由注册
func InitRouter() *gin.Engine {

	r := gin.New()

	// 日志中间件
	//r.Use(gin.Logger())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	gin.SetMode(setting.G_cfg_yaml.Server.Runmode)

	// add 报文耗时中间件
	r.Use(metric.Timing)

	// add 性能信息
	monitor.Wrap(r)
	r.GET("/debug/vars", monitor.GetCurrentRunningStats)


	r.GET("/api/v1/users/login", api.CheckAdmin)
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/query", v1.Query)
		apiv1.POST("/admins", v1.QueryAdminList)

		// 通用api
		apiv1.POST("/images", api.UploadImage)
		apiv1.StaticFS("/images", http.Dir(upload.GetImageFullPath()))

		// coderepo
		// To do list
		// 获取具体代码工程、分支提交次数List
		apiv1.GET("code/getcommitInfoByRepo", v1.GetcodeCommitInfoByRepo)
		// 查看项目下代码工程代码贡献率
		apiv1.GET("code/getsharedByRepo", api.UploadImage)
		// 查看项目下代码工程代码贡献率TopN
		apiv1.GET("/codeSharedByRepoTop", api.UploadImage)

		// pipiline
		// To do list
		// 查看流水线历史
		apiv1.GET("/pipelineHistory", api.UploadImage)

	}

	return r
}
