# tsvToJson

* 引数
  * 第1引数：入力ファイル
  * 第2引数：出力ファイル
  * 第3引数以降：Json定義（[例]0:key 5:title）
    * tsvの0カラム目をjsonのkeyという項目名で出力
  
* コマンド例
  * `$ go run tsvToJson.go digital_20160130.tsv digital.json 0:_key 7:title 14:author`
