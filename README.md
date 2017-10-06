# go-diff

・指定のサイトをクロール
　url.Parse()
　http.Get()

※後々、スクレイピングも加える
　goquery.NewDocumentFromReader()

・前回クロールした内容をSQLiteから取得
　https://github.com/mattn/go-sqlite3

・前回と今回とで差分があればDIFFを出力
　・標準出力
　・マークダウン
　・CSV

◆サブコマンド◆
・init　・・・SQLiteデータベース初期化等
・reinit　・・・SQLiteデータベース初期化等
・

○設定○
・クロールするサイトURL
・DIFF出力先
