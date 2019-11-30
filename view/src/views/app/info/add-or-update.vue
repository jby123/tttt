<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="应用名称" prop="appName">
        <el-input v-model="dataForm.appName" placeholder="请输入应用名称"></el-input>
      </el-form-item>

      <el-form-item label="支付配置">
        <el-select filterable class="max-width" v-model="dataForm.appPayId" placeholder="请选择所属应用">
          <el-option
            v-for="item in appPayList"
            :key="item.id"
            :label="item.name"
            :value="item.id">
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="应用描述" prop="description" size="500">
        <el-input rows="6" type="textarea" v-model="dataForm.description" placeholder="请输入应用描述"></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-radio-group v-model="dataForm.status">
          <el-radio label="规划中">规划中</el-radio>
          <el-radio label="开发中">开发中</el-radio>
          <el-radio label="已上线">已上线</el-radio>
          <el-radio label="已下线">已下线</el-radio>
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
        appPayList:[{
          'id':null,
          'name':'无'
        }],
        visible: false,
        dataForm: {
          id: '',
          appName: '',
          appPayId:null,
          description: '',
          status: '已上线'
        },
        dataRule: {
          appName: [
            { required: true, message: '应用名称不能为空', trigger: 'blur' }
          ],
          description: [
            { required: true, message: '应用描述不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    mounted(){
      this.initPayList()
    },
    methods: {
      init (id) {
        this.dataForm.id = id || ''
        this.visible = true
        this.$nextTick(() => {
          this.$refs['dataForm'].resetFields()
          if (this.dataForm.id) {
            API.appInfo.info(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                //this.dataForm.id = data.appInfo.id
                this.dataForm.appPayId = data.appInfo.appPayId
                this.dataForm.appName = data.appInfo.appName
                this.dataForm.description = data.appInfo.description
                this.dataForm.status = data.appInfo.status
              }
            })
          }
        })
      },
      //初始化应用列表
      initPayList(){
        API.appPay.allSimpleList().then(({data}) => {
          if (data && data.code === 0) {
            data.appPayList.forEach(item => {
              this.appPayList.push(item);
            });
          }
        })
      },
      // 表单提交
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            var params = {
              'id': this.dataForm.id || undefined,
              'appPayId':this.dataForm.appPayId,
              'appName': this.dataForm.appName,
              'description': this.dataForm.description,
              'status': this.dataForm.status
            }
            var tick = !this.dataForm.id ? API.appInfo.add(params) : API.appInfo.update(params)
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
