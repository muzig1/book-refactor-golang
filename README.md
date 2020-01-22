<p align="center">
    <img src="/go&#32;runner.jpg" height="400"></img>
    <h1 align="center">Refactor with Go</h1>
</p>

将[《重构》](https://baike.baidu.com/item/%E9%87%8D%E6%9E%84/2182519?fr=aladdin)一书中的Java 示例修改为Golang 示例

## 下载

### 代码

```bash
git clone --depth=1 https://github.com/nxlsBoy/book-refactor-golang.git
```

### PDF

请支持正版图书 - [下载链接](https://github.com/nxlsBoy/book-refactor-golang/raw/master/%E9%87%8D%E6%9E%84(%E7%AC%AC%E4%B8%80%E7%89%88).pdf)

## 目录

- [第八章-重新组织数据](#第八章-重新组织数据)
- [第十章-简化函数调用](#第十章-简化函数调用)

## 第八章-重新组织数据

从数据结构层面去重构代码

| Patern                                                                              | Expression                                                       | Progress |
| :---------------------------------------------------------------------------------- | :--------------------------------------------------------------- | :------- |
| [Self Encapsulate Field](./chapter8/self_encapsulate_field/README.md)               | 自我封装 -<br>将直接访问成员字段的方式修改为间接访问             | ✔        |
| [Replace Data Value with Object](chapter8/replace_data_value_with_object/README.md) | 结构对象替换简单类型 -<br> 将简单数据类型替换为结构类型,易于扩展 | ✔        |
| [Change Value to Reference](chapter8/change_value_to_reference/README.md)           | 值类型替换为引用类型 -<br> 修改一处, 即可修改其他所有引用的地方  | ✔        |
| Change Reference to Value                                                           |                                                                  | ✘        |
| Replace Array with Object                                                           |                                                                  | ✘        |
| Duplicate Observed Data                                                             |                                                                  | ✘        |
| Change Unidirectional Association to Bidirectional                                  |                                                                  | ✘        |
| Change Bidirectional Association to Unidirectional                                  |                                                                  | ✘        |
| Replace Magic Number with Symbolic Constant                                         |                                                                  | ✘        |
| Encapsulate Field                                                                   |                                                                  | ✘        |
| Encapsulate Collection                                                              |                                                                  | ✘        |
| Replace Record with Data Class                                                      |                                                                  | ✘        |
| Replace Type Code with Class                                                        |                                                                  | ✘        |
| Replace Type Code with Subclasses                                                   |                                                                  | ✘        |
| Replace Type Code with State/Strategy                                               |                                                                  | ✘        |
| Replace Subclass with Fields                                                        |                                                                  | ✘        |

## 第十章-简化函数调用

针对函数调用的优化

| Patern                                             | Expression           | Prograss |
| :------------------------------------------------- | :------------------- | :------- |
| [Rename Method](chapter10/rename_method/README.md) | 修改函数名           | ✔        |
| Add Parameter                                      | 添加参数             | ✘        |
| Remove Parameter                                   | 移除参数             | ✘        |
| Separate Query from Modifier                       | 查询与修改函数独立   | ✘        |
| Parameterize Method                                | 让函数携带参数       | ✘        |
| Replace Parameter with Explicit Methods            | 以明确函数取代参数   | ✘        |
| Preserve Whole Object                              | 保持对象完整         | ✘        |
| Replace Parameter with Methods                     | 以函数取代参数       | ✘        |
| Introduce Paramter Object                          | 引入参数对象         | ✘        |
| Remove Setting Method                              | 溢出设置值函数       | ✘        |
| Hide Method                                        | 隐藏函数             | ✘        |
| Replace Constructor with Factory Method            | 工厂函数取代构造函数 | ✘        |
| Encapsulate Downcast                               | 封装向下转型         | ✘        |
| Replace Error Code with Exception                  | 以异常取代错误骂     | ✘        |
| Replace Exception with Test                        | 以测试取代异常       | ✘        |
