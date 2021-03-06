# go-frpc
frp 客户端可视化（目前只支持windows、Linux系统）

<h2>源码下载后打包</h2>
<h4>Windows 打包命令</h4>
`fyne package -os windows -icon icon.png`
<h4>Linux 打包命令</h4>
`fyne package -os linux -icon icon.png`
<h4>Macos 打包命令</h4>
`fyne package -os darwin -icon icon.png`

<h5>注意：
在Linux平台开发打包时需要注释 src/utils/cmd/windows/windows.go 文件中 `RunCommandBg` 方法中的行 `command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}`<br/> 
在windows打包时必须打开</h5>


<h2>使用说明</h2>

1. 下载frp

![image](https://github.com/luoqiz/go-frpc/raw/master/readme/1.jpg)

2. 修改配置文件，在这里需要添加所有要用到节点，也可以在第三步中单独修改每个节点（前提是此处添加了节点）

![image](https://github.com/luoqiz/go-frpc/raw/master/readme/2.jpg)

3. 修改每个节点内容

![image](https://github.com/luoqiz/go-frpc/raw/master/readme/3.jpg)

4. 启动服务或停止服务

![image](https://github.com/luoqiz/go-frpc/raw/master/readme/4.jpg)


<h2>更新日志：</h2>

- 2020-11-19 v0.2
    1. 更新依赖组件、调整ui
    2. 提取 `frpc.ini` 文件到根目录，可更新 `frp` 软件
    3. 增加 http 代理，加快 github 下载速度
    4. 增加 linux 平台使用
    
- 2020-8-31  v0.1基本使用

## 致谢

- (https://www.jetbrains.com/?from=mogu_blog_v2)感谢 **[jetbrains](https://www.jetbrains.com/?from=mogu_blog_v2)** 提供的开源License 
