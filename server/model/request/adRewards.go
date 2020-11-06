package request

import "gin-vue-admin/model"

type AdRewardsSearch struct{
    model.AdRewards
    PageInfo
}