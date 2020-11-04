package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"time"
)

// @Tags AdReports
// @Summary 创建AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdReports true "创建AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adReports/createAdReports [post]
func CreateAdReports(c *gin.Context) {
	var adReports model.AdReports
	_ = c.ShouldBindJSON(&adReports)
	err := service.CreateAdReports(adReports)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdReports
// @Summary 删除AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdReports true "删除AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adReports/deleteAdReports [delete]
func DeleteAdReports(c *gin.Context) {
	var adReports model.AdReports
	_ = c.ShouldBindJSON(&adReports)
	err := service.DeleteAdReports(adReports)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdReports
// @Summary 批量删除AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adReports/deleteAdReportsByIds [delete]
func DeleteAdReportsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAdReportsByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdReports
// @Summary 更新AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdReports true "更新AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adReports/updateAdReports [put]
func UpdateAdReports(c *gin.Context) {
	var adReports model.AdReports
	_ = c.ShouldBindJSON(&adReports)
	adReports.UpdatedAt = time.Now()
	err := service.UpdateAdReports(&adReports)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdReports
// @Summary 用id查询AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdReports true "用id查询AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adReports/findAdReports [get]
func FindAdReports(c *gin.Context) {
	var adReports model.AdReports
	_ = c.ShouldBindQuery(&adReports)
	err, readReports := service.GetAdReports(adReports.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readReports": readReports}, c)
	}
}

// @Tags AdReports
// @Summary 分页获取AdReports列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdReportsSearch true "分页获取AdReports列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adReports/getAdReportsList [get]
func GetAdReportsList(c *gin.Context) {
	var pageInfo request.AdReportsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAdReportsInfoList(pageInfo)
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
