# Zinc Architecture Overview

## Compiler Phases

1. **Lexical Analysis**  
   Converts source code into tokens (keywords, literals, operators).

2. **Parsing**  
   Transforms tokens into an Abstract Syntax Tree (AST).

3. **Semantic Analysis**  
   Validates AST for type correctness, scope rules, and language semantics.

4. **Intermediate Representation (IR) Generation**  
   Converts AST to IR suitable for backend codegen or interpretation.

5. **Code Generation**  
   Emits target code (initially Go code or LLVM IR, later native binaries or bytecode).

6. **Runtime**  
   Provides built-in library and runtime support for Zinc programs.

## Modular Design

- Clear separation between compiler phases enables easier testing and future enhancements.
- Internal packages encapsulate implementation details (`internal/lexer`, `internal/parser`, `internal/sema`, etc.).
- Public packages expose API for tooling and embedding Zinc compilation.

## Tooling Integration

- CLI (`cmd/zn`) supports building, running, testing Zinc code.
- Additional tools for formatting (`znfmt`), linting (`znvet`), and documentation (`zndoc`).

## Bootstrapping Strategy

- Initial compiler written in Go (bootstrap stage).
- Gradual port of compiler core to Zinc itself for full self-hosting.