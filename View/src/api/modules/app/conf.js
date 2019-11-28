import request from '../../request'
import requestUrl from '../../requestUrl'
import requestParam from '../../requestParam'
import isInteger from 'lodash/isInteger'

// 获取配置列表
export function list (params) {
  return request({
    url: requestUrl('/app/conf/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取配置信息
export function info (id) {
  return request({
    url: requestUrl('/app/conf/info/' + id),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 添加配置
export function add (params) {
  return request({
    url: requestUrl('/app/conf/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改配置
export function update (params) {
  return request({
    url: requestUrl('/app/conf/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除应用
export function del (params) {
  return request({
    url: requestUrl('/app/conf/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}
