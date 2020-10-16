# go-frpc
frp 客户端可视化（目前只支持windows系统）

<h2>源码下载后打包</h2>
<h4>Windows 打包命令</h4>
```fyne package -os windows -icon icon.png```
<h4>Linux 打包命令</h4>
```fyne package -os linux -icon icon.png```
<h4>Macos 打包命令</h4>
```fyne package -os darwin -icon icon.png```

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

2020-8-31  v0.1基本使用