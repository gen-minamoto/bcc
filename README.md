# bcc
Bank Code Checker

## 使い方
```
$ bcc <bank code(comma delimited)>
```
例
```
$ bcc 0001,0005,0009
0001,みずほ,ミズホ,みずほ,mizuho
0005,三菱ＵＦＪ,ミツビシユ－エフジエイ,みつびしゆ－えふじえい,mitsubishiyu-efujiei
0009,三井住友,ミツイスミトモ,みついすみとも,mitsuisumitomo
```

## メモ
金融機関コードの情報は [ここ](https://raw.githubusercontent.com/zengin-code/source-data/master/data/banks.json) の情報を使っています

一度取得した情報はtmpに置かれ、ファイルが存在しない場合のみリクエストして保存します
