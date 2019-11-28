<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="主体名称" prop="name">
        <el-input v-model="dataForm.name" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item label="商户ID" prop="mchId">
        <el-input v-model="dataForm.mchId" placeholder="请输入商户ID"></el-input>
      </el-form-item>
      <el-form-item label="商户秘钥" prop="mchKey">
        <el-input v-model="dataForm.mchKey" placeholder="请输入商户秘钥"></el-input>
      </el-form-item>
      <el-form-item label="证书路径" prop="keyPath">
        <el-input :disabled="true" v-model="dataForm.keyPath" placeholder="请输入证书路径"></el-input>
      </el-form-item>
      <el-form-item label="回调地址" prop="notifyUrl">
        <el-input v-model="dataForm.notifyUrl" placeholder="请输入异步回调地址"></el-input>
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
        visible: false,
        dataForm: {
          id: 0,
          name:'',
          mchId: '',
          mchKey: '',
          keyPath: 'classpath:apiclient_cert.p12',
          notifyUrl:''
        },
        dataRule: {
          name: [
            { required: true, message: '名称不能为空', trigger: 'blur' }
          ],
          mchId: [
            { required: true, message: '商户ID不能为空', trigger: 'blur' }
          ],
          mchKey: [
            { required: true, message: '商户秘钥不能为空', trigger: 'blur' }
          ],
          notifyUrl: [
            { required: true, message: '回调地址不能为空', trigger: 'blur' },
            {pattern: '[a-zA-z]+://[^\\s]*',message: '请输入正确的网址 注：带http://或者https://',trigger: 'blur'}
          ]
        }
      }
    },
    methods: {
      init (id) {
        this.dataForm.id = id || 0
        this.visible = true
        this.$nextTick(() => {
          this.$refs['dataForm'].resetFields()
          if (this.dataForm.id) {
            API.appPay.info(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.id = data.appPay.id
                this.dataForm.name = data.appPay.name
                this.dataForm.mchId = data.appPay.mchId
                this.dataForm.mchKey = data.appPay.mchKey
                this.dataForm.keyPath = data.appPay.keyPath
                this.dataForm.notifyUrl = data.appPay.notifyUrl
              }
            })
          }
        })
      },
      // 表单提交
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            var params = {
              'id': this.dataForm.id || undefined,
              'name':this.dataForm.name,
              'mchId': this.dataForm.mchId,
              'mchKey': this.dataForm.mchKey,
              'keyPath': this.dataForm.keyPath,
              'notifyUrl': this.dataForm.notifyUrl
            }
            var tick = !this.dataForm.id ? API.appPay.add(params) : API.appPay.update(params)
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
