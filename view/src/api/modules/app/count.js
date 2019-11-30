import request from '../../request'
import requestUrl from '../../requestUrl'
import requestParam from '../../requestParam'
import isInteger from 'lodash/isInteger'


//首页统计
export function home (params) {
  return request({
    url: requestUrl('/app/count/home'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取所有应用列表
export function appList () {
  return request({
    url: requestUrl('/app/count/app-list'),
    method: 'get',
    params: requestParam(null, 'get')
  })
}

//概况趋势统计
export function surveyCount (params) {
  return request({
    url: requestUrl('/app/count/survey-count'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

//访问趋势统计
export function visitCount (params) {
  return request({
    url: requestUrl('/app/count/visit-count'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

