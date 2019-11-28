import Vue from 'vue'
import Router from 'vue-router'

// 开发环境不使用懒加载, 因为懒加载页面太多的话会造成webpack热更新太慢, 所以只有开发环境使用懒加载
const _import = require('./import-' + process.env.NODE_ENV)

Vue.use(Router)

// 路由定义使用说明:
// 1. 代码中路由统一使用name属性跳转.
// 2. 开放path属性用做简短路由, 比如: '/a1', 访问地址: www.renren.io/#/a1
export default new Router({
  mode: 'hash',
  routes: [
    { path: '/404', component: _import('error/404'), name: '404', desc: '404未找到' },
    { path: '/login', component: _import('login/index'), name: 'login', desc: '登录' },
    {
      path: '/',
      component: _import('layout/index'),
      name: 'layout',
      redirect: { name: 'home' },
      desc: '上左右整体布局',
      children: [
        // 通过isTab属性, 设定是否通过tab标签页展示内容
        { path: '/home', component: _import('home/index'), name: 'home', desc: '首页' },
        { path: '/layout-setting', component: _import('layout/setting'), name: 'setting', desc: '布局设置' },
        { path: '/user', component: _import('user/index'), name: 'user', desc: '管理员管理', meta: { isTab: true } },
        { path: '/role', component: _import('role/index'), name: 'role', desc: '角色管理', meta: { isTab: true } },
        { path: '/menu', component: _import('menu/index'), name: 'menu', desc: '菜单管理', meta: { isTab: true } },
        { path: '/sql', component: _import('sql/index'), name: 'sql', desc: 'SQL监控', meta: { isTab: true } },
        { path: '/schedule', component: _import('schedule/index'), name: 'schedule', desc: '定时任务', meta: { isTab: true } },
        { path: '/config', component: _import('config/index'), name: 'config', desc: '参数管理', meta: { isTab: true } },
        { path: '/oss', component: _import('oss/index'), name: 'oss', desc: '文件上传', meta: { isTab: true } },
        { path: '/log', component: _import('log/index'), name: 'log', desc: '系统日志', meta: { isTab: true } },
        { path: '/app-info', component: _import('app/info/index'), name: 'appinfo', desc: '应用列表',meta: { isTab: true } },
        { path: '/app-conf', component: _import('app/conf/index'), name: 'appconf', desc: '应用配置',meta: { isTab: true } },
        { path: '/app-pay', component: _import('app/pay/index'), name: 'apppay', desc: '支付配置',meta: { isTab: true } },
        { path: '/app-count-survey', component: _import('app/count/survey/index'), name: 'appcountsurvey', desc: '概况趋势',meta: { isTab: true } },
        { path: '/app-count-visit', component: _import('app/count/visit/index'), name: 'appcountvisit', desc: '访问趋势',meta: { isTab: true } },
        { path: '/app-pay-order', component: _import('app/payorder/index'), name: 'apppayorder', desc: '订单列表',meta: { isTab: true } },



        { path: '/mp-car-order', component: _import('miniapp/car/order/index'), name: 'mpcarorder', desc: '订单管理',meta: { isTab: true } },
        { path: '/mp-car-user', component: _import('miniapp/car/user/index'), name: 'mpcaruser', desc: '用户管理',meta: { isTab: true } },
        { path: '/mp-car-deposit', component: _import('miniapp/car/deposit/index'), name: 'mpcardeposit', desc: '提现管理',meta: { isTab: true } },
        { path: '/mp-car-service', component: _import('miniapp/car/service/index'), name: 'mpcarservice', desc: '求助管理',meta: { isTab: true } },


        { path: '/mp-story-user', component: _import('miniapp/story/user/index'), name: 'mpstoryuser', desc: '用户管理',meta: { isTab: true } },
        { path: '/mp-story-order', component: _import('miniapp/story/order/index'), name: 'mpstoryorder', desc: '订单管理',meta: { isTab: true } },
        { path: '/mp-story-rechargecard', component: _import('miniapp/story/rechargecard/index'), name: 'mpstoryrechargecard', desc: '充值卡管理',meta: { isTab: true } },

        { path: '/mp-story-article', component: _import('miniapp/story/article/index'), name: 'mpstoryarticle', desc: '文章管理',meta: { isTab: true } },
        { path: '/mp-story-add-article', component: _import('miniapp/story/articletemplate/addArticle'), name: 'mpstoryAddTemplateArticle', desc: '新增模板文章',meta: { isTab: true } },
        { path: '/mp-storycar-articletemplate', component: _import('miniapp/story/articletemplate/index'), name: 'mpstoryarticletemplate', desc: '模板管理',meta: { isTab: true } },

      ],
      beforeEnter (to, from, next) {
        let token = Vue.cookie.get('token')
        if (!token || !/\S/.test(token)) {
          next({ name: 'login' })
        }
        next()
      }
    },
    { path: '*', redirect: { name: '404' } }
  ]
})
