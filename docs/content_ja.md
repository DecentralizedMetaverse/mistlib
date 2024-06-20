このコードは、特定の「世界」データを操作するためのコマンドラインツールです。以下に各コマンドの使い方を日本語で説明します。

### `fw get-world-info`

現在の「世界」情報を取得し、ローカルファイルに保存します。

**使用例:**
```
fw get-world-info
```

### `fw download-world`

指定したCIDから「世界」データをダウンロードし、ローカルファイルに保存します。

**使用例:**
```
fw download-world <cid>
```
例:
```
fw download-world QmXoYziZsG5r4VbnwZprN9JkQtiMZ3kzgjXo3aE5Dyb53R
```

### `fw get-world-cid`

現在の「世界」データのCIDを取得します。

**使用例:**
```
fw get-world-cid
```

### `fw switch`

新しい「世界」を作成して切り替えます。

**使用例:**
```
fw switch <world-name>
```
例:
```
fw switch new_world
```

### `fw get`

指定したCIDからファイルをダウンロードし、ローカルに保存します。

**使用例:**
```
fw get <cid> <file-name>
```
例:
```
fw get QmXoYziZsG5r4VbnwZprN9JkQtiMZ3kzgjXo3aE5Dyb53R example.txt
```

### `fw put`

指定したファイルを指定の座標で「世界」に追加します。

**使用例:**
```
fw put <file> <x> <y> <z> <rx> <ry> <rz> <sx> <sy> <sz>
```
例:
```
fw put example.txt 1.0 2.0 3.0 0.0 0.0 0.0 1.0 1.0 1.0
```

### `fw update`

既存のメタデータを更新します。

**使用例:**
```
fw update <meta-cid> <x> <y> <z> <rx> <ry> <rz> <sx> <sy> <sz>
```
例:
```
fw update QmXoYziZsG5r4VbnwZprN9JkQtiMZ3kzgjXo3aE5Dyb53R 1.0 2.0 3.0 0.0 0.0 0.0 1.0 1.0 1.0
```

### `fw set-custom-data`

指定したメタデータにカスタムデータを設定します。

**使用例:**
```
fw set-custom-data <meta-cid> <value>
```
例:
```
fw set-custom-data QmXoYziZsG5r4VbnwZprN9JkQtiMZ3kzgjXo3aE5Dyb53R custom_value
```

### `fw set-parent`

指定したメタデータを親子関係として設定します。

**使用例:**
```
fw set-parent <child-cid> <parent-cid>
```
例:
```
fw set-parent QmXoYziZsG5r4VbnwZprN9JkQtiMZ3kzgjXo3aE5Dyb53R QmPARENTxG5r4VbnwZprN9JkQtiMZ3kzgjXo3aE5Dyb53R
```

### `fw set-password`

パスワードを設定します。

**使用例:**
```
fw set-password <password>
```
例:
```
fw set-password my_secret_password
```

### `fw init`

新しい「世界」リポジトリを初期化します。

**使用例:**
```
fw init
```

### `fw cat`

指定したパスのファイルの内容を表示します。

**使用例:**
```
fw cat <path>
```
例:
```
fw cat QmXoYziZsG5r4VbnwZprN9JkQtiMZ3kzgjXo3aE5Dyb53R
```

これらのコマンドを使って、「世界」データの操作や管理を行うことができます。