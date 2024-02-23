### 日志模块

这个日志模块提供了一种灵活的方式来记录应用程序的运行信息。它支持多种日志级别，可以根据需要过滤和格式化日志消息。

#### 主要功能

- 支持多种日志级别：`Debug`，`Info`，`Warn`，`Error`，`Fatal`。
- 提供了全局日志函数，可以在任何地方方便地记录日志。
- 支持日志过滤，可以根据日志级别，键或值来过滤日志。
- 提供了格式化日志的函数，可以按照特定的格式输出日志。

#### 使用方法

##### 设置日志级别

使用 `FilterLevel` 函数来设置日志级别。例如，如果你只想记录 `Warn` 级别以上的日志，你可以这样设置：

```golang
log.FilterLevel(log.LevelWarn)
```

##### 记录日志

使用 `Debug`，`Info`，`Warn`，`Error`，`Fatal` 函数来记录日志。例如：

```golang
log.Debug("This is a debug message")
log.Info("This is an info message")
log.Warn("This is a warning message")
log.Error("This is an error message")
```

##### 格式化日志

使用 `Debugf`，`Infof`，`Warnf`，`Errorf`，`Fatalf` 函数来格式化日志。例如：

```golang
log.Debugf("This is a %s message", "debug")
log.Infof("This is a %s message", "info")
log.Warnf("This is a %s message", "warning")
log.Errorf("This is a %s message", "error")
```

##### 格式化 KV 键值对日志

使用 `Debugw`, `Infow`，`Warnw`，`Errorw`，`Fatalw` 函数来格式化日志。例如：

```golang
log.Debugw("This is a %s message", "debug", "error", "This is a error")
log.Infow("This is a %s message", "info", "error", "This is a error")
log.Warnw("This is a %s message", "warning", "error", "This is a error")
log.Errorw("This is a %s message", "error", "error", "This is a error")
```

##### 过滤日志

使用 `FilterKey` 和 `FilterValue` 函数来过滤日志。例如，如果你想过滤键为 "password" 的日志，你可以这样设置：

```golang
log.FilterKey("password")
```

如果你想过滤值为 "error" 的日志，你可以这样设置：

```golang
log.FilterValue("error")
```
