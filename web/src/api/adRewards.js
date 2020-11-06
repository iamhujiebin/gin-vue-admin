import service from '@/utils/request'

// @Tags AdRewards
// @Summary 创建AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdRewards true "创建AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adRewards/createAdRewards [post]
export const createAdRewards = (data) => {
     return service({
         url: "/adRewards/createAdRewards",
         method: 'post',
         data
     })
 }


// @Tags AdRewards
// @Summary 删除AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdRewards true "删除AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adRewards/deleteAdRewards [delete]
 export const deleteAdRewards = (data) => {
     return service({
         url: "/adRewards/deleteAdRewards",
         method: 'delete',
         data
     })
 }

// @Tags AdRewards
// @Summary 删除AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adRewards/deleteAdRewards [delete]
 export const deleteAdRewardsByIds = (data) => {
     return service({
         url: "/adRewards/deleteAdRewardsByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdRewards
// @Summary 更新AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdRewards true "更新AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adRewards/updateAdRewards [put]
 export const updateAdRewards = (data) => {
     return service({
         url: "/adRewards/updateAdRewards",
         method: 'put',
         data
     })
 }


// @Tags AdRewards
// @Summary 用id查询AdRewards
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdRewards true "用id查询AdRewards"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adRewards/findAdRewards [get]
 export const findAdRewards = (params) => {
     return service({
         url: "/adRewards/findAdRewards",
         method: 'get',
         params
     })
 }


// @Tags AdRewards
// @Summary 分页获取AdRewards列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdRewards列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adRewards/getAdRewardsList [get]
 export const getAdRewardsList = (params) => {
     return service({
         url: "/adRewards/getAdRewardsList",
         method: 'get',
         params
     })
 }