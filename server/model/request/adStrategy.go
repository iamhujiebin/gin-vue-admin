package request

import "gin-vue-admin/model"

type AdStrategySearch struct{
    model.AdStrategy
    PageInfo
}