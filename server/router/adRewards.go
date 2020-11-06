package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdRewardsRouter(Router *gin.RouterGroup) {
	AdRewardsRouter := Router.Group("adRewards").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AdRewardsRouter.POST("createAdRewards", v1.CreateAdRewards)   // 新建AdRewards
		AdRewardsRouter.DELETE("deleteAdRewards", v1.DeleteAdRewards) // 删除AdRewards
		AdRewardsRouter.DELETE("deleteAdRewardsByIds", v1.DeleteAdRewardsByIds) // 批量删除AdRewards
		AdRewardsRouter.PUT("updateAdRewards", v1.UpdateAdRewards)    // 更新AdRewards
		AdRewardsRouter.GET("findAdRewards", v1.FindAdRewards)        // 根据ID获取AdRewards
		AdRewardsRouter.GET("getAdRewardsList", v1.GetAdRewardsList)  // 获取AdRewards列表
	}
}
