import service from '@/utils/request'

// Create Rating
export const createRating = (data) => {
  return service({
    url: '/rating/createRating',
    method: 'post',
    data
  })
}

// Delete Rating (single)
export const deleteRating = (data) => {
  return service({
    url: '/rating/deleteRating',
    method: 'delete',
    data
  })
}

// Delete Rating by IDs (batch)
export const deleteRatingByIds = (data) => {
  return service({
    url: '/rating/deleteRatingByIds',
    method: 'delete',
    data
  })
}

// Update Rating
export const updateRating = (data) => {
  return service({
    url: '/rating/updateRating',
    method: 'put',
    data
  })
}

// Find Rating by ID
export const findRating = (params) => {
  return service({
    url: '/rating/findRating',
    method: 'get',
    params
  })
}

// Get Rating list (paginated)
export const getRatingList = (params) => {
  return service({
    url: '/rating/getRatingList',
    method: 'get',
    params
  })
}

// Get all Ratings
export const getRatingListAll = () => {
  return service({
    url: '/rating/getRatingListAll',
    method: 'get'
  })
}
