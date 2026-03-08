# Go アプリケーションテンプレート

Go アプリケーションのテンプレートリポジトリ。

## 依存ツール

| ツール | 用途 | インストール |
|---|---|---|
| [Go](https://go.dev/) | 言語ランタイム | [インストール手順](https://go.dev/doc/install) |
| [golangci-lint](https://golangci-lint.run/) | lint / format チェック | [インストール手順](https://golangci-lint.run/welcome/install/) |
| [lefthook](https://github.com/evilmartians/lefthook) | pre-commit hooks | [インストール手順](https://github.com/evilmartians/lefthook#installation) |

## ディレクトリ構成

```
（後で追加）
```

## 依存の方向

（後で追加）

## セットアップ

```bash
git clone https://github.com/wisteriahuman/go-clean-template.git
cd go-clean-template

# pre-commit hooksの登録
lefthook install

# 動作確認
go build ./...
go test ./...
```

## 開発フロー

```
issue作成 → ブランチ作成 → 実装 → PR → レビュー → merge
```

| ブランチプレフィックス | 用途 |
|---|---|
| `feat/#XX/概要` | feature / task（機能追加） |
| `fix/#XX/概要` | bug（バグ修正） |
| `chore/#XX/概要` | chore（雑務） |
| `docs/#XX/概要` | docs（ドキュメント） |

## CI

PRおよびmainへのpush時に以下が自動実行される。

- `golangci-lint run`（lint / formatチェック）
- `go test ./...`（テスト）

## ラベル

| ラベル | 説明 |
|---|---|
| `epic` | 大きな機能のまとまり |
| `feature` | 実装する機能単位 |
| `task` | featureを分解したもの |
| `chore` | 雑務・設定ファイル修正など |
| `bug` | バグ修正 |
| `arch` | アーキテクチャ設計・レイヤー分離 |
| `api` | APIエンドポイントの実装 |
| `perf` | パフォーマンスチューニング・計測 |
| `test` | テスト・CI |
| `obs` | 可観測性・ロギング |
| `docs` | ドキュメント整備 |
