<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" label-width="80px">
      <el-form-item label="参数名" prop="key">
        <el-input v-model="dataForm.key" placeholder="参数名"></el-input>
      </el-form-item>
      <el-form-item label="参数值" prop="value">
        <el-input rows="6" type="textarea" v-model="dataForm.value" placeholder="参数值"></el-input>
      </el-form-item>
      <el-form-item label="所属应用">
        <el-select filterable class="max-width" v-model="dataForm.appInfoId" placeholder="请选择所属应用">
          <el-option
            v-for="item in appList"
            :key="item.appInfoId"
            :label="item.appName"
            :value="item.appInfoId">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input rows="6" type="textarea" v-model="dataForm.remark" placeholder="备注"></el-input>
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
        appList:[{
          'appInfoId':'null',
          'appName':'系统'
        }],
        visible: false,
        dataForm: {
          id: 0,
          key: '',
          value: '',
          remark: '',
          appInfoId:null
        },
        dataRule: {
          key: [
            { required: true, message: '参数名不能为空', trigger: 'blur' }
          ],
          value: [
            { required: true, message: '参数值不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    mounted(){
      this.initAppList()
    },
    methods: {
      init (id) {
        this.dataForm.id = id || 0
        this.visible = true
        this.$nextTick(() => {
          this.$refs['dataForm'].resetFields()
          if (this.dataForm.id) {
            API.config.info(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.key = data.config.key
                this.dataForm.value = data.config.value
                this.dataForm.remark = data.config.remark
                this.dataForm.appInfoId = data.config.appInfoIdStr
              }
            })
          }
        })
      },
      //初始化应用列表
      initAppList(){
        API.appCount.appList().then(({data}) => {
          if (data && data.code === 0) {
            data.appList.forEach(item => {
              this.appList.push(item);
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
              'key': this.dataForm.key,
              'value': this.dataForm.value,
              'remark': this.dataForm.remark,
              'appInfoId':this.dataForm.appInfoId
            }
            var tick = !this.dataForm.id ? API.config.add(params) : API.config.update(params)
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
