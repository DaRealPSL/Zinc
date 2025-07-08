# Zinc Language Features

This document expands on the core language features outlined in the README.

---

## Data Types

- **Primitive Types:**
  - `int` — 64-bit signed integers
  - `float` — 64-bit floating point numbers
  - `bool` — Boolean values: `true` or `false`
  - `string` — UTF-8 encoded text
  - `void` — Special type for functions that do not return a value

- **Composite Types:**
  - Arrays: homogeneous sequences of values, e.g. `int[]`, `string[]`
  - Objects: key-value maps/dictionaries with string keys and typed values

---

## Variables

- Mutable with explicit type annotation:
  ```zinc
  int x = 5;
  ```

* Mutable with type inference using `let`:

  ```zinc
  let x = 5;
  ```
* Immutable constants declared with `const`:

  ```zinc
  const pi = 3.1415;
  ```

---

## Control Flow

* Conditional statements:

    * `if`, `else` blocks
* Loops:

    * `for` loops with initialization, condition, and increment
    * `while` loops
* Early function exit with `return` statements

---

## Functions

* Declared using the `function` keyword:

  ```zinc
  function add(int a, int b) -> int {
      return a + b;
  }
  ```
* Typed parameters and explicit return types
* Support for function overloading planned as a future enhancement

---

## Classes & Objects

* Classes with fields and methods:

  ```zinc
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
* `init` function acts as the constructor
* `this` keyword for instance context
* Single inheritance supported initially; interfaces/traits planned for future

---

## Modules

* External files imported via `import` keyword:

  ```zinc
  import "utils.vix";
  ```
* Basic namespace support to avoid symbol collisions

---

## Built-in Library

* Core libraries included out of the box:

    * **Math**: constants (e.g., `Math.PI`), functions (e.g., `Math.sqrt`)
    * **Strings**: length, casing functions, concatenation
    * **IO**: `print()`, `input()`
    * **Collections**: arrays and objects manipulation

---
