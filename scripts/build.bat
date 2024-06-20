@echo off
setlocal

REM ビルド設定
set OUTPUT_DIR=bin
set ENTRY_POINT=cmd\content\main.go
set BINARY_NAME=fw.exe

REM 出力ディレクトリを作成
if not exist %OUTPUT_DIR% (
    mkdir %OUTPUT_DIR%
)

REM Goモジュールの依存関係を整理
echo Tidy up Go modules...
go mod tidy

REM Goビルドコマンドの実行
echo Building %BINARY_NAME%...
go build -o %OUTPUT_DIR%\%BINARY_NAME% %ENTRY_POINT%

REM ビルド結果の確認
if %ERRORLEVEL% neq 0 (
    echo Build failed!
    exit /b %ERRORLEVEL%
) else (
    echo Build succeeded!
    echo Binary is located at %OUTPUT_DIR%\%BINARY_NAME%
)

endlocal
pause
