package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdStrategyRouter(Router *gin.RouterGroup) {
	AdStrategyRouter := Router.Group("adStrategy").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AdStrategyRouter.POST("createAdStrategy", v1.CreateAdStrategy)   // 新建AdStrategy
		AdStrategyRouter.DELETE("deleteAdStrategy", v1.DeleteAdStrategy) // 删除AdStrategy
		AdStrategyRouter.DELETE("deleteAdStrategyByIds", v1.DeleteAdStrategyByIds) // 批量删除AdStrategy
		AdStrategyRouter.PUT("updateAdStrategy", v1.UpdateAdStrategy)    // 更新AdStrategy
		AdStrategyRouter.GET("findAdStrategy", v1.FindAdStrategy)        // 根据ID获取AdStrategy
		AdStrategyRouter.GET("getAdStrategyList", v1.GetAdStrategyList)  // 获取AdStrategy列表
	}
}
