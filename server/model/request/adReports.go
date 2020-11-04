package request

import "gin-vue-admin/model"

type AdReportsSearch struct{
    model.AdReports
    PageInfo
}