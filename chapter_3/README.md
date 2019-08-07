# 代码的坏味道

注意：当看到第三章的第四个点内容的时候，发现以下重构手法，似乎在后续有具体的描述，所以前三个内容，均为自己的感悟和理解，酌情理解和使用。

后续此章节内容，小生只会更新README.md 内容，类似整理一个手法列表吧

- Duplicated Code - 重复代码
    - Extract Method - 提炼函数
    - Pull Up Method - 提升函数权限
    - Extract Class - 提炼到类内部
- Long Method - 超长函数
    - Extract Method - 提炼函数
    - Replace Temp with Query
    - Introduce Parameter Object | Preserve Whole Object
    - Replace Method with Method Object
    - Decompose Conditional
- Large Class - 过大的类
    - Extract Method
    - Extract Class | Extract Subclass
- Long Parameter List - 过长的参数列表
    - Replace Parameter with Method
    - Preserve Whole Object
    - Introduce Parameter Object - 参数对象
