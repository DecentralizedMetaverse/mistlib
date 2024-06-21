#!/bin/bash

# ビルド設定
OUTPUT_DIR=bin
ENTRY_POINT=cmd/content/main.go
BINARY_NAME=fw

# 出力ディレクトリを作成
if [ ! -d "$OUTPUT_DIR" ]; then
    mkdir -p "$OUTPUT_DIR"
fi

# Goモジュールの依存関係を整理
echo "Tidy up Go modules..."
go mod tidy

# Goビルドコマンドの実行
echo "Building $BINARY_NAME..."
go build -o "$OUTPUT_DIR/$BINARY_NAME" "$ENTRY_POINT"

# ビルド結果の確認
if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
else
    echo "Build succeeded!"
    echo "Binary is located at $OUTPUT_DIR/$BINARY_NAME"
fi
