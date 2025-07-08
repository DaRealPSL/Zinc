# Step 1: Define Language Goals and Specs

---

## 1.1 Core Language Philosophy (From README)

* Beginner-friendly and lightweight
* Explicit typing but also inference (`let`, `var`)
* Clear and readable syntax inspired by JS + C++
* Support for constants (`const`)
* OOP support with classes, methods, constructors
* Modular imports for clean code organization

---

## 1.2 Language Features (expanded from README)

* **Data types**:

    * Primitive: `int`, `float`, `bool`, `string`, `void` (for functions)
    * Composite: arrays, objects (maps/dictionaries)
* **Variables**:

    * Mutable with explicit type: `int x = 5;`
    * Mutable with inference: `let x = 5;`
    * Immutable with `const`
* **Control flow**:

    * `if`, `else`
    * Loops: `for`, `while`
    * Early `return` from functions
* **Functions**:

    * `function` keyword, typed parameters and return type
    * Support for function overloading (optional future)
* **Classes & Objects**:

    * Classes with fields and methods
    * Constructor via `init` function
    * `this` keyword to access instance fields
    * Single inheritance (future: interfaces or traits)
* **Modules**:

    * `import` keyword for external files
    * Namespace support (basic for now)
* **Built-in library**:

    * Math, Strings, IO (print, input), Collections

---

## 1.3 Syntax Style Summary

* Statement terminator: semicolon `;` mandatory
* Block delimiters: `{ ... }`
* Comments:

    * Single-line: `// comment`
    * Multi-line: `/* comment */`

---

## 1.4 Data Types Summary

| Type     | Description                | Syntax Example         |
| -------- | -------------------------- | ---------------------- |
| `int`    | Integer                    | `int x = 10;`          |
| `float`  | Floating-point number      | `float pi = 3.14;`     |
| `bool`   | Boolean                    | `bool flag = true;`    |
| `string` | Text                       | `string s = "hello";`  |
| `void`   | No value (function return) | `function foo(): void` |

---

## 1.5 Control Flow Syntax

```Zinc
if (x > 5) {
    print("x is greater");
} else {
    print("x is not greater");
}

for (int i = 0; i < 10; i++) {
    print(i);
}

while (x > 0) {
    x = x - 1;
    print(x);
}
```

---

## 1.6 Functions

* Functions declared with `function` keyword
* Parameters have explicit types
* Return type declared after parameters with colon
* Example:

```Zinc
function add(int a, int b): int {
    return a + b;
}
```

---

## 1.7 Classes

* Declared with `class` keyword
* Fields declared inside class body
* Methods declared as functions inside class
* Constructor named `init`
* Example:

```Zinc
class Animal {
    string name;

    function init(string name) {
        this.name = name;
    }

    function speak() {
        print(this.name + " makes a sound");
    }
}
```

---

## 1.8 Modules & Imports

* Import external files with `import` keyword
* Simple path string as import target (future: module namespaces)

```Zinc
import "utils.zn";

function main() {
    let sum = utils.add(1, 2);
    print(sum);
}
```

---

# Draft of docs/dev/grammar\_spec.md

```markdown
# Zinc Language Grammar Specification

## Overview
This document specifies the formal grammar and syntax rules of Zinc.

## Lexical Elements

### Comments
- Single-line comments start with `//` and extend to end of line
- Multi-line comments are enclosed in `/* ... */`

### Identifiers
- Start with a letter or underscore, followed by letters, digits, or underscores
- Case-sensitive

### Keywords
```

const, let, var, int, float, bool, string, void, function, class, if, else, for, while, return, import, this, true, false, null

````

### Literals
- Integer literals: `123`, `0`, `-42`
- Float literals: `3.14`, `0.0`
- Boolean literals: `true`, `false`
- String literals: `"hello"`, `"abc\n"`
- Null literal: `null`

### Operators
- Arithmetic: `+`, `-`, `*`, `/`, `%`
- Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical: `&&`, `||`, `!`
- Assignment: `=`
- Member access: `.`
- Others: `;`, `,`, `:`, `(`, `)`, `{`, `}`, `[`, `]`

## Syntax Rules

### Program
A program consists of zero or more import statements followed by zero or more declarations (functions, classes, variables).

```ebnf
program ::= { import_stmt } { declaration }
````

### Import Statement

```ebnf
import_stmt ::= "import" STRING_LITERAL ";"
```

### Declarations

* Variable Declaration:

```ebnf
var_decl ::= type identifier "=" expression ";"
          | "let" identifier "=" expression ";"
          | "const" identifier "=" expression ";"
```

* Function Declaration:

```ebnf
function_decl ::= "function" identifier "(" [param_list] ")" [":" type] block
param_list ::= param { "," param }
param ::= type identifier
```

* Class Declaration:

```ebnf
class_decl ::= "class" identifier "{" { class_member } "}"
class_member ::= var_decl | function_decl
```

### Statements

* If Statement:

```ebnf
if_stmt ::= "if" "(" expression ")" block [ "else" block ]
```

* For Loop:

```ebnf
for_stmt ::= "for" "(" var_decl expression ";" expression ")" block
```

* While Loop:

```ebnf
while_stmt ::= "while" "(" expression ")" block
```

* Return Statement:

```ebnf
return_stmt ::= "return" [ expression ] ";"
```

* Expression Statement:

```ebnf
expr_stmt ::= expression ";"
```

* Block:

```ebnf
block ::= "{" { statement } "}"
statement ::= var_decl | if_stmt | for_stmt | while_stmt | return_stmt | expr_stmt | block
```

### Expressions

* (To be expanded: binary, unary, function calls, member access, literals, identifiers)

---