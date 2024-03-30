### Config 模块

这个模块提供了一个用于读取和监视配置文件的功能。

#### 功能

- 从指定的文件路径读取配置
- 将配置文件的内容解析到指定的对象
- 监视配置文件的更改，并在文件更改时重新加载配置

#### 使用方法

首先，你需要创建一个 `Config` 对象。你可以使用 `New` 函数来创建，这个函数接受一系列的 `Option`，这些 `Option` 可以用来配置 `Config` 对象。
例如：

```golang
cfg, err := config.New(
    config.WithLogger(myLogger),
    config.LoadFile("path/to/my/config.json", &myConfigObject),
)
if err != nil {
    // 处理错误
}
```

然后，你可以使用 Read 方法来读取和解析配置文件：

```golang
err := cfg.Read()
if err != nil {
    // 处理错误
}
```

你也可以使用 `Watch` 方法来监视配置文件的更改：

```golang
err := cfg.Watch()
if err != nil {
    // 处理错误
}
```

#### 注意事项

- `LoadFile` 方法接受一个文件路径和一个目标对象。这个目标对象应该是一个指向你的配置结构体的指针，这个结构体的字段会被用来存储配置文件的内容。
- `WithLogger` 方法接受一个 `log.Logger` 对象，这个对象会被用来记录配置的读取和监视过程中的事件。
- `New` 方法在创建 `Config` 对象时会自动调用 `reflection.SetUp` 方法来初始化目标对象。这个方法会使用 `reflect` 包来遍历目标对象的所有字段，并为每个字段设置一个默认值。这个默认值可以通过字段的 `default` 标签来指定。
- `Read` 和 `Watch` 方法在遇到错误时会返回一个 `errors.Internal` 错误。这个错误包含了一个错误消息和一个错误详情，你可以使用 `errors.Message` 和 `errors.Detail` 函数来获取这些信息。
