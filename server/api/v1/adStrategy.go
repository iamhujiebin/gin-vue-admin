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

// @Tags AdStrategy
// @Summary 创建AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdStrategy true "创建AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adStrategy/createAdStrategy [post]
func CreateAdStrategy(c *gin.Context) {
	var adStrategy model.AdStrategy
	_ = c.ShouldBindJSON(&adStrategy)
	err := service.CreateAdStrategy(adStrategy)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdStrategy
// @Summary 删除AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdStrategy true "删除AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adStrategy/deleteAdStrategy [delete]
func DeleteAdStrategy(c *gin.Context) {
	var adStrategy model.AdStrategy
	_ = c.ShouldBindJSON(&adStrategy)
	err := service.DeleteAdStrategy(adStrategy)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdStrategy
// @Summary 批量删除AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adStrategy/deleteAdStrategyByIds [delete]
func DeleteAdStrategyByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAdStrategyByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdStrategy
// @Summary 更新AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdStrategy true "更新AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adStrategy/updateAdStrategy [put]
func UpdateAdStrategy(c *gin.Context) {
	var adStrategy model.AdStrategy
	_ = c.ShouldBindJSON(&adStrategy)
	err := service.UpdateAdStrategy(&adStrategy)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdStrategy
// @Summary 用id查询AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdStrategy true "用id查询AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adStrategy/findAdStrategy [get]
func FindAdStrategy(c *gin.Context) {
	var adStrategy model.AdStrategy
	_ = c.ShouldBindQuery(&adStrategy)
	err, readStrategy := service.GetAdStrategy(adStrategy.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readStrategy": readStrategy}, c)
	}
}

// @Tags AdStrategy
// @Summary 分页获取AdStrategy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdStrategySearch true "分页获取AdStrategy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adStrategy/getAdStrategyList [get]
func GetAdStrategyList(c *gin.Context) {
	var pageInfo request.AdStrategySearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAdStrategyInfoList(pageInfo)
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
