import request from '../../request'
import requestUrl from '../../requestUrl'
import requestParam from '../../requestParam'
import isInteger from 'lodash/isInteger'

// 获取应用列表
export function list (params) {
  return request({
    url: requestUrl('/app/pay/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取所有简单的支付信息列表
export function allSimpleList() {
  return request({
    url: requestUrl('/app/pay/simple-list'),
    method: 'get',
    params: requestParam(null, 'get')
  })
}

// 获取应用信息
export function info (id) {
  return request({
    url: requestUrl('/app/pay/info/' + id),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 添加应用
export function add (params) {
  return request({
    url: requestUrl('/app/pay/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改应用
export function update (params) {
  return request({
    url: requestUrl('/app/pay/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除应用
export function del (params) {
  return request({
    url: requestUrl('/app/pay/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}
