# 🚀 Zinc Programming Language 🌟

Welcome to **Zinc**! A lightweight, beginner-friendly programming language designed for simplicity and power! Inspired by **JavaScript** and **C++**, Zinc makes coding intuitive without sacrificing robust features.

---

## ✨ Features

✅ **Explicit Typing**: Easily declare variable types for clarity.  
✅ **Dynamic Objects**: Use `let` or `var` for type inference when needed.  
✅ **Constants**: Keep values immutable with `const`.  
✅ **Functions**: Clean, readable syntax with the `function` keyword.  
✅ **Classes**: Enjoy object-oriented programming support.  
✅ **Modular Imports**: Keep your code clean with an `import` system.

---

## 🧠 Zinc Syntax Cheat Sheet

### 🔢 Variables

```Zinc
int x = 10;         // Integer variable
string y = "Hello";    // String variable
bool cool = true;   // Boolean variable
let z = 42;         // Generic variable (type inferred)
const pi = 3.1415;  // Immutable constant
```

### 🎉 Functions

```Zinc
function greet(str name) {
    print("Hello, " + name);
}
```

### 🛤️ Conditionals

```Zinc
if (x > 5) {
    print("x is greater than 5");
} else {
    print("x is 5 or less");
}
```

### 🔁 Loops

```Zinc
for (int i = 0; i < 10; i++) {
    print(i);
}
```

```Zinc
while (x > 0) {
    x = x - 1;
    print(x);
}
```

### 🧍 Objects

```Zinc
let person = {
    name: "DaPSL",         // string
    age: 14,                // int
    hobbies: ["coding", "gaming"] // array of strings
};
print(person.name);
```

### 🏛️ Classes

```Zinc
class Animal {
    str name;

    function init(str name) {
        this.name = name;
    }

    function speak() {
        print(this.name + " makes a sound");
    }
}
```

## 🛠️ Built-in Libraries

### 📚 Standard Libraries Included

Zinc has built-in tools for everyday coding tasks:

- **Math**: Handy constants like Math.PI and functions like Math.sqrt(x).
- **Strings**: Utilities like .length(), .toUpperCase(), and .toLowerCase().
- **Basic I/O**: Use print() and input() for quick input/output tasks.

### Example:

```Zinc
let radius = 5;
let area = Math.PI * (radius * radius); // Built-in Math library
print("Area: " + area);
```

## 📦 Importing Modules

Take it further with external or custom modules:

```Zinc
import "utils.zn"; // importing user-defined moduled

function main() {
    let result = utils.add(5, 10); // using a function from the imported module
    print(result);
}
```

## 🚀 Getting Started

1️⃣ Clone the Repository

```bash
git clone https://github.com/IAmThePSL/Zinc-programming-lang
cd Zinc
```

2️⃣ Build the Project

```bash

```

3️⃣ Run Zinc Programs

```bash
./Zinc examples/hello.zn
```

## 🤝 Contribution

We ❤️ contributors!
Found a bug? Have a cool feature idea? Feel free to open an issue or submit a pull request. Let’s make Zinc even better together! ✨

## 📜 License

This project is licensed under the [MIT License](https://github.com/DaRealPSL/Zinc/blob/main/LICENSE)

> [!NOTE] > **Zinc is currently under active development. Some features listed in this document may not yet be available. Thank you for your patience and support!**
