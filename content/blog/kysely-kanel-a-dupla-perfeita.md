---
ogImage: 
  component: BlogPost
  props:
    title: "[BR] Kysely + Kanel, a dupla perfeita"
    description: "Como configurar e integrar o Kysely, um query builder type-safe, com o Kanel, uma ferramenta que automatiza a geração de interfaces do banco de dados. Uma combinação perfeita para substituir o Raw SQL com segurança e eficiência."
sitemap:
  lastmod: 2025-04-15
robots: index, follow
schemaOrg:
  - "@type": "BlogPosting"
    headline: "[BR] Kysely + Kanel, a dupla perfeita"
    author: 
      type: "Person"
      name: "Rhyzzor"
    datePublished: "2024-12-24"
description: Como configurar e integrar o Kysely, um query builder type-safe, com o Kanel, uma ferramenta que automatiza a geração de interfaces do banco de dados. Uma combinação perfeita para substituir o Raw SQL com segurança e eficiência.
title: "[BR] Kysely + Kanel, a dupla perfeita"
date: "2024-12-24T18:00:00Z"

---

Eu sou uma pessoa que, desde que comecei a elaborar meus primeiros projetos (meu OT Pokémon e meus primeiros websites para o Habbo), sempre optei pelo **Raw SQL**. Sinceramente, ainda gosto bastante de escrever minhas próprias queries e ter um controle mais preciso sobre essa camada "low-level". Um ORM não me deixa totalmente confortável, pois já perdi dias analisando logs para identificar e otimizar queries ineficientes. 

Porém, em muitas codebases onde trabalhei com Raw SQL, a grande maioria não possuía controle de migrations, e o banco de dados tampouco era monitorado. Tudo funcionava na base do improviso: "Precisa de um novo campo? Roda um **ALTER TABLE** e adiciona uma nova coluna." Essa abordagem era extremamente prejudicial em todos os cenários, diversas questões surgiam, como: "Quais colunas devemos subir no ambiente de produção?", "Quais novas entidades foram criadas?", "Os ambientes estão sincronizados?" — e muitos outros problemas semelhantes.

### A solução para meus problemas

Diante de todos esses problemas, decidi adotar novas ferramentas para tornar minha rotina e as das equipes com quem trabalhei mais saudáveis. Eu não queria abrir mão da flexibilidade que tinha, mas também desejava controlar melhor os **graus de liberdade** da aplicação. Após muita pesquisa, encontrei uma ferramenta que considero a mais completa para resolver esses problemas: o **[Kysely](https://github.com/kysely-org/kysely)**, ele é um query builder para TypeScript que, além de prático, é completamente type-safe — um ponto super importante para mim. Essa lib chamou tanto minha atenção que comecei a contribuir ativamente na comunidade, tanto diretamente quanto indiretamente, criando plugins para outras bibliotecas open source integradas ao Kysely.

No entanto, uma das maiores dificuldades ao trabalhar com o Kysely é que, diferente de ORM's, ele não possui uma entidade ou uma geração automática de tipos/interfaces. Todo esse trabalho precisa ser feito manualmente, o que pode ser um pouco exaustivo. Durante minhas pesquisas por soluções, encontrei uma ferramenta que acabei adotando em todos os meus projetos envolvendo PostgreSQL: o **[Kanel](https://github.com/kristiandupont/kanel)**. O Kanel gera automaticamente as tipagens do banco de dados, complementando perfeitamente o Kysely. 

Além disso, o **Kanel** possui um recurso adicional para uso direto com o **Kysely**: o **[Kanel-Kysely](https://github.com/kristiandupont/kanel)**. Tenho contribuído ativamente para esse repositório, ajudando a desenvolver novas features, como filtros de tipos para tabelas de migrations e a conversão de objetos do Zod para camelCase.


### Configurando o Kysely

Estarei usando NestJS para ilustrar os exemplos a seguir. Então, se você não entender alguma sintaxe ou algo no código, sugiro dar uma lida na documentação do [NestJS](https://github.com/nestjs/nest). Na minha opinião, ele é o melhor framework JavaScript — ainda mais se você quiser "fugir" do JavaScript. Mas isso é assunto para outro post meu.

De antemão, você precisará ter um repositório com o NestJS inicializado, caso queira seguir os exemplos à risca. Porém, você também pode desenvolver seu próprio código.

De início, vamos precisar instalar o próprio **Kysely**, sua CLI e o módulo do PostgreSQL para o Node.js.

```bash
npm i kysely pg && npm i kysely-ctl --save-dev
```

Em seguida, vamos precisar criar um arquivo de configuração na raiz do projeto para o **Kysely**. Também vou utilizar o prefixo do Knex para os nossos arquivos de migrations e seeds.

```ts
// kysely.config.ts

import "dotenv/config";

import { defineConfig, getKnexTimestampPrefix } from "kysely-ctl";
import { Pool } from "pg";

export default defineConfig({
  dialect: "pg",
  dialectConfig: {
    pool: new Pool({ connectionString: process.env.DATABASE_URL }),
  },
  migrations: {
    migrationFolder: "src/database/migrations",
    getMigrationPrefix: getKnexTimestampPrefix,
  },
  seeds: {
    seedFolder: "src/database/seeds",
    getSeedPrefix: getKnexTimestampPrefix,
  },
});
```

Na sequência, vamos rodar o comando **`npx kysely migrate make create_user_table`** em nosso terminal. Ele será responsável por criar nossa primeira migration. Em seguida, vamos criar uma nova tabela de usuários e, assim que feito, vamos rodar essa migration no nosso banco de dados com o comando **`npx kysely migrate latest`**.

```ts
// 20241225222128_create_user_table.ts

import { sql, type Kysely } from 'kysely'


export async function up(db: Kysely<any>): Promise<void> {
  await db.schema
  .createTable("user")
  .addColumn("id", "serial", (col) => col.primaryKey())
  .addColumn("name", "text", (col) => col.notNull())
  .addColumn("email", "text", (col) => col.unique().notNull())
  .addColumn("password", "text", (col) => col.notNull())
  .addColumn("created_at", "timestamp", (col) =>
    col.defaultTo(sql`now()`).notNull(),
  )
  .execute();
}

export async function down(db: Kysely<any>): Promise<void> {
  await db.schema.dropTable("user").execute();
}
```

Com todos esses passos concluídos, vamos criar um módulo para a nossa base de dados. Repare também que estou usando um plugin do Kysely para converter nossas colunas para camelCase.

```ts
// src/database/database.module.ts

import { EnvService } from "@/env/env.service";
import { Global, Logger, Module } from "@nestjs/common";
import { CamelCasePlugin, Kysely, PostgresDialect } from "kysely";
import { Pool } from "pg";

export const DATABASE_CONNECTION = "DATABASE_CONNECTION";

@Global()
@Module({
  providers: [
    {
      provide: DATABASE_CONNECTION,
      useFactory: async (envService: EnvService) => {
        const dialect = new PostgresDialect({
          pool: new Pool({
            connectionString: envService.get("DATABASE_URL"),
          }),
        });

        const nodeEnv = envService.get("NODE_ENV");

        const db = new Kysely({
          dialect,
          plugins: [new CamelCasePlugin()],
          log: nodeEnv === "dev" ? ["query", "error"] : ["error"],
        });

        const logger = new Logger("DatabaseModule");

        logger.log("Successfully connected to database");

        return db;
      },
      inject: [EnvService],
    },
  ],
  exports: [DATABASE_CONNECTION],
})
export class DatabaseModule {}
```

### Configurando o Kanel

Vamos começar instalando nossas dependências.

```bash
npm i kanel kanel-kysely --save-dev
```

Em seguida, vamos criar nosso arquivo de configuração para o Kanel começar a fazer o seu trabalho. Repare que estarei utilizando alguns plugins, como o camelCaseHook (para transformar nossas interfaces em camelCase) e o kyselyTypeFilter (para excluir as tabelas de migrations do Kysely), uma dessas features eu tive o prazer de poder contribuir e facilitar ainda mais o trabalho que tínhamos.

```js
// .kanelrc.js

require("dotenv/config");

const { kyselyCamelCaseHook, makeKyselyHook, kyselyTypeFilter } = require("kanel-kysely");

/** @type {import('kanel').Config} */
module.exports = {
  connection: {
    connectionString: process.env.DATABASE_URL,
  },
  typeFilter: kyselyTypeFilter,
  preDeleteOutputFolder: true,
  outputPath: "./src/database/schema",
  preRenderHooks: [makeKyselyHook(), kyselyCamelCaseHook],
};

```

Assim que o arquivo for criado, vamos rodar o comando **`npx kanel`** em nosso terminal. Repare que foi criado um diretório no caminho especificado no arquivo de configuração. Esse diretório corresponde ao nome do seu schema, no nosso caso, o **Public**, e dentro dele temos dois novos arquivos: **PublicSchema.ts** e **User.ts**. Provavelmente, o seu **User.ts** estará exatamente assim:

```ts
// @generated
// This file is automatically generated by Kanel. Do not modify manually.

import type { ColumnType, Selectable, Insertable, Updateable } from 'kysely';

/** Identifier type for public.user */
export type UserId = number & { __brand: 'UserId' };

/** Represents the table public.user */
export default interface UserTable {
  id: ColumnType<UserId, UserId | undefined, UserId>;

  name: ColumnType<string, string, string>;

  email: ColumnType<string, string, string>;

  password: ColumnType<string, string, string>;

  createdAt: ColumnType<Date, Date | string | undefined, Date | string>;
}

export type User = Selectable<UserTable>;

export type NewUser = Insertable<UserTable>;

export type UserUpdate = Updateable<UserTable>;

```

No entanto, o mais importante é o arquivo fora desse diretório **Public**, o arquivo **Database.ts**, porque é ele quem vamos repassar para o Kysely entender toda a estrutura da nossa base de dados. Dentro do nosso arquivo **app.service.ts**, vamos injetar o nosso provider do DatabaseModule e repassar para o Kysely o nosso tipo **Database**.

```ts
// src/app.service.ts

import { Inject, Injectable } from "@nestjs/common";
import { Kysely } from "kysely";
import { DATABASE_CONNECTION } from "./database/database.module";
import Database from "./database/schema/Database";

@Injectable()
export class AppService {
  constructor(@Inject(DATABASE_CONNECTION) private readonly db: Kysely<Database>) {}

  async findManyUsers() {
    const users = await this.db.selectFrom("user").select(["id", "name"]).execute();

    return users;
  }
}
```

Repare que a tipagem que o Kanel gerou está funcionando corretamente, porque nosso editor de código irá sugerir justamente as colunas que criamos em nossa primeira migration.

![Sugestão do Editor de Código](/img/kysely-example.png)

### Considerações finais

Essa é uma dupla que gosto bastante de utilizar em meus projetos pessoais e até mesmo no trabalho (quando tenho a liberdade para isso). Um query builder é a ferramenta essencial para todos que gostam da flexibilidade que o Raw SQL oferece, mas também optam por um caminho "mais seguro". O Kanel também já me poupou muitas horas de debug e criação de novas tipagens. Eu recomendo fortemente que você crie um projeto com esses dois, você com certeza não irá se arrepender.

**Link do Repositório:** [frankenstein-nodejs](https://github.com/rhyzzor/frankenstein-nodejs)