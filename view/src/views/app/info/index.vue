<template>
  <div class="mod-user">
    <el-form :inline="true" :model="dataForm">
      <el-form-item>
        <el-input v-model="dataForm.appName" @keyup.enter.native="getDataList()" placeholder="应用名称" clearable></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="getDataList()">查询</el-button>
        <el-button v-if="isAuth('app:info:save')" type="primary" @click="addOrUpdateHandle()">新增</el-button>
        <el-button v-if="isAuth('app:info:delete')" type="danger" @click="deleteHandle()" :disabled="dataListSelections.length <= 0">批量删除</el-button>
      </el-form-item>
    </el-form>
    <el-table
      :data="dataList"
      border
      v-loading="dataListLoading"
      @selection-change="selectionChangeHandle"
      @sort-change="sortChange"
      :default-sort = "{prop: 'createTime', order: 'descending'}"
      style="width: 100%;">
      <el-table-column
        type="selection"
        header-align="center"
        align="center"
        width="50">
      </el-table-column>
      <el-table-column
        prop="appInfoId"
        header-align="center"
        align="center"
        width="180"
        show-overflow-tooltip
        label="AppId">
      </el-table-column>
      <el-table-column
        prop="appSecret"
        header-align="center"
        align="center"
        width="180"
        show-overflow-tooltip
        label="AppSecret">
      </el-table-column>
      <el-table-column
        prop="appName"
        header-align="center"
        align="center"
        label="名称">
      </el-table-column>
      <el-table-column
        prop="appPayName"
        align="center"
        header-align="center"
        show-overflow-tooltip
        label="支付配置">
      </el-table-column>
      <el-table-column
        prop="description"
        header-align="center"
        show-overflow-tooltip
        label="描述">
      </el-table-column>
      <el-table-column
        prop="createTime"
        header-align="center"
        align="center"
        width="180"
        sortable="custom"
        label="创建时间">
      </el-table-column>
      <el-table-column
        prop="status"
        header-align="center"
        align="center"
        label="状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === '规划中'" size="small" type="info">规划中</el-tag>
          <el-tag v-else-if="scope.row.status === '开发中'" size="small" type="warning">开发中</el-tag>
          <el-tag v-else-if="scope.row.status === '已上线'" size="small" type="success">已上线</el-tag>
          <el-tag v-else="scope.row.status === '已下线'" size="small" type="danger">已下线</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        fixed="right"
        header-align="center"
        align="center"
        width="150"
        label="操作">
        <template slot-scope="scope">
          <el-button v-if="isAuth('app:info:update')" type="text" size="small" @click="addOrUpdateHandle(scope.row.appInfoId)">修改</el-button>
          <el-button v-if="isAuth('app:info:delete')" type="text" size="small" @click="deleteHandle(scope.row.appInfoId)">删除</el-button>
        </template>
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
    <!-- 弹窗, 新增 / 修改 -->
    <add-or-update v-if="addOrUpdateVisible" ref="addOrUpdate" @refreshDataList="getDataList"></add-or-update>
  </div>
</template>

<script>
  import API from '@/api'
  import AddOrUpdate from './add-or-update'
  export default {
    data () {
      return {
        dataForm: {
          appName: ''
        },
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
    components: {
      AddOrUpdate
    },
    activated () {
      //this.getDataList()
    },
    methods: {
      // 获取数据列表
      getDataList () {
        this.dataListLoading = true
        var params = {
          page: this.pageIndex,
          limit: this.pageSize,
          appName: this.dataForm.appName,
          sidx:this.sidx,
          order:this.order
        }
        API.appInfo.list(params).then(({data}) => {
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
      // 多选
      selectionChangeHandle (val) {
        this.dataListSelections = val
      },
      // 新增 / 修改
      addOrUpdateHandle (id) {
        this.addOrUpdateVisible = true
        this.$nextTick(() => {
          this.$refs.addOrUpdate.init(id)
        })
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
      },
      // 删除
      deleteHandle (id) {
        var ids = id ? [id] : this.dataListSelections.map(item => {
          return item.id
        })
        this.$confirm(`确定对[id=${ids.join(',')}]进行[${id ? '删除' : '批量删除'}]操作?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          API.appInfo.del(ids).then(({data}) => {
            if (data && data.code === 0) {
              this.$message({
                message: '操作成功',
                type: 'success',
                duration: 1500,
                onClose: () => {
                  this.getDataList()
                }
              })
            } else {
              this.$message.error(data.msg)
            }
          })
        })
      }
    }
  }
</script>
