package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdReportsRouter(Router *gin.RouterGroup) {
	AdReportsRouter := Router.Group("adReports").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AdReportsRouter.POST("createAdReports", v1.CreateAdReports)   // 新建AdReports
		AdReportsRouter.DELETE("deleteAdReports", v1.DeleteAdReports) // 删除AdReports
		AdReportsRouter.DELETE("deleteAdReportsByIds", v1.DeleteAdReportsByIds) // 批量删除AdReports
		AdReportsRouter.PUT("updateAdReports", v1.UpdateAdReports)    // 更新AdReports
		AdReportsRouter.GET("findAdReports", v1.FindAdReports)        // 根据ID获取AdReports
		AdReportsRouter.GET("getAdReportsList", v1.GetAdReportsList)  // 获取AdReports列表
	}
}
