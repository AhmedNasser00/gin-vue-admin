import service from '@/utils/request'

export const createKpiCategory = (data) => {
  return service({ url: '/kpi/createCategory', method: 'post', data })
}

export const deleteKpiCategory = (data) => {
  return service({ url: '/kpi/deleteCategory', method: 'delete', data })
}

export const deleteKpiCategoryByIds = (data) => {
  return service({ url: '/kpi/deleteCategoryByIds', method: 'delete', data })
}

export const updateKpiCategory = (data) => {
  return service({ url: '/kpi/updateCategory', method: 'put', data })
}

export const findKpiCategory = (params) => {
  return service({ url: '/kpi/findCategory', method: 'get', params })
}

export const getKpiCategoryList = (params) => {
  return service({ url: '/kpi/getCategoryList', method: 'get', params })
}


