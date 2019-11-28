<template>
  <div v-if="homeShow" class="mod-home" v-loading="loading" element-loading-text="获得统计数据中">
    <el-row class="top-app">
      <el-col>
        <div>
          <div style="float:left;cursor:pointer;" @click="initData()">
            <icon-svg name="refresh" class="top-app-count-icon"></icon-svg>
            <span class="top-app-count-icon-text">刷新</span>
          </div>
          <div>
            <icon-svg name="count" class="top-app-count-icon" style="margin-left: 10px"></icon-svg>
            <span class="top-app-count-icon-text">实时模式</span>
            <el-select v-model="appInfoId" class="top-app-select" @change="initData" placeholder="选择应用">
              <el-option
                v-for="item in appList"
                :key="item.perms"
                :label="item.name"
                :value="item.perms">
              </el-option>
            </el-select>
          </div>
        </div>
      </el-col>
    </el-row>


    <el-row :gutter="16">
      <el-col :span="6">
        <el-card shadow="always">
          <div style="padding-bottom: 53px">
            <h3 style="float: left">用户总数</h3>
            <h3 style="float: right"><span class="top-car-count">{{countData.userCount}}人</span></h3>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="always">
          <div style="padding-bottom: 53px">
            <h3 style="float: left">总成交金额</h3>
            <h3 style="float: right"><span class="top-car-count">{{countData.orderTotalMoney}}元</span></h3>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="always">
          <div style="padding-bottom: 53px">
            <h3 style="float: left">应用总数</h3>
            <h3 style="float: right"><span class="top-car-count">{{countData.appCount}}个</span></h3>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="always">
          <div style="padding-bottom: 53px">
            <h3 style="float: left">IM在线</h3>
            <h3 style="float: right"><span class="top-car-count">{{countData.imUserCount}}人</span></h3>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-row class="row-count">
        <el-card shadow="always">
            <el-col :span="12">
              <div id="chartOrderMoney" class="chart-order-money"></div>
            </el-col>
            <el-col :span="12">
                <div id="chartUserMouth" class="chart-user-mouth"></div>
            </el-col>
        </el-card>
    </el-row>
  </div>
  <div class="no-authority-f" v-else v-show="showNoAuthority">
    <div class="no-authority-c">
      <img src="../../../static/img/home/no-authority.svg"/>
      <h3 style="text-align: center">无查看首页数据权限</h3>
    </div>
  </div>

</template>

<script>
  import API from '@/api'
  export default {
    name: 'hello',
    data () {
      return {
        appMenuList:[],
        appList:[{
          'perms':null,
          'name':'全部'
        }],
        loading:false,
        appInfoId:null,
        dayCount:7,//默认查询最近七天的数据
        countData:{},
        msg: 'Welcome to Your Vue.js App',
        homeShow: true,
        showNoAuthority: false,
      }
    },
    mounted(){
      //this.initAppList();
      //监听菜单初始化完成
      this.MyEvent.$on('menu-init', (text) => {
        if (!this.initStatus){
          this.initData()
          this.initAppList()
        }
      })
      if (this.appList.length === 1){
        this.initAppList()
        if ((sessionStorage.getItem('permissions') || '[]')!=='[]') {
          this.initData()
        }
      }
    },
    methods: {
      initData(){
        this.loading = true;
        var params = {
          'appInfoId':this.appInfoId,
          'dayCount':this.dayCount
        }
        if (this.isAuth('sys:comm:home')) {
          API.appCount.home(params).then(({data}) => {
            this.loading = false;
            if (data && data.code === 0) {
              this.countData = data.countData
              this.drawUserLineCount(data.countData.userLineCount);
              this.drawOrderBarCount(data.countData.orderMoneyBarCount);
            }else {
              this.homeShow = false
              this.showNoAuthority = true
            }
          })
        }else {
          this.homeShow = false
          this.showNoAuthority = true
        }
      },
      //初始化应用列表
      initAppList(){
        if (this.appList.length === 1){
          let appMenuJson = sessionStorage.getItem('appMenuList')
          if (appMenuJson) {
            this.appMenuList = JSON.parse(appMenuJson)
            this.appMenuList.forEach(item => {
              this.appList.push(item);
            });
          }
        }
      },
      drawUserLineCount(userLineCount){
        // 基于准备好的dom，初始化echarts实例
        let myChart = this.$echarts.init(document.getElementById('chartUserMouth'),'light')
        let option = {
          title : {
            text: '新增用户',
            subtext: '最近7天'
          },
          tooltip : {
            trigger: 'axis'
          },
          calculable : true,
          xAxis : [
            {
              type : 'category',
              data : userLineCount.days
            }
          ],
          yAxis : [
            {
              type : 'value'
            }
          ],
          series : [
            {
              name:'新增用户',
              type:'line',
              data: userLineCount.counts,
              markPoint : {
                data : [
                  {type : 'max', name: '最大值'},
                  {type : 'min', name: '最小值'}
                ]
              },
              markLine : {
                data : [
                  {type : 'average', name: '平均值'}
                ]
              }
            }
          ]
        };
        // 绘制图表
        myChart.setOption(option);
      },
      drawOrderBarCount(orderMoneyBarCount){
        // 基于准备好的dom，初始化echarts实例
        let myChart = this.$echarts.init(document.getElementById('chartOrderMoney'),'light')
        let option = {
          title : {
            text: '成交金额(元)',
            subtext: '最近7天'
          },
          tooltip : {
            trigger: 'axis'
          },
          calculable : true,
          xAxis : [
            {
              type : 'category',
              data : orderMoneyBarCount.days
            }
          ],
          yAxis : [
            {
              type : 'value'
            }
          ],
          series : [
            {
              name:'成交',
              type:'bar',
              data:orderMoneyBarCount.moneys,
              markPoint : {
                data : [
                  {type : 'max', name: '最大值'},
                  {type : 'min', name: '最小值'}
                ]
              },
              markLine : {
                data : [
                  {type : 'average', name: '平均值'}
                ]
              }
            }
          ]
        };
        // 绘制图表
        myChart.setOption(option);
      }
    }
  }
</script>

<style>
  .top-app-count-icon-text{
    color: #6f7180;
  }
  .top-app-count-icon{
    width: 26px;
    height: 26px;
    position: relative;
    top: 50%;
    transform: translateY(30%);
  }
  .top-app-select{
    float: right;
  }
  .top-app{
    margin-bottom: 10px;
  }
  .top-car-count{
    color: #1c84c6;
  }
  .mod-home {
    line-height: 1.5;
  }
  .row-count{
    margin-top: 10px;
  }
  .chart-user-mouth{
    width: 100%;
    height: 350px;
    margin: 10px;
  }
  .chart-order-money{
    width: 100%;
    height: 350px;
    margin: 10px;
  }
  .no-authority-f{
    width:100%;
    height:500px;
    display:flex;
    justify-content:center;
    align-items:center;
  }
  .no-authority-c{
    width:200px;
    height:200px;
  }
</style>

