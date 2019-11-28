import * as common from './modules/common'
import * as user from './modules/user'
import * as role from './modules/role'
import * as menu from './modules/menu'
import * as log from './modules/log'
import * as config from './modules/config'
import * as oss from './modules/oss'
import * as schedule from './modules/schedule'
import * as appInfo from './modules/app/info'
import * as appConf from './modules/app/conf'
import * as appPay from './modules/app/pay'
import * as appCount from './modules/app/count'
import * as appPayOrder from './modules/app/payorder'




export default {
  common,     // 公共
  user,       // 管理员管理
  role,       // 角色管理
  menu,       // 菜单管理
  log,        // 系统日志
  config,     // 参数管理
  oss,        // 文件服务
  schedule,   // 定时任务
  appInfo,    // 应用管理
  appConf,    // 应用配置
  appPay,     // 应用支付配置
  appCount,   // 应用统计
  appPayOrder,   // 应用支付订单

}
