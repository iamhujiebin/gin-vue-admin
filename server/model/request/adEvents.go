package request

import "gin-vue-admin/model"

type AdEventsSearch struct{
    model.AdEvents
    PageInfo
}