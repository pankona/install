# install

Install some awesome tools on linux to make it development ready for me

## やることメモ

- vim
  - vim をインストール
  - vimrc を配置する & ftplugin などを置く
  - パッケージマネージャ (dein) をインストール
- bashrc
  - 自分の .bashrc をクローン
  - 既存の .bashrc に、クローンした自分の .bashrc を読み込ませるように処理を追記する
- build essential のインストール
  - Ubuntu だけかも
  - apt install build-essential する
- asdf
  - asdf をインストール
  - nodejs、yarn、ruby、ghq くらいは入れておく
- git
  - git config global user.email yosuke.akatsuka@gmail.com
  - git config global user.name pankona
  - git config global core.editor vim
- gh (GitHub の CLI)
  - gh を入れる
  - (認証は自動化しない)
- ghq
  - git config に設定を追加して、~/go/src を root として追加する
- docker、docker compose
  - docker 入れる
  - docker compose 入れる
- Go のインストール
  - なんとなく最新版を入れる

## さらにメモ

Mac にインストールしたり Ubuntu にインストールしたり Alpine にインストールしたり、色々なシチュエーションがあると思うので、
- install スクリプトのエントリーポイントは一個にする
- 明らかに OS に特化しているところ (apt とか brew とか) はファイルを分けて分岐する
- 共通で使えそうなところ (asdf するとか Go のインストールとか) はあえて分けなくてもいいかもしれないが、でもメンテナンスしやすさを考えたら分けておいたほうがいいだろう

## LICENSE

MIT

## Author


Yosuke Akatsuka (a.k.a pankona)
