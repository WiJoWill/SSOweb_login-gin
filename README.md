## Golang编写的sso单点登录系统

#### 项目简介

此项目是一个用go语言编写的sso单点登录系统，此仓库存储用户中心的代码，如需要测试子系统代码请到本文档底部[测试系统](#测试系统)，跳转至对应仓库。

此项目比较粗糙，是一个草稿式的demo，欢迎各路大神增删填补。如果有任何疑问欢迎评论或是联系我（联系方式在文档底部）。

（此仓库不包含整理/修改后保密协议内且不属于个人的代码，请谅解）

***

#### 功能特性

—  简单的注册、登录、和修改密码功能（这里要求用户第一次登录时必须修改原始密码，这个功能当然可以不加上去）

—  Token的生成、解析、验证

—  用户数据表的生成、储存、调取等常规功能

—  首页会显示一个实时更新的时间，首页需Token验证

—  db_info页传递部分数据库用户信息到网页上

—  粗糙的安全功能（CSRF、XSS、数据过滤、SQL注入）

----

#### 环境依赖

golang编程语言

gin框架(https://github.com/gin-gonic/gin)

（以及一些杂七杂八的，源码的import里都有）

----

#### 目录结构描述

##### Controllers

- homepage.go
- register_controller.go
- login_controller.go
- token_controller.go
- database_controller.go
- changepw_controller.go

##### Databases

- mysql.go
- salt.go

##### Models

- redis_model.go
- salt_model.go
- user_model.go

##### Routers

- router.go

##### Statics

- 2view
  - css
    - login.css
  - img
    - check.png
    - checked.png
    - login_background
  - js
    - app.js
    - particles.js
- js
  - lib
    - jQuery.url.js
    - jQuery - 3.3.1.min.js
  - db_info.js
  - login_web.js

##### Utility

- utils.go

##### Views

- Home.html
- login.html
- register.html
- change-password.html
- db_info.html

----

#### 内容更新

- Jul.31, 2019: 第一次上传整个草稿式代码
- Aug.1, 2019: 修改了前端页面显示Json信息，并且添加了部分Readme内容

----

#### 测试系统

这个仓库的代码只涉及到用户登录中心的部分，测试系统如下

[测试子系统A](https://github.com/WiJoWill/SSOweb_login_test_systemA-gin)

----

#### 联系

邮箱：WillWuHJ@outlook.com

[发送邮箱时，请注明来者和来意]


