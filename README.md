# Sasa:tanabata_tree:

> Sasa is an Automatic Scoring Application for "PandA1".

Sasa は PandA1 の課題の採点を自動化するための CLI アプリケーションです。

## Concepts

Sasa は以下の作業を自動化します。

- ソースファイルの UTF-8 へのエンコード
- エンコードされたソースファイルのコンパイル
- バイナリファイルの実行
- 実行結果の採点

Sasa はコマンド一つで採点を終わらせます。

## Requirements

`gcc`あるいは`clang`をインストール済みであること

## Installation

[ここ](https://github.com/tighug/sasa/releases)からバイナリファイル`sasa`をダウンロードできます。

## Quick Start

1. 作業ディレクトリを新規作成（例 : `workspace/`）
2. 作業ディレクリにダウンロードした`sasa`を配置
3. ターミナル上で作業ディレクリに移動

```bash
$ cd workspace
$ ls
sasa
```

4. ターミナル上で`./sasa init`を実行
   （`src/`、`answer.txt`、`.sasarc.yaml`が自動で生成される）

```bash
$ ./sasa init
$ ls
src/  sasa  answer.txt  .sasarc.yaml
```

5. 模範解答プログラムを用意して実行（例 : `hello_ans.c`）

```bash
$ gcc hello_ans.c -o hello_ans
$ ./hello_ans
hello, world
```

6. 上記 5 の出力結果（`hello, world`）を`answer.txt`にコピー&ペースト
7. LETUS から課題をダウンロード
   （"提出をフォルダに入れてダウンロードする"のチェックは外す）

<div align="center">
   <img src="https://github.com/tighug/sasa/blob/asset/image/letus_options.png?raw=true" width=400px />
</div>

8. ダウンロードした Zip を展開し、中のソースファイルを`src/`内に移動

<div align="center">
   <img src="https://github.com/tighug/sasa/blob/asset/image/source_files.png?raw=true" width=300px />
</div>

9. ターミナル上で`./sasa all`を実行
   （`encoded/`、`build/`、`output/`、`database.csv`が自動で生成される）

```bash
$ ./sasa all
[1/4] Encoding...
[2/4] Building...
[3/4] Running...
[4/4] Checking...
$ ls
src/  encoded/  build/  output/
sasa  answer.txt  .sasarc.yaml  database.csv
```

10. 採点結果が記録された`database.csv`を見ながら、LETUS 上で評点を行う
    (Score が低いプログラムは、`encoded/`内のファイルを個別で確認する)

## Commands

<div align="center">
   <img src="https://github.com/tighug/sasa/blob/asset/image/sasa_cmd.jpg?raw=true" width=400px/>
</div>

### `help`

コマンドのヘルプを表示します。

### `version`

Sasa のバージョンを表示します。

### `init`

作業ディレクトリの初期化コマンドです。

`src/`、`answer.txt`、`.sasarc.yaml`を自動生成します。既に上記のディレクトリ・ファイルが存在する場合には、それを上書きせずに無視します。

### `encode`

ソースファイルのエンコードコマンドです。現在、Shift JIS / UTF-8 に対応しています。

`src/`のソースファイルをエンコードし、`encoded/`に出力します。UTF-8 はそのまま、UTF-8 以外のものは UTF-8 にエンコードして出力します。出力ファイル名は`[学籍番号]_[名前].c`です。

`src/`が存在しない場合には、エラーになります。

### `build`

ソースファイルのコンパイルコマンドです。

`encoded/`のソースファイルをコンパイルし、`build/`に出力します。コンパイル時に警告・エラーを吐く場合、ログファイルを出力します。エラーを吐く場合、バイナリファイルは出力されません。バイナリファイル名は`[学籍番号]_[名前]`、ログファイル名は`[学籍番号]_[名前].log`です。

`encoded/`が存在しない場合には、エラーになります。

### `run`

バイナリファイルの実行コマンドです。

`build/`のバイナリファイルを実行し、標準出力をログファイルとして`output/`に出力します。ログファイル名は`[学籍番号]_[名前].log`です。

`build/`が存在しない場合には、エラーになります。`input.txt`が存在する場合には、その中身を標準入力に使用します。

### `check`

出力結果の採点コマンドです。

`output/`のログファイルと`answer.txt`の内容を比較して採点します。同じ行 1 つにつき、1 点を加点します。採点結果は`database.csv`に記録されます。

`output/`や`answer.txt`が存在しない場合には、エラーになります。

### `all`

上記`encode`/`build`/`run`/`check`の 4 つを一括で行うコマンドです。

## Directories & Files

### `src/`

LETUS からダウンロードしたソースファイルを配置するディレクトリです。

コマンド`init`によって自動生成されます。

### `encoded/`

UTF-8 にエンコードされたソースファイルが出力されるディレクトリです。

コマンド`encode`によって自動生成されます。

### `build/`

コンパイルされたバイナリファイルが出力されるディレクトリです。コンパイル時のエラーや警告は、ログファイルとして出力されます。

コマンド`build`によって自動生成されます。

### `output/`

バイナリファイルを実行した際の標準出力が、ログファイルとして出力されるディレクトリです。

コマンド`run`によって自動生成されます。

### `.sasarc.yaml`

Sasa の設定ファイルです。ディレクトリ名・ファイル名などを変更できます。

コマンド`run`によって自動生成されます。

### `answer.txt`

採点時に比較される出力結果ファイルです。このファイルの内容を正しい出力結果とみなし、対象の出力結果と同じ行 1 つにつき、対象を 1 点加点します。

コマンド`run`によって自動生成されます。模範解答プログラムの出力結果をコピー&ペーストして下さい。

### `input.txt`

標準入力として使用するファイルです。

課題で`scanf()`等の標準入力を扱う際に使用します。必要な場合、自身で作成して下さい。

### `database.csv`

採点結果等が記録されるデータベースファイルです。

コマンド`encode`によって自動生成され、`build`/`run`/`check`によって参照・更新されます。

各列は以下を表しています。

- ID : 学籍番号
- Name : 名前
- Charset : 推測された元の文字コード
- CanCompile : コンパイルできたか
- Score : 出力結果において`answer.txt`と同じ行の個数

## Tips

### Excel ではなく VSCode で CSV ファイルを確認する

[Rainbow CSV](https://marketplace.visualstudio.com/items?itemName=mechatroner.rainbow-csv)という拡張機能が便利です。

### `fopen()`等のファイル入力を使った課題を採点する

**作業ディレクトリから見たパス**にファイルを配置して下さい。

例えば、以下のような場合には

```c
fp = fopen("input.csv")
```

作業ディレクトリ（例：workspace）直下に読み取るファイルを配置して下さい。

```bash
$ cd workspace
$ ls
src/  input.csv  ...
```

以下のような場合には

```c
fp = fopen("./resource/input.csv")
```

`workspace/resource/input.csv`に読み取るファイルを配置して下さい。

```bash
$ cd workspace
$ ls
src/  resource/  ...
$ ls resource
input.csv  ...
```

### `scanf()`等の標準入力を使った課題を採点する

作業ディレクリ直下に`input.txt`を作成し、中に標準入力を記述して下さい。

例えば、以下のような場合には

```c
printf("xを整数で入力してください：\n");
scanf("%d", &x);
printf("yを整数で入力してください：\n");
scanf("%d", &y);
```

`input.txt`内に以下のように記述して下さい。改行区切りで一つずつ読み込まれます。空行は無視されます。

```txt
12
345
```

この時、`answer.txt`に使用する出力結果は、同じ`input.txt`を使用し、リダイレクト`>`を用いて取得して下さい。

```bash
$ ls
input_prob_ans.c  input.txt  ...
$ gcc input_prob_ans.c -o input_prob_ans
$ input.txt > ./input_prob_ans
xを整数で入力してください： ...
yを整数で入力してください： ...
...
```

## License

[MIT](./LICENSE)
