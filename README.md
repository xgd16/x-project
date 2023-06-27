# xInit 项目自动初始化

### 打包
```shell
make build-pack
```
### 创建指令 (在需要创建项目的目录运行)
```shell
./xInit_darwin_arm64 -p testDemo
```
##### 指令解析
``./xInit_darwin_arm64`` 下载的执行文件名称

``-p`` 执行项目名称 请使用 ``golang`` 支持的 ``mod`` 名称标准 *PS: 不指定 -p 默认使用* ``x-project``

``testDemo`` 生成出的项目名称