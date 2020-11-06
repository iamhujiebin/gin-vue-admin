package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdEventsRouter(Router *gin.RouterGroup) {
	AdEventsRouter := Router.Group("adEvents").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AdEventsRouter.POST("createAdEvents", v1.CreateAdEvents)   // 新建AdEvents
		AdEventsRouter.DELETE("deleteAdEvents", v1.DeleteAdEvents) // 删除AdEvents
		AdEventsRouter.DELETE("deleteAdEventsByIds", v1.DeleteAdEventsByIds) // 批量删除AdEvents
		AdEventsRouter.PUT("updateAdEvents", v1.UpdateAdEvents)    // 更新AdEvents
		AdEventsRouter.GET("findAdEvents", v1.FindAdEvents)        // 根据ID获取AdEvents
		AdEventsRouter.GET("getAdEventsList", v1.GetAdEventsList)  // 获取AdEvents列表
	}
}
