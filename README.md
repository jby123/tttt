# goAdmin

#### 介绍
gin + gorm + jwt + vue  权限框架

#### 软件架构

##### 1.技术框架 gin
##### 2.ORM框架 gorm
##### 3.token机制 jwt
##### 4.缓存 go-redis
##### 5.前端框架 vue  

  
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
###### 2.  接口访问根路径：        http://localhost:7001/${Gin.BasePath--为对应各个环境配置的基础地址}/${具体controller路由地址}
