## memo
ひとまずechoのバージョンを下げることで動作には漕ぎ付ける。
2.x系なら問題ないらしい。

### バージョンの下げ方
バージョンリスト
https://github.com/labstack/echo/releases?after=v3.0.0-beta.1

下記コマンドで、V2.1になる

```
cd ~/gopath/src/github.com/labstack/echo/
git checkout 04e6901
```
