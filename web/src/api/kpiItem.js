import service from '@/utils/request'

export const createKpiItem = (data) => {
  return service({ url: '/kpi/createItem', method: 'post', data })
}

export const deleteKpiItem = (data) => {
  return service({ url: '/kpi/deleteItem', method: 'delete', data })
}

export const deleteKpiItemByIds = (data) => {
  return service({ url: '/kpi/deleteItemByIds', method: 'delete', data })
}

export const updateKpiItem = (data) => {
  return service({ url: '/kpi/updateItem', method: 'put', data })
}

export const findKpiItem = (params) => {
  return service({ url: '/kpi/findItem', method: 'get', params })
}

export const getKpiItemList = (params) => {
  return service({ url: '/kpi/getItemList', method: 'get', params })
}


