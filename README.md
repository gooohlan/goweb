# goweb
封装gin完成最基本的web功能，自定义错误码及返回格式

# 目录结构

- controller 用于接收和返回前端传递及所需参数
- db 数据库相关操作内容
- errors 错误信息处理（日志写入待做）
- helper 公共函数
- model 数据模型
- router 路由
- service 业务逻辑相关

# 验证器
- 自定义验证规则在`vali/validator`下,参考checkMobile
- controller下使用验证器,调用`api.API.ParseRequest`


# 自定义返回码

在errors -> code -> table 中 设置对应Code和Msg

# 日志处理
待做中。。。
