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

// @Tags AdEvents
// @Summary 创建AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdEvents true "创建AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adEvents/createAdEvents [post]
func CreateAdEvents(c *gin.Context) {
	var adEvents model.AdEvents
	_ = c.ShouldBindJSON(&adEvents)
	err := service.CreateAdEvents(adEvents)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdEvents
// @Summary 删除AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdEvents true "删除AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adEvents/deleteAdEvents [delete]
func DeleteAdEvents(c *gin.Context) {
	var adEvents model.AdEvents
	_ = c.ShouldBindJSON(&adEvents)
	err := service.DeleteAdEvents(adEvents)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdEvents
// @Summary 批量删除AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adEvents/deleteAdEventsByIds [delete]
func DeleteAdEventsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAdEventsByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdEvents
// @Summary 更新AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdEvents true "更新AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adEvents/updateAdEvents [put]
func UpdateAdEvents(c *gin.Context) {
	var adEvents model.AdEvents
	_ = c.ShouldBindJSON(&adEvents)
	err := service.UpdateAdEvents(&adEvents)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdEvents
// @Summary 用id查询AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdEvents true "用id查询AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adEvents/findAdEvents [get]
func FindAdEvents(c *gin.Context) {
	var adEvents model.AdEvents
	_ = c.ShouldBindQuery(&adEvents)
	err, readEvents := service.GetAdEvents(adEvents.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readEvents": readEvents}, c)
	}
}

// @Tags AdEvents
// @Summary 分页获取AdEvents列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdEventsSearch true "分页获取AdEvents列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adEvents/getAdEventsList [get]
func GetAdEventsList(c *gin.Context) {
	var pageInfo request.AdEventsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAdEventsInfoList(pageInfo)
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
