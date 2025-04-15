---
title: "[BR] C++ Addons no Node.js"
description: Uma pequena introdução ao uso de C++ Addons no Node.js
date: 2024-12-02 18:00:00
---
### O que são C++ Addons?

São uma maneira de aumentar o desempenho e a funcionalidade do seu código em Node.js, estendendo um código C++ diretamente no ambiente de JavaScript. Em outras palavras, um addon em C++ é um módulo que você cria para o Node.js que permite que você escreva funções e bibliotecas em C++ e as use no seu código JavaScript.

São basicamente uma ponte entre o JavaScript e o C++, o que faz com que o Node.js possa executar códigos mais pesados sem perder a flexibilidade e a acima de tudo, sua simplicidade.

### Razões para usar C++ Addons no Node.js

- **Perfomance:** C++ é uma linguagem compilada, ele geralmente apresenta uma performance superior quando comparado ao JavaScript, que é uma linguagem interpretada. Além de permitir um uso mais controlado sobre a alocação e desalocação de memória.
- **Bibliotecas e código legado:** Muitos sistemas e bibliotecas legados foram desenvolvidos em C++. Os addons permitem integrar essas bibliotecas de forma eficiente no ambiente, facilitando a reutilização de bases de código já existentes e reduzindo o esforço de reescrita. Além de possuir o acesso a uma gama de bibliotecas que não funcionam diretamente no JavaScript
- **Nível de sistema:** Certos recursos a nível de sistema, que não são acessíveis pelo JavaScript, podem ser utilizados com a ajuda dos addons, permitindo o uso de funcionalidades específicas sempre que necessário.

### Como os C++ Addons funcionam?

- **Compilação:** Um addon precisa ser compilado antes de ser usado. Para isso, é necessário ter o [node-gyp](https://github.com/nodejs/node-gyp) configurado no seu repositório. Essa ferramenta compila o código C++ para um módulo "nativo" que o Node.js consegue executar.
- **Bindings:** Para a criação de uma "ponte" entre o JavaScript e o C++, você pode usar os pacotes [N-API](https://nodejs.org/api/n-api.html), [NAN](https://github.com/nodejs/nan) e/ou diretamente através do V8, libuv, e bibliotecas alternativas do Node.js.
- **Carregando as funcionalidades:** Assim que o código for compilado e uma "ponte" existir entre os dois mundos, as funções criadas podem ser chamadas através do **require()**, fazendo com que elas sejam acessíveis para o JavaScript

### Exemplo Prático

Primeiramente, vamos criar um diretório para nosso código

```bash
 mkdir addon
 cd addon
```

Em seguida, vamos criar inicializar nosso **package.json** e baixar a lib [**node-gyp**](https://github.com/nodejs/node-gyp) (certifique-se de instalar corretamente as dependências no seu SO)

```bash
npm init -y
npm i node-gyp
```

Dentro do nosso **package.json**, vamos criar um script de build para nossa addon.

```json
 "scripts": {
    "build": "node-gyp configure build"
  }
```

Agora, iremos criar nossa addon em C++ e vamos configurar nossa binding.

```cpp
// array_sum.cpp
#include <node.h>
namespace demo
{

  using v8::Array;
  using v8::Context;
  using v8::Exception;
  using v8::FunctionCallbackInfo;
  using v8::Isolate;
  using v8::Local;
  using v8::Number;
  using v8::Object;
  using v8::Value;

  void SumArray(const FunctionCallbackInfo<Value> &args)
  {
    Isolate *isolate = args.GetIsolate();
    Local<Context> context = isolate->GetCurrentContext();

    Local<Array> array = Local<Array>::Cast(args[0]);
    uint32_t length = array->Length();
    double sum = 0;

    for (uint32_t i = 0; i < length; ++i)
    {
      Local<Value> element = array->Get(context, i).ToLocalChecked();
      if (element->IsNumber())
      {
        sum += element->NumberValue(context).FromJust();
      }
    }

    args.GetReturnValue().Set(Number::New(isolate, sum));
  }

  void Initialize(Local<Object> exports)
  {
    NODE_SET_METHOD(exports, "sum", SumArray);
  }

  NODE_MODULE(NODE_GYP_MODULE_NAME, Initialize)
}

```

```json
// binding.gyp
{
  "targets": [{
    "target_name": "array_sum",
    "sources": [ "array_sum.cpp" ],
  }]
}
```

Pronto, assim que criado esses dois arquivos, podemos rodar nosso script de build (**npm run build**) e aguardar todo o processo para podermos desfrutar da nossa addon. Em seguida, iremos criar um novo arquivo e executar no terminal o comando **node index.js**

```js
//index.js
const addon = require('./build/Release/array_sum');

const big = [45, 'teste', 5]
const sum = addon.sum(big);
console.log('Resultado:', sum);
// Resultado: 50
```

### Considerações finais

Os addons são recursos poderosos quando o objetivo é otimizar o desempenho em operações críticas ou realizar integração com código nativo. Embora demandem conhecimento em C++ e aumentem a complexidade do projeto, eles podem ser a solução perfeita para situações em que o JavaScript puro não oferece a solução ideal. Graças a ferramentas como a **N-API**, o desenvolvimento de addons se tornou mais acessível e estável, permitindo que seus projetos combinem a praticidade do JavaScript com a eficiência do C++.

**Link do Repositório:** [cpp-addon-nodejs](https://github.com/rhyzzor/cpp-addon-nodejs)

**Referências:** [Node.js](https://nodejs.org/docs/latest/api/)