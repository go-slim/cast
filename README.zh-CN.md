# cast

[English](README.md) | [简体中文](README.zh-CN.md)

一个用于 Go 字符串到原生类型转换的极简、无依赖工具库。

模块路径：`go-slim.dev/cast`

- 将字符串输入转换为 Go 的数值/布尔/字符串类型
- 支持时间和持续时间类型转换
- 提供反射友好的辅助函数来转换为目标 `reflect.Type`
- 通过 `FromType` 提供简单的切片支持（逗号分隔值）

## 安装

```bash
go get go-slim.dev/cast
```

## 快速开始

```go
package main

import (
    "fmt"
    "reflect"
    "time"

    cast "go-slim.dev/cast"
)

func main() {
    // 基本类型转换
    v, _ := cast.FromString("42", "int")
    fmt.Printf("%T %v\n", v, v) // int 42

    v, _ = cast.FromString("true", "bool")
    fmt.Printf("%T %v\n", v, v) // bool true

    // 时间类型转换
    v, _ = cast.FromString("2023-12-25T10:30:45Z", "time.Time")
    fmt.Printf("%T %v\n", v, v) // time.Time 2023-12-25 10:30:45 +0000 UTC

    // 持续时间转换
    v, _ = cast.FromString("1h30m", "time.Duration")
    fmt.Printf("%T %v\n", v, v) // time.Duration 1h30m0s

    // 使用 reflect.Type，包括切片（逗号分隔）
    t := reflect.TypeOf([]int(nil)) // []int
    v, _ = cast.FromType("1,2,3", t)
    fmt.Printf("%T %v\n", v, v) // []int [1 2 3]
}
```

## API

### 核心入口点

- **`FromString(s string, targetType string) (any, error)`**
  - 支持的类型名称：
    - 整数：`int`, `int8`, `int16`, `int32`, `int64`
    - 无符号整数：`uint`, `uint8`, `uint16`, `uint32`, `uint64`
    - 浮点数：`float32`, `float64`
    - 布尔：`bool`
    - 字符串：`string`
    - 时间：`time.Time`, `time.Duration`

- **`FromType(s string, targetType reflect.Type) (any, error)`**
  - 如果 `targetType` 是切片（例如 `[]int`），输入会按逗号 `,` 分割，每个项目会被修剪空白，然后通过 `FromString` 转换为切片元素类型

### 数值辅助函数（字符串 → 类型化值）

- **整数转换**：`Int`, `Int8`, `Int16`, `Int32`, `Int64`
- **无符号整数转换**：`Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
- **浮点数转换**：`Float32`, `Float64`
- **精确小数转换**：`Decimal` (使用 `github.com/shopspring/decimal`)

### 布尔辅助函数

- **`Bool(s string) (bool, error)`**
  - 支持：`true`, `false`, `1`, `0`, `t`, `f`, `TRUE`, `FALSE`, 等

### 时间辅助函数

- **`Time(s string) (time.Time, error)`**
  - 支持多种时间格式：
    - RFC3339：`2023-12-25T10:30:45Z`
    - 日期时间：`2023-12-25 10:30:45`, `2023/12/25 10:30:45`
    - 日期：`2023-12-25`, `12/25/2023`
    - RFC 格式：RFC1123, RFC822
    - Unix 时间戳：`1703505045`
    - 纳秒精度支持

- **`Duration(s string) (time.Duration, error)`**
  - 支持多种格式：
    - 标准 Go 格式：`1h30m45s`, `500ms`, `2h`
    - 整数（纳秒）：`1000000000`
    - 浮点数（秒）：`1.5`, `0.001`

## 实现位置

- `cast/cast.go` - 调度器 `FromString`, `FromType`
- `cast/num.go` - 数值/布尔解析
- `cast/time.go` - 时间/持续时间解析

## 错误行为

- 对于不支持的类型名称，`FromString` 返回：

  ```
  fmt.Errorf("cast: type %v is not supported", targetType)
  ```

- 解析失败时，辅助函数返回来自 `strconv` 的解析错误（`Int` 除外，它会格式化类似 `cast: cannot cast "%v" to type "%v"` 的消息）

- 使用切片类型的 `FromType` 会在遇到第一个项目错误时停止并返回

## 示例

### 基本类型转换

```go
// 无符号整数
v, err := cast.FromString("1", "uint16")
// v == uint16(1), err == nil

// 解析错误
v, err = cast.FromString("str", "int")
// err != nil (无法转换)

// 切片转换
t := reflect.TypeOf([]uint8(nil))
v, err = cast.FromType("1, 2, 255", t)
// v == []uint8{1,2,255}, err == nil
```

### 时间转换

```go
// RFC3339 格式
t, err := cast.Time("2023-12-25T10:30:45Z")
// t == 2023-12-25 10:30:45 +0000 UTC

// 常见日期格式
t, err = cast.Time("2023-12-25")
// t == 2023-12-25 00:00:00 +0000 UTC

// Unix 时间戳
t, err = cast.Time("1703505045")
// t == 2023-12-25 10:30:45 +0000 UTC

// 十六进制整数
v, err := cast.Int("0xFF")
// v == 255

// 八进制
v, err = cast.Int("0755")
// v == 493

// 二进制
v, err = cast.Int("0b1010")
// v == 10
```

### 持续时间转换

```go
// 标准格式
d, err := cast.Duration("1h30m45s")
// d == 1*time.Hour + 30*time.Minute + 45*time.Second

// 浮点秒数
d, err = cast.Duration("1.5")
// d == 1500 * time.Millisecond

// 整数纳秒
d, err = cast.Duration("1000000000")
// d == time.Second
```

### 数值基数支持

```go
// 十六进制
v, err := cast.Int("0xFF")
// v == 255

// 八进制
v, err = cast.Int("0755")
// v == 493

// 二进制
v, err = cast.Int("0b1010")
// v == 10
```

## 测试

运行测试：

```bash
go test -v
```

运行基准测试：

```bash
go test -bench . -benchmem
```

测试套件包括以下覆盖：

- `FromString` 的正常路径和失败情况
- 所有数值宽度和有符号/无符号边界
- 布尔解析（多种格式）
- 时间解析（20+ 种格式）
- 持续时间解析（多种输入格式）
- 带切片目标的 `FromType`
- 边界情况和错误场景
- 性能基准测试

## 版本控制与兼容性

- 推荐 Go 1.20+（已使用 Go 1.24 测试）
- API 小巧且稳定；变更将以向后兼容为目标

## 许可证

MIT
