# Codex + Clash Verge 自动化安装工具

跨平台（Windows / macOS）自动下载并安装最新版 Clash Verge Rev 和 Codex（ChatGPT 桌面应用）的 CLI 工具。

## 主要功能
- 实时获取最新版本
- 多源 + 自动代理测速优选（系统代理、Clash 常见端口、环境变量）
- 多线程下载，显示总进度条与分段进度
- 安装包下载到程序当前目录
- SHA256 校验
- Windows 引导流程：窗口置顶 → 打开注册链接 → Clash GUI 安装 → Codex MSIX（Add-AppxPackage）→ 打开 chatgpt.com → 最终校验 → 自行销毁
- macOS 支持双击启动脚本

## 编译
```bash
make
```

## 使用
运行对应平台的二进制文件。Windows 下工具会按引导流程执行，完成校验后自动删除自身。
