# Stella

Olá! Esse é um compilador escrito só por diversão! 

Baseado no "Writing an Interpreter in Go" de Thorsten Ball.

## Análise Léxica (Linear)
Análise léxica seria uma leitura de um fluxo de caracteres escrito em código-fonte de um programa e traduzido para um
conjunto de Tokens. Estes tokens são sequências de caracteres que têm um significado coletivo (Aho, Sethi e Ullman).

```go
montante := deposito_inicial + taxa_juros * 60
```

Estes seriam lidos pelo scanner como:

1.  Identificadores: montante, deposito_inicial e taxa_juros
2.  Atribuição: :=
3.  Adição: +
4.  Multiplicação: *
5.  Número inteiro: 60

Obs: Em algumas linguagens como o Python, a identação é significante e deve ser considerada pelo interpretador, no entanto não será avaliada nessa implementação. Espaços em branco entre os tokens também não serão considerados.

## Referências

1. "Writing an Interpreter in Go" by Thorsten Ball
2. "Compilers, Principles, Techniques and Tools" by Aho, Sethi e Ullman (Dragon's cover book)
3. ""