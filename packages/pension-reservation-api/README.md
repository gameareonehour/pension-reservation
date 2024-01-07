# Pension Reservation API

## 概要

ペンション予約サイトに必要なAPIを提供する.


## アーキテクチャ

### モジュラーモノリス

`pension-reservation-api` では[モジュラーモノリス](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity)をベースとしたアーキテクチャを採択している。

モジュラーモノリスで設計することで、コンテキスト毎にコードを凝集でき、開発者のシステム理解に貢献できる。
ただモジュールは必要に応じてマイクロサービス化可能であることが前提となるので、モジュール自体は疎結合でなければならない。

今回モジュールの独立性を確保するために、モジュールに対してデータベースの内部構造を隠蔽している。
この点について、[CQRS](https://learn.microsoft.com/ja-jp/azure/architecture/patterns/cqrs)パターンを採用し、モジュール側でインターフェースを定義し、Main がそれに対して具体クラスを依存性注入する設計となっている。

### DI

今回、モジュールの独立性を高める、またモジュールのテスト容易性を高めるためにDIコンテナを採択し、レイヤー毎の依存関係の管理を柔軟にしている。

### OpenAPI

また YAMLで書かれた[API定義](https://www.openapis.org/)を元にサーバーサイド/クライエントサイドのコード生成を行っており、生成したコードではルーティング周りは担当している。

## コードベース

**ディレクトリ構造**

```
├── main.go       -- アプリケーションを起動するためのファイル
├── cmd           -- CLIアプリケーションを提供するディレクトリ
├── core          -- ログ、環境変数、API設定などアプリケーションの核となる機能を配置する
├── manipulation  -- データベースへのCRUDを行うオブジェクト(CQRSの具体オブジェクト)を配置する
├── mod           -- コンテキストで分けられたモジュールを配置する、ドメイン領域でアプリケーションの中心
├── model         -- ORMファイルを配置する
├── openapi       -- OpenAPIに関連するコードを配置する
│   ├── generated -- oapi-codegenが生成したコードを配置する
│   └── server    -- generated と mod を仲介するためのコード
├── tmp           -- airが使う作業用ディレクトリ
└── static        -- フロントエンドに配信する画像等を配置する
```

**実行順序**

1. Manipulation 生成
2. ModuleService 作成
3. ModuleService をDIコンテナに登録
4. Controller を作成
5. Controller をDIコンテナに登録
6. OpenAPIServer を作成
7. OpenAPI側に OpenAPIServerを渡し、ルーティング処理を委譲

**依存関係**

![Untitled-2024-01-07-0316](https://github-production-user-asset-6210df.s3.amazonaws.com/91713318/294710194-94754dfd-c04d-4d7e-afd9-7c2b38fc9bf1.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20240106%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240106T184700Z&X-Amz-Expires=300&X-Amz-Signature=4783265135191b4326024d50ac29c1aa104aad3ddeea89dc3e88fd5a23569e90&X-Amz-SignedHeaders=host&actor_id=91713318&key_id=0&repo_id=697608441)
