# Zinc Syntax Style Summary

- **Statement terminator**: semicolon `;` is mandatory for all statements
- **Block delimiters**: curly braces `{ ... }` define code blocks for functions, loops, conditionals, classes, etc.
- **Comments**:
    - Single-line comments start with `//` and continue to end of line
    - Multi-line comments are enclosed by `/*` and `*/` and can span multiple lines

---

Example:

```zinc
// Single-line comment
let x = 5;

if (x > 0) {
    /*
     Multi-line comment
     explaining the following block
    */
    print("Positive");
}
