# goAdmin

#### 介绍
gin + jwt +vue  权限框架

#### 软件架构
软件架构说明
  #####1.技术框架 gin
  #####2.ORM框架 gorm
  #####3.token机制 jwt
  #####4.缓存 go-redis
  #####5.前端框架 vue  

  
#### 安装教程

##### 1.  环境
 ###### go 版本 1.13
 ###### mysql 5.6以上
 ###### redis  
 
##### 2.  环境配置
 ###### 2.1 自行修改 src/resources/目录下的 bootstrap.yaml 对应的端口 和使用环境 默认端口：7001 环境：dev
 ###### 2.2 自行修改自己对应环境的 mysql配置、redis配置
 ###### 2.3 动态配置是否开启验证码、权限校验对应环境配置(便于开发联调)
 
##### 3.  开发工具
 ###### goland 开发工具
 
##### 4.   线上部署
 ###### 4.1 动态配置是否开启验证码、权限校验对应环境配置(便于开发联调)

##### 5.   vue前端代码地址 https://cool-admin.com/
 
#### 使用说明

###### 1.  启动 项目的main方法：   src/main/AdminApplication.go 
###### 2.  接口访问根路径：        http://localhost:7001/${Iris.BasePath--为对应各个环境配置的基础地址}/${具体controller路由地址}


#### 码云特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  码云官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解码云上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是码云最有价值开源项目，是码云综合评定出的优秀开源项目
5.  码云官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  码云封面人物是一档用来展示码云会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
