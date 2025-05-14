# XbookLab 前端项目文档

## 技术栈
- 框架：Vue 3（使用`<script setup>`语法糖）
- 路由管理：Vue Router 4
- UI组件库：Ant Design Vue 4
- HTTP客户端：Axios 1
- 日期处理：Day.js 1
- 构建工具：Vite 6
- 样式预处理器：Sass 1
- Mock工具：Mock.js + Vite Plugin Mock

## 功能模块划分
- `src/views/`：核心业务页面（如用户中心、书籍详情页）
- `src/components/`：通用可复用组件（导航栏、卡片组件等）
- `src/layouts/`：页面布局容器（主布局、空布局）
- `src/router/`：路由配置文件
- `src/api/`：接口请求封装
- `src/utils/`：工具函数库（日期格式化、本地存储等）
- `src/assets/`：静态资源（图片、字体）
- `src/style/`：全局样式文件

## 开发规范
模板使用Vue 3 `<script setup>` SFCs，详细语法参考[script setup文档](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup)。

IDE支持可查看[Vue工具链指南](https://vuejs.org/guide/scaling-up/tooling.html#ide-support)。
