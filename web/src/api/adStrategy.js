import service from '@/utils/request'

// @Tags AdStrategy
// @Summary 创建AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdStrategy true "创建AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adStrategy/createAdStrategy [post]
export const createAdStrategy = (data) => {
     return service({
         url: "/adStrategy/createAdStrategy",
         method: 'post',
         data
     })
 }


// @Tags AdStrategy
// @Summary 删除AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdStrategy true "删除AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adStrategy/deleteAdStrategy [delete]
 export const deleteAdStrategy = (data) => {
     return service({
         url: "/adStrategy/deleteAdStrategy",
         method: 'delete',
         data
     })
 }

// @Tags AdStrategy
// @Summary 删除AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adStrategy/deleteAdStrategy [delete]
 export const deleteAdStrategyByIds = (data) => {
     return service({
         url: "/adStrategy/deleteAdStrategyByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdStrategy
// @Summary 更新AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdStrategy true "更新AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adStrategy/updateAdStrategy [put]
 export const updateAdStrategy = (data) => {
     return service({
         url: "/adStrategy/updateAdStrategy",
         method: 'put',
         data
     })
 }


// @Tags AdStrategy
// @Summary 用id查询AdStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdStrategy true "用id查询AdStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adStrategy/findAdStrategy [get]
 export const findAdStrategy = (params) => {
     return service({
         url: "/adStrategy/findAdStrategy",
         method: 'get',
         params
     })
 }


// @Tags AdStrategy
// @Summary 分页获取AdStrategy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdStrategy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adStrategy/getAdStrategyList [get]
 export const getAdStrategyList = (params) => {
     return service({
         url: "/adStrategy/getAdStrategyList",
         method: 'get',
         params
     })
 }