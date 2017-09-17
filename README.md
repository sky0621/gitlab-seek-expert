# gitlab-seek-expert

### ■ツールの説明

##### gitlabのリポジトリ一覧（コミット数の多い順にコミッターも表示）を生成

### ■実行方法

##### [前提]

###### ・go version 1.7

###### ・glide インストール済み

###### ・オープンソースのcommunity editionの方を使用

##### 1. glide up を実行

###### ※Mac or Linux環境の場合、下記コマンドでインストール可能

###### curl https://glide.sh/get | sh

###### 参考：「https://github.com/Masterminds/glide」

##### 2. 下記の起動時パラメータを与えて cmd 配下で go run main.go を実行

###### ・host ・・・GitLabホスト（「http(s)」プロトコルはじまりのホスト名）

###### ・pkey ・・・GitLabでの自分のアカウントのPrivate token

### ■参考

##### https://docs.gitlab.com/ce/api/README.html

##### https://github.com/xanzy/go-gitlab
