<template>
  <div class="mod-user">
    <el-form :inline="true" :model="dataForm">
      <el-form-item>
        <el-input v-model="dataForm.nickName" @keyup.enter.native="getDataList()" placeholder="用户昵称" clearable></el-input>
      </el-form-item>
      <el-form-item>
        <el-select v-model="appInfoId" class="top-app-select" @change="getDataList" placeholder="选择应用">
          <el-option
            v-for="item in appList"
            :key="item.perms"
            :label="item.name"
            :value="item.perms">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button @click="getDataList()">查询</el-button>
      </el-form-item>
    </el-form>
    <el-table
      :data="dataList"
      border
      v-loading="dataListLoading"
      @sort-change="sortChange"
      :default-sort = "{prop: 'createTime', order: 'descending'}"
      style="width: 100%;">
      <el-table-column
        align="center"
        label="头像">
        <template slot-scope="scope">
          <img :src="handleImg(scope.row.avatarUrl)" width="50" height="50" class="head_pic"/>
        </template>
      </el-table-column>
      <el-table-column
        prop="nickName"
        header-align="center"
        align="center"
        label="昵称">
      </el-table-column>
      <el-table-column
        prop="orderNum"
        header-align="center"
        align="center"
        label="订单号">
      </el-table-column>
      <el-table-column
        prop="goodsName"
        header-align="center"
        align="center"
        label="商品名称">
      </el-table-column>
      <el-table-column
        prop="money"
        header-align="center"
        align="center"
        label="价格">
      </el-table-column>
      <el-table-column
        prop="appName"
        header-align="center"
        align="center"
        label="应用">
      </el-table-column>
      <el-table-column
        prop="status"
        header-align="center"
        align="center"
        label="状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === '待支付'" size="small" type="danger">待支付</el-tag>
          <el-tag v-else-if="scope.row.status === '已支付'" size="small" type="success">已支付</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        prop="createTime"
        header-align="center"
        align="center"
        width="180"
        sortable="custom"
        label="创建">
      </el-table-column>
      <el-table-column
        prop="updateTime"
        header-align="center"
        align="center"
        width="180"
        sortable="custom"
        label="更新">
      </el-table-column>

    </el-table>
    <el-pagination
      @size-change="sizeChangeHandle"
      @current-change="currentChangeHandle"
      :current-page="pageIndex"
      :page-sizes="[10, 20, 50, 100]"
      :page-size="pageSize"
      :total="totalPage"
      layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>
  </div>
</template>

<script>
  import API from '@/api'
  export default {
    data () {
      return {
        appList:[],
        dataForm: {
          nickName: ''
        },
        appInfoId:null,
        dataList: [],
        pageIndex: 1,
        pageSize: 10,
        totalPage: 0,
        dataListLoading: false,
        dataListSelections: [],
        addOrUpdateVisible: false,
        sidx:'createTime',
        order:'DESC'
      }
    },
    activated () {
      //this.getDataList()
    },
    mounted(){
      if (this.appList.length === 0){
        this.initAppList()
      }
    },
    methods: {
      // 获取数据列表
      getDataList () {
        this.dataListLoading = true
        var params = {
          page: this.pageIndex,
          limit: this.pageSize,
          appInfoId: this.appInfoId,
          nickName: this.dataForm.nickName,
          sidx:this.sidx,
          order:this.order
        }
        API.appPayOrder.list(params).then(({data}) => {
          if (data && data.code === 0) {
            this.dataList = data.page.list
            this.totalPage = data.page.totalCount
          } else {
            this.dataList = []
            this.totalPage = 0
          }
          this.dataListLoading = false
        })
      },
      //初始化应用列表
      initAppList(){
        let appMenuJson = sessionStorage.getItem('appMenuList')
        if (appMenuJson) {
          this.appList = JSON.parse(appMenuJson)
          this.appInfoId = this.appList[0].perms
          this.getDataList()
        }
      },
      handleImg(url){
        return url+"?x-oss-process=image/resize,m_lfit,h_100,w_100";
      },
      // 每页数
      sizeChangeHandle (val) {
        this.pageSize = val
        this.pageIndex = 1
        this.getDataList()
      },
      // 当前页
      currentChangeHandle (val) {
        this.pageIndex = val
        this.getDataList()
      },
      //排序
      sortChange(val){
        let prop = val.prop;
        if (prop){
          this.sidx = prop;
          this.order = val.order=='ascending'?'ASC':'DESC';
          this.getDataList();
        }else {
          this.sidx = null;
          this.order = null
        }
      }
    }
  }
</script>
