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
          <div id="chartSurvey" class="chart-user-mouth"></div>
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
        API.appCount.surveyCount(params).then(({data}) => {
          this.loading = false;
          if (data && data.code === 0) {
            this.drawVisitLineCount(data.data);
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
      drawVisitLineCount(data){
        // 基于准备好的dom，初始化echarts实例
        let myChart = this.$echarts.init(document.getElementById('chartSurvey'),'light')
        let option = {
          title: {
            text: '概况',
            subtext: '最近7天'
          },
          tooltip: {
            trigger: 'axis'
          },
          legend: {
            data:['累计用户数','转发人数','转发次数']
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
              name:'累计用户数',
              type:'bar',
              data: data.counts.visitTotals,
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
              name:'转发次数',
              type:'bar',
              data: data.counts.sharePvs,
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
              name:'转发人数',
              type:'bar',
              data: data.counts.shareUvs,
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

