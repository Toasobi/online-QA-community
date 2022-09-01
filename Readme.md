这是最后考核作业（单体架构实现）
在基础内容上多加了redis缓存，点赞，简单关注，邮箱发送验证码注册功能。（虽然有个dockerfile但其实docker部署没部署成功。。。）

swag 接口访问地址:http://localhost:8080/swagger/index.html

生成token：
安装jwt ： go get github.com/dgrijalva/jwt-go

通过邮箱发送验证码功能
使用网上现成的包 ： github.com/jordan-wright/email
安装: go get github.com/jordan-wright/email
