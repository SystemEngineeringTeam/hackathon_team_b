###### tags: `SyskenHACKTeamB`
# シス研ハッカソンTeamB 本番
[TOC]

## 組み合わせたもの
* 1 友達の作り方&好きなアニメがマイナーで悲しい 
* 2 春から〇〇大学の為に、共有できる情報とかをlineに限定しないものが欲しい&地雷授業に当たるのをなくしたい(確率統計とか)レビューサイトみたいなのがあればいいのに
* 3 上手いラーメン屋に行きたい&友達とご飯食べに行く時どこにしようか悩む&昼食を何にしたらいいか決めきれない
* 4 ふくだくんが寝坊する & 彼女ができない & 時刻前に起きれない
* 5 春から〇〇大学の為に、共有できる情報とかをlineに限定しないものが欲しい＆技術力が向上しない 

## 作るもの
* 2 春から〇〇大学の為に、共有できる情報とかをlineに限定しないものが欲しい&地雷授業に当たるのをなくしたい(確率統計とか)レビューサイトみたいなのがあればいいのに


## 機能書き出し
* 授業表示
* レビュー
* ↑投稿日時タイムスタンプ
* 学科別ソート
* 専攻別ソート
* 年別ソート
* 役に立ったorいいねボタン
* 人気別ソート
* 5段階or3段階の星で評価
* この口コミは参考になったのボタン
* レビューアーランキング
* ログイン
## 実装する機能(上から優先度高い)
* 授業表示
* 絞り込み(前期後期)
* ソート
* キーワード検索
* レビュー　星だけでもok
## 時間があれば実装したい
* 役に立ったボタン（レビューに対する評価)
* ログイン　
* ユーザー継続できる機能を追加したい
    
## 技術選定
フロント
* vue.js

バックエンド
* Golang（MySQL,firebaseのアクセス)

データベース
* MySQL

認証
* firebase

## api設計
* 授業関連
    * /lecture
      * get
      * parameter
          * grade:学年
          * Department:学部
          * semester:学期
          * dayofweek:曜日
          * time:時限
          * teacher:講師
          * lectureName:科目名
      
      * return
          * grade:学年
          * Department:学部
          * semester:学期
          * dayofweek:曜日
          * time:時限
          * teacher:講師
          * lectureName:科目名
          * grade:int
          * reviewStarAverage:int
          * index
   * /reviews
       * get
           * parameter
              * index?
           * return
              * favorite:int
              * good:int　　←これってカウントとかは...
              * review:string
              * 下の方に例を記述
       * post
           * body
               * reviewStar:int
               * sentence:string 
           
   * /search
       * get
       * parameter
           * keyword
       * return 
          * grade:学年
          * Department:学部
          * semester:学期
          * dayofweek:曜日
          * time:時限
          * teacher:講師
          * lectureName:科目名
          * grade:int
          * star:int
          * favorite:int
          * index



* 例（こんな感じで書いたら見やすいかも）
    | 変数名       |    型      |   意味        |
    | --------    |  --------  | --------     |
    | lectureName | string     | 授業の名前     |
    | grade       | int        | 学年          |
    | star        | int        | ?            |

      
<!-- [{
    favorite:1,
    review:くそつまらん,
},
{
    favorite:2,
    review:楽しい,
}] -->
    

## webページ構成
学部,開校学期,曜日,時限,科目名,担当教員





