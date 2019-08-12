Golang编写的sso单点登录系统

#### 项目简介

此项目是一个用go语言编写的sso单点登录系统，此仓库存储用户中心的代码，如需要测试子系统代码请到本文档底部[测试系统](#测试系统)，跳转至对应仓库。

此项目比较粗糙，是一个草稿式的demo，欢迎各路大神增删填补。如果有任何疑问欢迎评论或是联系我（联系方式在文档底部）。

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

  - 「测试用的Homepage，无实际功能」

- register_controller.go 「注释已详细解释代码」

  - 获取页面

    ```go
    func RegisterGet (c *gin.Context){}
    ```

  - 注册功能

  - ```go
    func RegisterPost (c *gin.Context){}
    ```

- login_controller.go 「注释已详细解释代码」

  - 获取页面

  - ```go
    func LoginGet (c *gin.Context){}
    ```

  - 登录功能

  - ``` go
    func LoginPost (c *gin.Context)
    ```

  - 生成token

  - ```go
    func generateToken(c *gin.Context, user model.User) string {}
    ```

  - 生成一个子token，逻辑与上一个方法相同

  - ```go
    func (j *JWT) CreateTokenSub (claims UserClaims) (string, error){}
    ```

  - 登录结果

    ```go
    type LoginResult struct {}
    ```

- token_controller.go

  - 常数，用来作为登录信息「查看源码和注释」

  - Userclaims

    ```go
    type UserClaims struct{}
    ```

  - 新建jwt实体

  - ``` go
    func NewJWT() *JWT{}
    ```

  - 获取SignKey

  - ```go
    func GetSignKey() string{}
    ```

  - 设置SignKey

  - ```go
    func SetSignKey() string{}
    ```

  - JWT中间件，用来检验token信息，只服务于用户认证中心，子系统有类似功能但应该根据子系统目的进行增删查改

  - ```go
    func (j *JWT) CreateToken(claims UserClaims) (string, error) {}
    ```

  - 生成Token

  - ```go
    func (j *JWT) CreateTokenSub (claims UserClaims) (string, error){}
    ```

  - 给子系统生成token，可不启用

  - ```go
    func (j *JWT) CreateTokenSub (claims UserClaims) (string, error){}
    ```

  - 解析token，用来判断token

  - ```go
    func (j *JWT) ParseToken(tokenString string) (*UserClaims, error) {}
    ```

  - 更新token，一般不启用

  - ```go
    func (j *JWT) RefreshToken(tokenString string) (string, error) {}
    ```

  - 验证用户的token信息，要使用redis

  - ```go
    func CheckUserToken (token string) bool{}
    ```

  - 从redis库通过读取token获取用户的信息

  - ```go
    func GetUserFromToken (token string) bool{}
    ```

- database_controller.go

  - 显示页面，Get方法
  - 显示前五个用户的基础数据，Post方法
  - 请求用户信息，具有验证IP的功能，如果IP不一致则让用户重新登录（这里未实施，如果有需要请自行添加Redirect）

- changepw_controller.go

  - 显示页面，Get方法

  - ```go
    func ChangepwGet (C *gin.Context){}
    ```

  - 修改密码，Post方法

  - ```go
    func ChangepwPost (c *gin.Context){}
    ```

##### Databases

- mysql.go

  - 启动Mysql，并且验证是否有用户表格

  - ```go
    func InitMysql(){}
    ```

  - 创建用户表

  - ```go
    func CreateTableWithUser(){}
    ```

  - 执行数据库，使用db.Exec命令

  - ```go
    func ModifyDB(sql string, args ... interface{})(int64, error){}
    ```

  - 查询数据库

  - ```go
    func QueryRowDB(sql string) *sql.Row{}
    ```

- salt.go

  - 创建盐表

  - ```go
    func ModifyDBSalt(sql string, args ...interface{})(int64, error){}
    ```

##### Models

- redis_model.go

  - 连接对应的redis

  - ```go
    func ConnectRedis(){}
    ```

  - 存储token，key为token， value为username，存储于db0

  - ```go
    func SetToken(token string, username string){}
    ```

  - 存储用户登录的ip，ip为token，value为token，存储于db2

  - ```go
    func SetTokenIP(ip string, token string){}
    ```

  - 查询ip地址是否与对应的token相符合，如果不是则说明用户不于登录时的ip访问网站，则不允许访问

  - ```go
    func CheckIPAndToken(ip string, token string) bool {}
    ```

  - 检查密钥是否存在于redis库中

  - ```go
    func CheckToken(token string) bool{}
    ```

  - 从redis库中获取对应token的username

  - ```go
    func GetTokenValue (token string) string{}
    ```

- salt_model.go

  - usersalt的structure，包含id, username, saltstring

  - ```go
    type UserSalt struct{}
    ```

  - 生成新的用户盐，并且插入到表中

  - ```go
    func InsertUserSalt (salt UserSalt) (int64, error){}
    ```

  - 更新用户盐信息（但感觉没有什么用，除非后期提供修改用户名的功能）

  - ```go
    func UpdateUserSalt (salt UserSalt)(int64, error){}
    ```

  - 根据用户的id返回对应的用户盐

  - ```go
    func QueryUserSaltWithID(id int)string {}
    ```

  - 根据用户名返回对应的用户盐

  - ```go
    func QueryUserSaltWithUsername(username string) string{}
    ```

- user_model.go

  - 这是一个user的structure，包含id，username，password和status

  - ```go
    type User struct{}
    ```

  - 生成新的用户并插入至用户表

  - ```go
    func InsertUser(user User)(int64, error){}
    ```

  - 修改用户的资料信息，更新相关内容

  - ``` go
    func UpdateUser(user User)(int64 error){}
    ```

  - 按条件查询对应用户，需要输入对应的sql

  - ```go
    func QueryUserWightCon(con string) int{}
    ```

  - 按条件查询用户状态，需要输入对应的sql

  - ```go
    func QueryUserStatusWightCon(con string) int{}
    ```

  - 根据用户名查询用户状态

  - ```go
    func QueryUserStatusWithusername(username string) int{}
    ```

  - 根据用户名查询用户id

  - ```go
    func QueryUserWithusername(username string) int{}
    ```

  - 根据用户名和密码，查询id，一般用来作为用户登录的信息检验

  - ```go
    func QueryUserWithParam(username, password string) int{}
    ```

  - 根据用户id查询用户状态

  - ```go
    func QueryUserInfoWithID(id int) string {}
    ```

##### Routers

- router.go

  - 启动路由，每个部分有对应注释

  - ```go
    func InitRouter() *gin.Engine{}
    ```

##### Statics

- 2view
  - css
    - login.css
      - 登录页面的css文件「请看注释」
  - img - 服务登录页面的图片
    - check.png
    - checked.png
    - login_background.png
  - js - 给动态页面服务的js
    - app.js
    - particles.js
- js
  - lib - 一些js源文件
    - jQuery.url.js
    - jQuery - 3.3.1.min.js
  - db_info.js
    - 给db_info页提供功能
  - login_web.js
    - 给login_web提供登录和注册功能

##### Utility

- utils.go

  - MD5加密功能，注释里提供了加盐的方式

  - ```go
    func MD5 (str string) string{}
    ```

##### Views

- Home.html
  - 「请直接读代码和注释」
- login.html
  - 「请直接读代码和注释」
- register.html
  - 「请直接读代码和注释」
- change-password.html
  - 「请直接读代码和注释」
- db_info.html
  - 「请直接读代码和注释」

----

#### 内容更新

- Jul.31, 2019: 第一次上传整个草稿式代码
- Aug.1, 2019: 修改了前端页面显示Json信息，并且添加了部分Readme内容
- Aug.5, 2019: 换了个好看点的ui，并且添加了README的内容，并且提供了更加安全的用户信息保存方法（在注释里，未启用）。

----

#### 项目实例

（扔到了云服务器）

[Example](http://139.155.74.24:8081/)

----

#### 测试系统

这个仓库的代码只涉及到用户登录中心的部分，测试系统如下

[测试子系统A](https://github.com/WiJoWill/SSOweb_login_test_systemA-gin)

----

#### 联系

邮箱：WillWuHJ@outlook.com

[发送邮箱时，请注明来者和来意]


