# Zinc Grammar Specification (Draft)

## Lexical Tokens

- **Keywords:** `function`, `class`, `if`, `else`, `for`, `while`, `return`, `let`, `const`, `import`, `true`, `false`, `null`, `void`
- **Identifiers:** `[a-zA-Z_][a-zA-Z0-9_]*`
- **Literals:**
    - Integer: `[0-9]+`
    - Float: `[0-9]+\.[0-9]+`
    - String: `"([^"\\]|\\.)*"`
    - Boolean: `true` | `false`
- **Operators:** `+`, `-`, `*`, `/`, `%`, `=`, `==`, `!=`, `<`, `>`, `<=`, `>=`, `&&`, `||`, `!`
- **Delimiters:** `;`, `{`, `}`, `(`, `)`, `,`, `.`

## Grammar (EBNF-like)

```ebnf
program         ::= { statement } ;

statement       ::= variable_decl
                  | function_decl
                  | class_decl
                  | expression_stmt
                  | if_stmt
                  | for_stmt
                  | while_stmt
                  | return_stmt
                  | import_stmt
                  ;

variable_decl   ::= ("let" | "const" | type) identifier [ "=" expression ] ";" ;

function_decl   ::= "function" identifier "(" [ param_list ] ")" [ "->" type ] block ;

param_list      ::= param { "," param } ;

param           ::= type identifier ;

class_decl      ::= "class" identifier block ;

block           ::= "{" { statement } "}" ;

expression_stmt ::= expression ";" ;

if_stmt         ::= "if" "(" expression ")" block [ "else" block ] ;

for_stmt        ::= "for" "(" variable_decl expression ";" expression ")" block ;

while_stmt      ::= "while" "(" expression ")" block ;

return_stmt     ::= "return" [ expression ] ";" ;

import_stmt     ::= "import" string_literal ";" ;

expression      ::= assignment ;

assignment      ::= logical_or [ "=" assignment ] ;

logical_or      ::= logical_and { "||" logical_and } ;

logical_and     ::= equality { "&&" equality } ;

equality        ::= comparison { ("==" | "!=") comparison } ;

comparison      ::= addition { ("<" | ">" | "<=" | ">=") addition } ;

addition        ::= multiplication { ("+" | "-") multiplication } ;

multiplication  ::= unary { ("*" | "/" | "%") unary } ;

unary           ::= ("!" | "-" ) unary | primary ;

primary         ::= identifier
                  | literal
                  | "(" expression ")"
                  | function_call
                  | member_access
                  ;

function_call   ::= identifier "(" [ argument_list ] ")" ;

argument_list   ::= expression { "," expression } ;

member_access   ::= primary "." identifier ;

literal         ::= integer_literal
                  | float_literal
                  | string_literal
                  | boolean_literal
                  | "null"
                  ;

type            ::= "int" | "float" | "bool" | "string" | "void" | identifier ;

identifier      ::= [a-zA-Z_][a-zA-Z0-9_]* ;

integer_literal ::= [0-9]+ ;

float_literal   ::= [0-9]+"."[0-9]+ ;

string_literal  ::= "\"" ([^"\\]|\\.)* "\"" ;

boolean_literal ::= "true" | "false" ;
```

> Notes:
> 
> This grammar is a draft and may evolve.
> 
> The parser will be implemented as a recursive descent parser based on this spec.