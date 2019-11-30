<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="所属应用">
        <el-select filterable class="max-width" v-model="dataForm.appInfoId" placeholder="请选择所属应用">
          <el-option
            v-for="item in appInfoList"
            :key="item.appInfoId"
            :label="item.appName"
            :value="item.appInfoId">
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="appId" prop="appId">
        <el-input v-model="dataForm.appId" placeholder="请输入appId"></el-input>
      </el-form-item>
      <el-form-item label="appSecret" prop="appSecret">
        <el-input v-model="dataForm.appSecret" placeholder="请输入appSecret"></el-input>
      </el-form-item>
      <el-form-item label="token" prop="token">
      <el-input v-model="dataForm.token" placeholder="请输入token"></el-input>
      </el-form-item>
      <el-form-item label="aesKey" prop="aesKey">
        <el-input v-model="dataForm.aesKey" placeholder="请输入aesKey"></el-input>
      </el-form-item>

      <el-form-item label="消息格式" prop="msgDataFormat">
        <el-radio-group v-model="dataForm.msgDataFormat">
          <el-radio label="JSON">JSON</el-radio>
          <el-radio label="XML">XML</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" @click="dataFormSubmit()">确定</el-button>
    </span>
  </el-dialog>
</template>

<script>
  import API from '@/api'
  export default {
    data () {
      return {
        appInfoList:[],
        visible: false,
        dataForm: {
          id: 0,
          appInfoId: '',
          appId: '',
          appSecret: '',
          token:'',
          aesKey:'',
          msgDataFormat:'JSON'
        },
        dataRule: {
          appId: [
            { required: true, message: 'appId不能为空', trigger: 'blur' }
          ],
          appSecret: [
            { required: true, message: 'appSecret不能为空', trigger: 'blur' }
          ],
          appInfoId: [
            { required: true, message: '所属应用不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    mounted(){
      this.initInfoList()
    },
    methods: {
      init (id) {
        this.dataForm.id = id || 0
        this.visible = true
        this.$nextTick(() => {
          this.$refs['dataForm'].resetFields()
          if (this.dataForm.id) {
            API.appConf.info(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.id = data.appConf.id
                this.dataForm.appInfoId = data.appConf.appInfoIdStr
                this.dataForm.appId = data.appConf.appId
                this.dataForm.appSecret = data.appConf.appSecret
                this.dataForm.token = data.appConf.token
                this.dataForm.aesKey = data.appConf.aesKey
                this.dataForm.msgDataFormat = data.appConf.msgDataFormat
              }
            })
          }
        })
      },
      //初始化应用列表
      initInfoList(){
        API.appInfo.allList().then(({data}) => {
          if (data && data.code === 0) {
            this.appInfoList = data.appInfoList;
          }
        })
      },

      // 表单提交
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            var params = {
              'id': this.dataForm.id || undefined,
              'appInfoId': this.dataForm.appInfoId,
              'appId': this.dataForm.appId,
              'appSecret': this.dataForm.appSecret,
              'token': this.dataForm.token,
              'aesKey': this.dataForm.aesKey,
              'msgDataFormat': this.dataForm.msgDataFormat
            }
            var tick = !this.dataForm.id ? API.appConf.add(params) : API.appConf.update(params)
            tick.then(({data}) => {
              if (data && data.code === 0) {
                this.$message({
                  message: '操作成功',
                  type: 'success',
                  duration: 1500,
                  onClose: () => {
                    this.visible = false
                    this.$emit('refreshDataList')
                  }
                })
              } else {
                this.$message.error(data.msg)
              }
            })
          }
        })
      }
    }
  }
</script>

<style>
  .max-width{
    width: 100%;
  }
</style>

