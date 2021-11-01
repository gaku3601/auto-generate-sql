# 導入方法
https://github.com/gaku3601/auto-generate-sql/releases  
から各環境に適した実行ファイルをダウンロードする  
## mac
agsをダウンロードして.zshrcなどで定義しているpathの通っているところへ配置する(/usr/local/bin等)
以下コマンドで実行権限を与える
```
chmod 777 ags
```

ags -hでヘルプが表示されれば利用可能です  
セキュリティ周りでpopupが出る場合はmacのセキュリティープライバシーの設定からagsを許可してください　　

## win
ags.exeをダウンロードしてpathの通っているところへ配置する

# 使い方
```
ags generate -p [対象エクセルファイル]
```
を実行することで自動的にSQLファイルが作成される  
SQLは対象エクセルファイルと同一のフォルダに配置される