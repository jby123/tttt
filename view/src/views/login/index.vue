<template>
<div class="htmleaf-container">
  <div class="wrapper">
    <div class="container">
      <h1>17car管理平台</h1>

      <form id="loginForm" style="width: 602px" data-parsley-validate="" novalidate=""
            role="form"
            method="post">
        <div>
          <input v-model="dataForm.userName" id="username" name="username" placeholder="请输入用户名" type="text">
          <input id="password" v-model="dataForm.password" name="password"  placeholder="请输入密码" type="password">
        </div>
        <div style="width: 390px;margin-left: 118px">
          <input type="text"  v-model="dataForm.captcha"   style="float: left;width: 169px" placeholder="验证码">
          <img id="captcha_img" :src="captchaPath" style="height:43px;cursor: pointer;" @click="getCaptcha"/>
        </div>
        <br><br><br>
        <div>
          <button @click="dataFormSubmit"  style="height: 50px;font-size: 20px" id="ok" type="button">{{btnText}}</button>
        </div>


        <br><br><br>
        <label style="color: white">Copyright ©  YouYuNet. All Rights Reserved. v1.0.1</label>
      </form>
    </div>


    <ul class="bg-bubbles">
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
    </ul>
  </div>
</div>
</template>

<script>
  import API from '@/api'
  import { getUUID } from '@/utils'
  export default {
    data () {
      return {
        dataForm: {
          userName: '',
          password: '',
          uuid: '',
          captcha: ''
        },
        captchaPath: '',
        btnText: '登录'
      }
    },
    created () {
      this.getCaptcha()
    },
    mounted () {
      document.onkeydown=function(event) {
        var e = event || window.event || arguments.callee.caller.arguments[0]
        if(e && e.keyCode==13) {
          var oBtn = document.getElementById('ok')
          oBtn.click()
        }
      }
    },
    methods: {
      // 提交表单
      dataFormSubmit () {
        if (!this.dataForm.userName) {
          this.$message({
            message: '用户名不能为空',
            type: 'warning'
          })
          return
        }
        if (!this.dataForm.password) {
          this.$message({
            message: '密码不能为空',
            type: 'warning'
          })
          return
        }
        if (!this.dataForm.captcha) {
          this.$message({
            message: '验证码不能为空',
            type: 'warning'
          })
          return
        }
        var params = {
          'username': this.dataForm.userName,
          'password': this.dataForm.password,
          'uuid': this.dataForm.uuid,
          'captcha': this.dataForm.captcha
        }
        this.btnText = '登录中...'
        API.common.login(params).then(({data}) => {
          if (data && data.code === 0) {
            this.btnText = '登录成功'
            this.$cookie.set('token', data.token, { expires: `${data.expire || 0}s` })
            this.$router.replace({ name: 'home' })
          } else {
            this.btnText = '登录失败'
            this.getCaptcha()
            this.$message.error(data.msg)
          }
        })
      },
      // 获取验证码
      getCaptcha () {
        this.dataForm.uuid = getUUID()
        this.captchaPath = API.common.captcha(this.dataForm.uuid)
      }
    }
  }
</script>

<style scoped lang="scss">
  .wrapper {
    background: #50a3a2;
    background: -webkit-linear-gradient(top left, #50a3a2 0%, #53e3a6 100%);
    background: linear-gradient(to bottom right, #50a3a2 0%, #53e3a6 100%);
    opacity: 0.8;
    position: absolute;
    top: 50%;
    left: 0;
    width: 100%;
    height: 600px;
    margin-top: -293px;
    overflow: hidden;

  }

  .wrapper.form-success .container h1 {
    -webkit-transform: translateY(85px);
    -ms-transform: translateY(85px);
    transform: translateY(85px);
  }
  .container {
    max-width: 600px;
    margin: 0 auto;
    padding: 80px 0;
    height: 400px;
    text-align: center;
  }
  .container h1 {
    font-size: 40px;
    -webkit-transition-duration: 1s;
    transition-duration: 1s;
    -webkit-transition-timing-function: ease-in-put;
    transition-timing-function: ease-in-put;
    font-weight: 200;
    color: white;
  }
  form {
    padding: 20px 0;
    position: relative;
    z-index: 2;
  }
  form input {
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    outline: 0;
    border: 1px solid rgba(255, 255, 255, 0.4);
    width: 366px;
    border-radius: 3px;
    padding: 10px 15px;
    margin: 0 auto 10px auto;
    display: block;
    text-align: center;
    font-size: 18px;
    color: black;
    -webkit-transition-duration: 0.25s;
    transition-duration: 0.25s;
    font-weight: 300;
  }
  form input:hover {
    background-color: rgba(255, 255, 255, 0.4);
  }
  form input:focus {
    background-color: white;
    width: 400px;
    color: black;
  }
  form button {
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    outline: 0;
    background-color: white;
    border: 0;
    padding: 10px 15px;
    color: #53e3a6;
    border-radius: 3px;
    width: 366px;
    cursor: pointer;
    font-size: 18px;
    -webkit-transition-duration: 0.25s;
    transition-duration: 0.25s;
  }
  form button:hover {
    background-color: #f5f7f9;
  }
  .bg-bubbles {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 80%;
    z-index: 1;
  }
  .bg-bubbles li {
    position: absolute;
    list-style: none;
    display: block;
    width: 40px;
    height: 40px;
    background-color: rgba(255, 255, 255, 0.15);
    bottom: -160px;
    -webkit-animation: square 25s infinite;
    animation: square 25s infinite;
    -webkit-transition-timing-function: linear;
    transition-timing-function: linear;
  }
  .bg-bubbles li:nth-child(1) {
    left: 10%;
  }
  .bg-bubbles li:nth-child(2) {
    left: 20%;
    width: 80px;
    height: 80px;
    -webkit-animation-delay: 2s;
    animation-delay: 2s;
    -webkit-animation-duration: 17s;
    animation-duration: 17s;
  }
  .bg-bubbles li:nth-child(3) {
    left: 25%;
    -webkit-animation-delay: 4s;
    animation-delay: 4s;
  }
  .bg-bubbles li:nth-child(4) {
    left: 40%;
    width: 60px;
    height: 60px;
    -webkit-animation-duration: 22s;
    animation-duration: 22s;
    background-color: rgba(255, 255, 255, 0.25);
  }
  .bg-bubbles li:nth-child(5) {
    left: 70%;
  }
  .bg-bubbles li:nth-child(6) {
    left: 80%;
    width: 120px;
    height: 120px;
    -webkit-animation-delay: 3s;
    animation-delay: 3s;
    background-color: rgba(255, 255, 255, 0.2);
  }
  .bg-bubbles li:nth-child(7) {
    left: 32%;
    width: 160px;
    height: 160px;
    -webkit-animation-delay: 7s;
    animation-delay: 7s;
  }
  .bg-bubbles li:nth-child(8) {
    left: 55%;
    width: 20px;
    height: 20px;
    -webkit-animation-delay: 15s;
    animation-delay: 15s;
    -webkit-animation-duration: 40s;
    animation-duration: 40s;
  }
  .bg-bubbles li:nth-child(9) {
    left: 25%;
    width: 10px;
    height: 10px;
    -webkit-animation-delay: 2s;
    animation-delay: 2s;
    -webkit-animation-duration: 40s;
    animation-duration: 40s;
    background-color: rgba(255, 255, 255, 0.3);
  }
  .bg-bubbles li:nth-child(10) {
    left: 90%;
    width: 160px;
    height: 160px;
    -webkit-animation-delay: 11s;
    animation-delay: 11s;
  }
  @-webkit-keyframes square {
    0% {
      -webkit-transform: translateY(0);
      transform: translateY(0);
    }
    100% {
      -webkit-transform: translateY(-700px) rotate(600deg);
      transform: translateY(-700px) rotate(600deg);
    }
  }
  @keyframes square {
    0% {
      -webkit-transform: translateY(0);
      transform: translateY(0);
    }
    100% {
      -webkit-transform: translateY(-700px) rotate(600deg);
      transform: translateY(-700px) rotate(600deg);
    }
  }

</style>
