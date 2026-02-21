.PHONY: build test lint fmt clean help

# デフォルトターゲット
all: fmt lint test build

# ビルド
build:
	go build ./...

# テスト
test:
	go test ./... -v

# テスト（カバレッジ付き）
test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# フォーマット
fmt:
	go fmt ./...

# リント（staticcheckがインストールされている場合）
lint:
	@which staticcheck > /dev/null 2>&1 && staticcheck ./... || echo "staticcheck not installed, skipping"

# 依存関係の整理
tidy:
	go mod tidy

# クリーンアップ
clean:
	rm -f coverage.out coverage.html
	go clean

# サンプル実行
run-example:
	go run ./examples/basic/main.go

# ヘルプ
help:
	@echo "使用可能なターゲット:"
	@echo "  make build         - パッケージをビルド"
	@echo "  make test          - テストを実行"
	@echo "  make test-coverage - カバレッジ付きテストを実行"
	@echo "  make fmt           - コードをフォーマット"
	@echo "  make lint          - 静的解析を実行"
	@echo "  make tidy          - go mod tidy を実行"
	@echo "  make clean         - 生成ファイルを削除"
	@echo "  make run-example   - サンプルを実行"
	@echo "  make all           - fmt, lint, test, build を実行"
