4-9
登录、注册页面的制作
对于用户名、电子邮箱，密码的验证，使用数据库进行存储（这里采用sqlite3）


数据库目前使用orm进行管理
表结构为
type Userinfo struct{
    UserName string `orm:PK` //显示说明用户名为主键
    Email string //电子邮箱
    Password string //密码
}