# Codex + Clash Verge 自动安装器

跨平台（Windows & macOS）命令行工具，用于自动下载并安装 **Clash Verge Rev** 和 **Codex / ChatGPT 桌面应用**。

## 主要功能
- 动态获取最新版本（GitHub API / CDN）
- 多源测速 + 自动检测代理（系统代理、Clash 常见端口等）
- 多线程下载 + 总进度条与分段指示
- 下载到程序当前目录
- SHA256 校验
- Windows 引导流程：窗口置顶 → 打开注册链接 → 下载 Clash → GUI 安装 → 下载 Codex MSIX → Add-AppxPackage → 打开 chatgpt.com → 最终校验 → 自行销毁
- macOS 支持（可使用 .command 启动器）

## 使用方法
直接运行对应平台的二进制文件。Windows 下会自动执行上述引导流程。

## 编译
需要 Go 1.21+。

```bash
make all
```

编译产物位于 `dist/` 目录。

## 免责声明
本工具仅自动化下载与安装公开可用的官方或镜像包。用户需自行遵守当地法律，并配置自己的代理订阅。