<template>
  <nav class="site-navbar" :class="navbarClasses">
    <div class="site-navbar__header">
      <h1 class="site-navbar__brand" @click="$router.push({ name: 'home' })">
        <a class="site-navbar__brand-lg" href="javascript:;">LEHA PLATFORM</a>
        <a class="site-navbar__brand-mini" href="javascript:;">LEHA</a>
      </h1>
    </div>
    <div class="site-navbar__body clearfix">
      <el-menu
        class="site-navbar__menu"
        mode="horizontal">
        <el-menu-item class="site-navbar__switch" index="0" @click="switchSidebarCollapseHandle()">
          <icon-svg name="zhedie"></icon-svg>
        </el-menu-item>
      </el-menu>
      <el-menu
        class="site-navbar__menu site-navbar__menu--right"
        mode="horizontal">
        <el-menu-item index="4" v-show="showExit" @click="exitMiniapp">
          <el-badge value="退出">
             <i class="el-icon-back"></i>
          </el-badge>
        </el-menu-item>
        <el-submenu v-show="showSubmenu" index="3">
          <template slot="title">{{submenuName}}</template>
          <el-menu-item @click="changeLeftMenu(item.name)" v-for="(item, index) in appMenuList" :key='index' :index="getIndex(index)"><a href="javaScript:;">{{item.name}}</a></el-menu-item>
        </el-submenu>
        <el-menu-item index="1" @click="$router.push({ name: 'setting' })">
          <template slot="title">
            <el-badge value="new">
              <icon-svg name="shezhi" class="el-icon-setting"></icon-svg>
            </el-badge>
          </template>
        </el-menu-item>
        <el-menu-item class="site-navbar__avatar" index="3">
          <el-dropdown :show-timeout="0" placement="bottom">
            <span class="el-dropdown-link">
              <img src="~@/assets/img/avatar.png" :alt="$store.state.user.name">
              {{ $store.state.user.name }}
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="updatePasswordHandle()">修改密码</el-dropdown-item>
              <el-dropdown-item @click.native="logoutHandle()">退出</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </el-menu-item>
      </el-menu>
    </div>
    <!-- 弹窗, 修改密码 -->
    <update-password v-if="updatePassowrdVisible" ref="updatePassowrd"></update-password>
  </nav>
</template>

<script>
  import UpdatePassword from './update-password'
  import API from '@/api'
  import { mapMutations } from 'vuex'
  export default {
    data () {
      return {
        updatePassowrdVisible: false,
        appMenuList: [],
        leftMenuList: [],
        newLeftMenuList: [],
        submenuName: '小程序管理',
        showSubmenu: false,
        showExit: false
      }
    },
    components: {
      UpdatePassword
    },
    mounted () {
      this.$router.push('home')
      this.MyEvent.$on('menu-init', (text) => {
        this.initMenuList()
      })
    },
    computed: {
      navbarClasses () {
        let type = this.$store.state.navbarLayoutType
        return [
          !/\S/.test(type) || type === 'default' ? '' : `site-navbar--${type}`
        ]
      }
    },
    methods: {
      exitMiniapp () {
        this.submenuName = '小程序管理'
        this.showExit = false
        this.UPDATE_MENU_NAV_LIST(this.leftMenuList)
        this.$router.push('home')
      },
      changeLeftMenu (menuName) {
        this.newLeftMenuList = []
        this.submenuName = menuName
        this.appMenuList.forEach(item => {
          if (item.name === menuName) {
            this.newLeftMenuList.push(item)
          }
        })
        this.showExit = true
        this.UPDATE_MENU_NAV_LIST(this.newLeftMenuList)
      },
      getIndex (index) {
        return '2-' + index
      },
      initMenuList () {
        let appMenuJson = sessionStorage.getItem('appMenuList')
        if (appMenuJson) {
          this.appMenuList = JSON.parse(appMenuJson)
        }
        if (this.appMenuList.length>0){
          this.showSubmenu = true;
        }
        let leftMenuJson = sessionStorage.getItem('leftMenuList')
        if (leftMenuJson) {
          this.leftMenuList = JSON.parse(leftMenuJson)
        }
      },
      // 切换侧边栏, 水平折叠收起状态
      switchSidebarCollapseHandle () {
        this.SWITCH_SIDEBAR_COLLAPSE({ collapse: !this.$store.state.sidebarCollapse })
      },
      // 修改密码
      updatePasswordHandle () {
        this.updatePassowrdVisible = true
        this.$nextTick(() => {
          this.$refs.updatePassowrd.init()
        })
      },
      // 退出
      logoutHandle () {
        this.$confirm(`确定进行[退出]操作?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          API.common.logout().then(({data}) => {
            if (data && data.code === 0) {
              this.DELETE_CONTENT_TABS()
              this.$cookie.delete('token')
              this.$router.replace({ name: 'login' })
            }
          })
        }).catch(() => {})
      },
      ...mapMutations(['UPDATE_MENU_NAV_LIST','SWITCH_SIDEBAR_COLLAPSE', 'DELETE_CONTENT_TABS'])
    }

  }
</script>
