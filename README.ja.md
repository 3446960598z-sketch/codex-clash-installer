# Codex + Clash Verge 自動インストーラー

最新の Clash Verge Rev と Codex（ChatGPT デスクトップアプリ）を自動ダウンロード・インストールするクロスプラットフォーム（Windows / macOS）CLI ツールです。

## 主な機能
- リアルタイムで最新バージョンを取得
- 複数ソース + 自動プロキシ検出と速度比較（システムプロキシ、Clash ポート、環境変数）
- マルチスレッドダウンロードと進捗バー表示
- インストーラーを実行ファイルと同じディレクトリに保存
- SHA256 検証
- Windows ガイドフロー：コンソールウィンドウを最前面 → 登録リンクを開く → Clash GUI インストール → Codex MSIX（Add-AppxPackage）→ chatgpt.com を開く → 最終検証 → 自己削除
- macOS 対応（.command ランチャー）

## ビルド
```bash
make
```

## 使い方
対応プラットフォームのバイナリを実行します。Windows ではガイドフローに従い、検証完了後に自身を削除します。
