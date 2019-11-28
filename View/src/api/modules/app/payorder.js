import request from '../../request'
import requestUrl from '../../requestUrl'
import requestParam from '../../requestParam'
import isInteger from 'lodash/isInteger'

// 获取订单列表
export function list (params) {
  return request({
    url: requestUrl('/app/pay/order/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

