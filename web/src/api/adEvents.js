import service from '@/utils/request'

// @Tags AdEvents
// @Summary 创建AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdEvents true "创建AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adEvents/createAdEvents [post]
export const createAdEvents = (data) => {
     return service({
         url: "/adEvents/createAdEvents",
         method: 'post',
         data
     })
 }


// @Tags AdEvents
// @Summary 删除AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdEvents true "删除AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adEvents/deleteAdEvents [delete]
 export const deleteAdEvents = (data) => {
     return service({
         url: "/adEvents/deleteAdEvents",
         method: 'delete',
         data
     })
 }

// @Tags AdEvents
// @Summary 删除AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adEvents/deleteAdEvents [delete]
 export const deleteAdEventsByIds = (data) => {
     return service({
         url: "/adEvents/deleteAdEventsByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdEvents
// @Summary 更新AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdEvents true "更新AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adEvents/updateAdEvents [put]
 export const updateAdEvents = (data) => {
     return service({
         url: "/adEvents/updateAdEvents",
         method: 'put',
         data
     })
 }


// @Tags AdEvents
// @Summary 用id查询AdEvents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdEvents true "用id查询AdEvents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adEvents/findAdEvents [get]
 export const findAdEvents = (params) => {
     return service({
         url: "/adEvents/findAdEvents",
         method: 'get',
         params
     })
 }


// @Tags AdEvents
// @Summary 分页获取AdEvents列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdEvents列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adEvents/getAdEventsList [get]
 export const getAdEventsList = (params) => {
     return service({
         url: "/adEvents/getAdEventsList",
         method: 'get',
         params
     })
 }