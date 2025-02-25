# `glog` 包使用说明文档

## 一、概述

`glog` 是一个用于日志记录和上传的 Go 语言包，支持将日志信息打印到控制台，并可选择将日志上传到指定的 OpenObserve 服务。它提供了不同级别的日志记录方法，如 `Info`、`Debug`、`Warn` 和 `Error`，方便在不同场景下使用。

## 二、安装

### 前提条件

确保你已经安装了 Go 语言环境，版本要求 Go 1.16 及以上。

### 安装步骤

在你的 Go 项目中，使用以下命令安装 `glog` 包：

```sh
go get github.com/genez233/glog
```

## 三、使用方法

### 1. 导入包

在你的 Go 代码中导入 `glog` 包：

```go
import (
    "github.com/genez233/glog"
)
```

### 2. 配置 `glog`

在使用 `glog` 之前，需要进行配置。配置信息包含应用名称、版本、是否打印到控制台、是否上传日志、运行环境、上传地址和 OpenObserve 的 Token 等。示例代码如下：

```go
config := &glog.Config{
    ServerName:       "your_app_name",
    Version:          "1.0.0",
    ConsoleLog:       true,
    IsUpload:         true,
    RunMode:          "development",
    LogUrl:           "http://your.domain.com/api/default/%s/_json",
    OpenobserveToken: "your_openobserve_token",
}
```

### 3. 创建 `glog` 实例

使用配置信息创建 `glog` 实例：

```go
logger := glog.New(config)
```

### 4. 记录日志

`glog` 提供了不同级别的日志记录方法，使用方式如下：

```go
// 记录 Info 级别的日志
logger.Info("This is an info message")

// 记录 Debug 级别的日志
logger.Debug("This is a debug message")

// 记录 Warn 级别的日志
logger.Warn("This is a warning message")

// 记录 Error 级别的日志
logger.Error("This is an error message")
```

## 四、配置说明

`glog` 的配置信息通过 `Config` 结构体进行设置，各字段说明如下：

| 字段名             | 类型   | 说明                                                         |
| ------------------ | ------ | ------------------------------------------------------------ |
| `ServerName`       | string | 应用名称，对应 OpenObserve 仓库名，默认为 "default"          |
| `Version`          | string | 应用版本                                                     |
| `ConsoleLog`       | bool   | 是否将日志打印到控制台，`true` 表示打印，`false` 表示不打印  |
| `IsUpload`         | bool   | 是否将日志上传到 OpenObserve 服务，`true` 表示上传，`false` 表示不上传 |
| `RunMode`          | string | 运行环境，如 "development"、"production" 等                  |
| `LogUrl`           | string | 日志上传地址，支持格式化字符串，使用 `%s` 占位应用名称       |
| `OpenobserveToken` | string | OpenObserve 的 Token，用于身份验证                           |

## 五、注意事项

1. **错误处理**：在 `post` 方法中，对于 HTTP 请求和响应的错误处理只是简单地打印错误信息，实际使用中可以根据需求进行更详细的错误处理。
2. **日志上传**：如果 `IsUpload` 为 `true`，需要确保 `LogUrl` 和 `OpenobserveToken` 配置正确，否则日志上传可能会失败。
3. **性能考虑**：频繁的日志上传可能会影响性能，可根据实际情况调整上传频率或采用异步上传的方式。
