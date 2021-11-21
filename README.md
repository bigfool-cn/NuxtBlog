# NuxtBlog 极简风格博客
[Nuxt.js + Golang (Gin + Grom + Logrus) 打造的极简风格博客](https://www.bigfool.cn)。

### 效果图
![首页](./1.png)
![后台标签页](./2.png)
![文章页](./3.png)
![文章新增页](./4.png)

### 使用
> 服务端

&nbsp;&nbsp;首先将nuxt-blog-api目录下nuxt-blog.sql文件导入到自己的MySQL数据库，作者用的数据库版本为5.7。
然后需要修改nuxt-blog-api目录下的/configs/config.yaml配置文件的配置值，登录账号密码(username和pwd)、日志存储路径(logpath)、Redsi缓存配置cache和
MySQL数据库配置db。

&nbsp;&nbsp;nuxt-blog-api目录下还提供了Dcokerfile，可以通过该Dockerfile构建项目镜像。

>客户端

```bash
# 开发运行
npm run dev
# 发布打包
npm run build
```
&nbsp;&nbsp;部署在服务器上可以通过pm2跑在宿主机上，也可以通过根目录下的Dockerfile构建项目镜像。

PS:目前在线示例的客户端和服务端均跑在项目提供Dockerfile构建的镜像的容器中。

## License

[MIT](https://github.com/bigfool-cn/NuxtBlog/blob/master/LICENSE)

Copyright (c) 2021-present bigfool-cn
