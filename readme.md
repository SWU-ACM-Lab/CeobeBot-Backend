<div align="center"><img src="https://z3.ax1x.com/2021/10/14/5Q1m3n.png"/></div>

----

# 1. 说明

这里是CeobeBot的后端，目录结构如下：

```text
.
├── api
├── bin
├── controller
├── module
└── util
```

+ api是封装Http(s)Api的目录，其中v1提供简易版HttpApi服务，v2提供RESTful风格API服务。

# 2. 文档

请转到[CeobeBot-Document](https://github.com/SWU-ACM-Lab/CeobeBot-Document)项目。

# 3. 使用与部署

## 3.1 使用

如果您已经部署好了前端，不想自己托管数据，您可以直接使用我们的后端，具体使用方法详见上文的文档仓库。

但是在开始前，您需要向我们申请后端服务使用授权，我们需要您提供以下资料：

1. 您的联系方式
2. 您的用户规模
3. 您的机器人的NickName

您需要将上述信息发送至[support@mail.swu-acm.cn](mailto:support@mail.swu-acm.cn)，同时您需要确保您使用我们后端服务时不触犯中华人民共和国法律，同时不侵犯他人的权益，并承担您行为所带来的可能的后果。

## 3.2 部署

1. 克隆本项目到本地

> 请在合适的目录下执行

```shell
git clone https://github.com/SWU-ACM-Lab/CeobeBot-Backend.git
```

2. 自动安装依赖

> 请在本项目目录下执行

```shell
go mod tidy
```

3. 修改配置文件

> 请在本项目目录下执行

```shell
cp config.example.ini config.ini
```

将`config.ini`中的各项内容修改为您的实际配置

4. 构建并运行可执行文件

> 请在本项目目录下执行

```shell
go run main.go
```

# 4. 支持

如果你部署有困难，可以支付一定的报酬给`SWU-ACM-Lab`，我们会协助您进行部署，您的资助将全部用于本项目的维护费用。

您也可以按照issue模版提交您的疑惑，我们也将尽力解决您的问题。

# 5. 贡献

您可以fork本项目，然后通过提交pr来贡献代码。也可以提交issue来提出您的意见与建议。

# 鸣谢

感谢下列人员/组织以及他们为本项目做出的贡献。

| NickName | Contribution |
| :==: | :==: |
| [SunistC](https://github.com/sunist-c) | 后端构建者 |
| [Jecosine](https://github.com/Jecosine) | 前端构建者与数据库支持 |
| [pianfanshaan](https://github.com/pianfanshaan) | 项目测试与数据库支持 |
| [西南大学ACM实验室](https://github.com/SWU-ACM-Lab) | 托管服务器支持 |