# 简介
TP-GEN 是一个生成TP框架下的 `CURD` 的工具,可以在项目目录下生成对应的 `controller`, `models`, `service`

# 使用方式
`.\curd.exe -t versions -g aduli -m api`
- `-g` 生成控制方法
    - `a` 生成控制器创建操作
    - `d` 生成控制器删除操作
    - `u` 生成控制器更新操作
    - `l` 生成控制器生成列表操作
    - `i` 生成控制器获取详情操作
- `-m` 指定生成文件在哪个 `module` 下面(不指定使用配置文件中 `module` 配置)


# 配置说明
```
path: 项目路径,不配置则默认为当前目录
application_name: app目录名称
module: 默认生成在哪个模块下
```