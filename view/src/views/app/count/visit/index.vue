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
            <span class="top-app-count-icon-text">非实时</span>
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

    <el-row class="row-count">
        <el-card shadow="always">
            <!--<el-col :span="12">-->
              <!--<div id="chartOrderMoney" class="chart-order-money"></div>-->
            <!--</el-col>-->
            <el-col :span="24">
                <div id="chartVisit" class="chart-user-mouth"></div>
            </el-col>
        </el-card>
    </el-row>

    <el-row class="row-count">
      <el-card shadow="always">
        <!--<el-col :span="12">-->
        <!--<div id="chartOrderMoney" class="chart-order-money"></div>-->
        <!--</el-col>-->
        <el-col :span="24">
          <div id="chartVisitTime" class="chart-user-mouth"></div>
        </el-col>
      </el-card>
    </el-row>
  </div>


</template>

<script>
  import API from '@/api'
  export default {
    name: 'hello',
    data () {
      return {
        appMenuList:[],
        appList:[],
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
      if (this.appList.length === 0){
        this.initAppList()
      }
    },
    methods: {
      initData(){
        this.loading = true;
        var params = {
          'appInfoId':this.appInfoId,
          'dayCount':7
        }
          API.appCount.visitCount(params).then(({data}) => {
            this.loading = false;
            if (data && data.code === 0) {
              this.drawVisitBarCount(data.data);
              this.drawVisitTimeBarCount(data.data)
            }else {
              this.homeShow = false
              this.showNoAuthority = true
            }
          })
      },
      //初始化应用列表
      initAppList(){
        let appMenuJson = sessionStorage.getItem('appMenuList')
        if (appMenuJson) {
          this.appList = JSON.parse(appMenuJson)
          this.appInfoId = this.appList[0].perms
          this.initData()
        }
      },
      drawVisitBarCount(data){
        // 基于准备好的dom，初始化echarts实例
        let myChart = this.$echarts.init(document.getElementById('chartVisit'),'light')
        let option = {
          title: {
            text: '访问',
            subtext: '最近7天'
          },
          tooltip: {
            trigger: 'axis'
          },
          legend: {
            data:['打开次数','访问次数','访问人数','新增用户']
          },
          toolbox: {
            show: true,
            feature: {
              dataZoom: {
                yAxisIndex: 'none'
              },
              dataView: {readOnly: false},
              magicType: {type: ['line','bar']},
              restore: {},
              saveAsImage: {}
            }
          },
          xAxis:  {
            type: 'category',
            data: data.days
          },
          yAxis: {
            type: 'value',
          },
          series: [
            {
              name:'打开次数',
              type:'line',
              data: data.counts.sessionCnts,
              markPoint: {
                data: [
                  {type: 'max', name: '最大值'},
                  {type: 'min', name: '最小值'}
                ]
              },
              markLine: {
                data: [
                  {type: 'average', name: '平均值'}
                ]
              }
            },
            {
              name:'访问次数',
              type:'line',
              data: data.counts.visitPvs,
              markPoint: {
                data: [
                  {type: 'max', name: '最大值'},
                  {type: 'min', name: '最小值'}
                ]
              },
              markLine: {
                data: [
                  {type: 'average', name: '平均值'}
                ]
              }
            },
            {
              name:'访问人数',
              type:'line',
              data: data.counts.visitUvs,
              markPoint: {
                data: [
                  {type: 'max', name: '最大值'},
                  {type: 'min', name: '最小值'}
                ]
              },
              markLine: {
                data: [
                  {type: 'average', name: '平均值'}
                ]
              }
            },
            {
              name:'新增用户',
              type:'line',
              data: data.counts.visitUvNews,
              markPoint: {
                data: [
                  {type: 'max', name: '最大值'},
                  {type: 'min', name: '最小值'}
                ]
              },
              markLine: {
                data: [
                  {type: 'average', name: '平均值'}
                ]
              }
            }
          ]
        };
        // 绘制图表
        myChart.setOption(option);
      },
      drawVisitTimeBarCount(data){
        // 基于准备好的dom，初始化echarts实例
        let myChart = this.$echarts.init(document.getElementById('chartVisitTime'),'light')
        let option = {
          title: {
            text: '时间与深度',
            subtext: '最近7天'
          },
          tooltip : {
            trigger: 'axis',
            axisPointer: {
              type: 'cross',
              label: {
                backgroundColor: '#6a7985'
              }
            }
          },
          legend: {
            data:['人均停留时长[秒]','次均停留时长[秒]','平均访问深度']
          },
          toolbox: {
            feature: {
              saveAsImage: {}
            }
          },
          grid: {
            left: '8%',
            right: '9%',
            bottom: '3%',
            containLabel: true
          },
          xAxis : [
            {
              boundaryGap : false,
              type : 'category',
              data : data.days
            }
          ],
          yAxis : [
            {
              type : 'value'
            }
          ],
          series : [
            {
              name:'人均停留时长[秒]',
              type:'line',
              areaStyle: {normal: {}},
              data:data.counts.stayTimeUvs
            },
            {
              name:'次均停留时长[秒]',
              type:'line',
              areaStyle: {normal: {}},
              data:data.counts.stayTimeSessions
            },
            {
              name:'平均访问深度',
              type:'line',
              areaStyle: {normal: {}},
              data:data.counts.visitDepths
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

