import service from '@/utils/request'

// @Tags AdReports
// @Summary 创建AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdReports true "创建AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adReports/createAdReports [post]
export const createAdReports = (data) => {
     return service({
         url: "/adReports/createAdReports",
         method: 'post',
         data
     })
 }


// @Tags AdReports
// @Summary 删除AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdReports true "删除AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adReports/deleteAdReports [delete]
 export const deleteAdReports = (data) => {
     return service({
         url: "/adReports/deleteAdReports",
         method: 'delete',
         data
     })
 }

// @Tags AdReports
// @Summary 删除AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adReports/deleteAdReports [delete]
 export const deleteAdReportsByIds = (data) => {
     return service({
         url: "/adReports/deleteAdReportsByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdReports
// @Summary 更新AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdReports true "更新AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adReports/updateAdReports [put]
 export const updateAdReports = (data) => {
     return service({
         url: "/adReports/updateAdReports",
         method: 'put',
         data
     })
 }


// @Tags AdReports
// @Summary 用id查询AdReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdReports true "用id查询AdReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adReports/findAdReports [get]
 export const findAdReports = (params) => {
     return service({
         url: "/adReports/findAdReports",
         method: 'get',
         params
     })
 }


// @Tags AdReports
// @Summary 分页获取AdReports列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdReports列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adReports/getAdReportsList [get]
 export const getAdReportsList = (params) => {
     return service({
         url: "/adReports/getAdReportsList",
         method: 'get',
         params
     })
 }