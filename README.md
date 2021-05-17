# humtodo

## dev tools

### tools/gen-ts-entity/main.go

- `// +gen-ts-entity`のコメントの次にある構造体の定義からTypeScriptのinterfaceを生成する
- フィールド名がそのままTS側のフィールド名になる
    - jsonタグを指定している場合はその名前になる
    - jsonタグでomitemptyを指定している場合は`optional`になる
- タグに`ts-import`を指定するとTSでimportを挿入してくれる

```golang

// +gen-ts-entity
type TestEntity struct {
    ID string `json:"id,omitempty"`
    User entities.UserEntity `json:"user" ts-import="../entities/entities"`
}

```

```
$ make gen-entity
```

`gen/xxx/xxx.ts`に出力される

```

import {
	UserEntity,
} from '../entities/entities';
export interface SignUpRequest {
	name: string;
	password: string;
	user: UserEntity;
};
```