# 火币(Huobi)价格监控

由于部分交易对火币官方未提供价格监控，因此写了个小程序，长期屯币党可以用它来提醒各种现货价格。

该工具只需要提前安装Go环境和Redis即可。

消息推送使用的「钉钉」，需要提前配置好钉钉机器人（企业群类型、带webhook的机器人）。

## 使用方法

1. 下载本项目
2. 拷贝根目录下 `.env.sample` 文件至 `.env`，完成`.env`文件的配置
3. 执行 `go mod download` 和 `go build .` 即可获取该工具binary文件
4. 可以用supervisor之类来管理常驻后台进程

## 效果截图
[![4BqGgs.md.jpg](https://z3.ax1x.com/2021/09/24/4BqGgs.md.jpg)](https://imgtu.com/i/4BqGgs)
[![4Bq3CQ.md.jpg](https://z3.ax1x.com/2021/09/24/4Bq3CQ.md.jpg)](https://imgtu.com/i/4Bq3CQ)
[![4Bq83j.md.jpg](https://z3.ax1x.com/2021/09/24/4Bq83j.md.jpg)](https://imgtu.com/i/4Bq83j)
