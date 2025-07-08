# Zinc Design Decisions

## Explicit Semicolons

We chose mandatory semicolons to simplify parsing and reduce ambiguity.

## Static Typing with Type Inference

Zinc is explicitly typed for clarity and performance but supports `let` for inferred typing to improve ergonomics.

## Single Inheritance, No Multiple Inheritance

To keep the object model simple and avoid complexity, Zinc supports single inheritance. Future interfaces/traits will extend flexibility.

## Minimal Core Language

We prioritize a small core with powerful abstractions. Complex utilities belong in the standard library or external modules.

## Module System

Files are imported via string literals for straightforward dependency management, inspired by JavaScript and Go.

## Tooling

Built-in formatting, linting, and documentation tooling promote code quality and maintainability from the start.

## Bootstrapping

Initial bootstrap compiler written in Go for reliability; self-hosting via Zinc compiler is a future goal.

