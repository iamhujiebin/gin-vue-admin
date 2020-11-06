package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags AdRewards
// @Summary 创建AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdRewards true "创建AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adRewards/createAdRewards [post]
func CreateAdRewards(c *gin.Context) {
	var adRewards model.AdRewards
	_ = c.ShouldBindJSON(&adRewards)
	err := service.CreateAdRewards(adRewards)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdRewards
// @Summary 删除AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdRewards true "删除AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adRewards/deleteAdRewards [delete]
func DeleteAdRewards(c *gin.Context) {
	var adRewards model.AdRewards
	_ = c.ShouldBindJSON(&adRewards)
	err := service.DeleteAdRewards(adRewards)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdRewards
// @Summary 批量删除AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adRewards/deleteAdRewardsByIds [delete]
func DeleteAdRewardsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAdRewardsByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdRewards
// @Summary 更新AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdRewards true "更新AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adRewards/updateAdRewards [put]
func UpdateAdRewards(c *gin.Context) {
	var adRewards model.AdRewards
	_ = c.ShouldBindJSON(&adRewards)
	err := service.UpdateAdRewards(&adRewards)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdRewards
// @Summary 用id查询AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdRewards true "用id查询AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adRewards/findAdRewards [get]
func FindAdRewards(c *gin.Context) {
	var adRewards model.AdRewards
	_ = c.ShouldBindQuery(&adRewards)
	err, readRewards := service.GetAdRewards(adRewards.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readRewards": readRewards}, c)
	}
}

// @Tags AdRewards
// @Summary 分页获取AdRewards列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdRewardsSearch true "分页获取AdRewards列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adRewards/getAdRewardsList [get]
func GetAdRewardsList(c *gin.Context) {
	var pageInfo request.AdRewardsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAdRewardsInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
