# UPSIDERコーディング試験 2025/02/15

[課題内容](https://github.com/upsidr/coding-test/blob/main/web-api-language-agnostic/README.ja.md)

以下は課題内容のコピー

# APIコーディングテストについて

## はじめに

候補者の皆様へ

この度は、UPSIDERの選考に進んでいただき、誠にありがとうございます。

このテストは、「スーパー支払い君.com」という法人向けの架空のウェブサービスを立ち上げるためのREST APIを設計・実装することが目的です。このウェブサービスでは、ユーザーが未来の支払期日の請求書のデータを登録しておくと、期日に残高がなくとも自動的に銀行振り込みを行うことができ、現金の支出を最大一ヶ月遅らせることができるため、ユーザーにとって便利なウェブサービスとなることが期待されます。

このコーディングテストでは、下記のビジネス要件を満たすREST APIを設計・実装することが求められます。具体的には、次のユーザーストーリーに基づいたREST APIのエンドポイントを開発していただきます。

* 「ユーザーとして、請求書データを新規に作成することができる。なぜなら支払期日に確実に支払い処理を行いたいからだ。」
* 「ユーザーとして、指定期間内に支払いが発生する請求書データの一覧を取得することができる。なぜならどのような支払いを登録したか確認したいためだ。」

言語は Golang で記述してください。
フレームワークの指定は特にありません。その他のライブラリも自由にお使いください。

なお、このコーディングテストにおいて、データベースの指定はありませんが、NoSQLではなくRDBMSを使うことを前提にしています。

テストの目安としては、2~3時間程度を想定しております。

なお、GitHubを利用してソースコードを管理していただくため、Public RepositoryとしてPushしてURLを提出していただくこととなります。

ご自分の最大限の能力を発揮して、素晴らしいソフトウェアを開発してください！

## 課題/条件

### REST APIエンドポイントの設計

* POST /api/invoices
  * 新しい請求書データを作成する
    * `請求金額` は自動的に計算されるものとする
      * `支払金額` に手数料4%を加えたものに更に手数料の消費税を加えたものを請求金額とする
      * 例) 支払金額 `10,000` の場合は請求金額は `10,000 + (10,000 * 0.04 * 1.10) = 10,440`
* GET /api/invoices
  * 指定期間内に支払いが発生する請求書データの一覧を取得する

### データモデル

* 企業
  * 法人名
  * 代表者名
  * 電話番号
  * 郵便番号 
  * 住所
* ユーザー (企業に紐づく)
  * 氏名
  * メールアドレス
  * パスワード
* 取引先 (企業に紐づく)
  * 法人名
  * 代表者名
  * 電話番号
  * 郵便番号 
  * 住所
* 取引先銀行口座 (取引先に紐づく)
  * 取引先ID
  * 銀行名
  * 支店名
  * 口座番号
  * 口座名
* 請求書データ (企業・取引先に紐づく)
  * 発行日
  * 支払金額
  * 手数料
  * 手数料率
  * 消費税
  * 消費税率
  * 請求金額
  * 支払期日
  * ステータス (未処理、処理中、支払い済み、エラー)

### テストケース

* ユーザーが新しい請求書データを作成できる
  * リクエストが成功し、HTTPステータスコード200が返されること
  * レスポンスに作成された請求書データが含まれていること
* ユーザーが指定期間内に支払いが発生する請求書データの一覧を取得できる
  * リクエストが成功し、HTTPステータスコード200が返されること
  * レスポンスに指定期間内の請求書データが含まれていること

## コーディングテストの評価基準

* クラス、メソッド、構造体、関数などの責務や役割に対する考慮
* SOLID原則やアーキテクチャのベストプラクティスに関する理解度
* コーディングスタイルや可読性に対する考慮
* 適切なコミット単位でコードが管理されており、コードレビューがしやすい状態であるか
* 認証・認可に関する考慮がされているか
* 秘匿情報の取り扱いに関する考慮がされているか
* 適切なテストコードが実装されているか
* テストデータが大量にある場合や同時アクセス数が多い場合に、十分なパフォーマンスを発揮できるかどうか
* エラーや例外が発生した場合に、適切なエラー処理が行われているかどうか
* 必要に応じて、他の開発者がコーディングするために必要な知識をドキュメント化しているか

# Contributors

## How to started

1. `docker compose up` することでサービスを起動
2. `db/schema.sql` を実行することでスキーマ生成
3. `db/seeds/*.sql` を最初に実行し、Seedデータを生成

## Seeds

初期設定だけでは動作確認できないため最初に一定のcompaniesテーブルなどデータが必要になります。
参考までに、以下のようにテストデータの作成を行うことでテストが可能になります。

```sql

INSERT INTO companies
    (id, legal_name, representative_name, phone_number, postal_code, address)
VALUES
    ('01JM6BKQNM5RNPG3WDCX4YHJT2', '株式会社テスト', '山田太郎', '090-1234-5678', '123-4567', '東京都千代田区永田町1-7-1'),
    ('01JM6BKQNPK2VTHS5EFYM9KLU4', '株式会社テスト2', '佐藤一郎', '090-1234-5678', '123-4567', '東京都千代田区永田町1-7-1'),
    ('01JM6BKQNQW4XVJU7GHZN2MNV6', '株式会社テスト3', '鈴木花子', '090-1234-5678', '123-4567', '東京都千代田区永田町1-7-1');

INSERT INTO users
    (id, company_id, name, email, password_hash)
VALUES
    ('01JM6BKQNR8Y5XKV9JK2P4QRW8', '01JM6BKQNM5RNPG3WDCX4YHJT2', '山田太郎', 'yama@example.com', 'password1'),
    ('01JM6BKQNS9Z6YLW2KL3Q5RSX9', '01JM6BKQNPK2VTHS5EFYM9KLU4', '佐藤一郎', 'sato@example.com', 'password2'), 
    ('01JM6BKQNT2A7ZMX3LM4R6STY2', '01JM6BKQNQW4XVJU7GHZN2MNV6', '鈴木花子', 'suzuki@example.com', 'password3');

INSERT INTO business_partners
    (id, company_id, partner_company_id, partner_bank_account_id)
VALUES
    ('01JM6BKQNU3B8ANY4MN5S7TUZ3', '01JM6BKQNM5RNPG3WDCX4YHJT2', '01JM6BKQNM5RNPG3WDCX4YHJT2', '01JM6BKQNX6E3DQZ7PQ8V9WXC6'),
    ('01JM6BKQNV4C9BOZ5NO6T8UVA4', '01JM6BKQNPK2VTHS5EFYM9KLU4', '01JM6BKQNPK2VTHS5EFYM9KLU4', '01JM6BKQNY7F4ERZ8QR9W9XDD7'),
    ('01JM6BKQNW5D2CPZ6OP7U9VWB5', '01JM6BKQNQW4XVJU7GHZN2MNV6', '01JM6BKQNQW4XVJU7GHZN2MNV6', '01JM6BKQNZ8G5FSZ9RS0XAWE88');

INSERT INTO bank_accounts
    (id, company_id, bank_name, branch_name, account_number, account_name)
VALUES
    ('01JM6BKQNX6E3DQZ7PQ8V9WXC6', '01JM6BKQNM5RNPG3WDCX4YHJT2', '三井住友銀行', '本店', '1234567890', '山田太郎'),
    ('01JM6BKQNY7F4ERZ8QR9W9XDD7', '01JM6BKQNPK2VTHS5EFYM9KLU4', 'みずほ銀行', '本店', '1234567890', '佐藤一郎'),
    ('01JM6BKQNZ8G5FSZ9RS0XAWE88', '01JM6BKQNQW4XVJU7GHZN2MNV6', '三菱UFJ銀行', '本店', '1234567890', '鈴木花子');
```

## TODO

- [ ] package構成についてDocument書く
- [ ] テスト
  - [ ] httptestでE2Eテスト
- [ ] ロギング整備
- [ ] 認証整備
- [ ] GraphQLやgRPC導入でスキーマ駆動開発基盤整備
- [ ] インフラ構成と構成管理(terraformなどでIaC化)
- [ ] deploy
- [ ] CI/CD整備
