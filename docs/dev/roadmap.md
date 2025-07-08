# Zinc Development Roadmap

## Phase 1: Language Definition & Core Compiler (Months 1-3)

- Finalize language grammar and features
- Implement lexer, parser, AST structures
- Semantic analyzer and basic type checker
- Initial IR design and simple code generation (Go/LLVM IR)
- Basic CLI tooling (build, run)

## Phase 2: Standard Library & Tooling (Months 4-6)

- Build standard library modules (math, strings, IO, collections)
- Implement formatter (znfmt), linter (znvet), and doc generator (zndoc)
- Add more comprehensive tests (unit, integration, semantic)
- Start benchmarking and performance tuning

## Phase 3: Bootstrapping & Self-hosting (Months 7-9)

- Develop minimal Zinc compiler subset in Go for bootstrapping
- Port compiler core to Zinc language itself
- Create build pipeline for self-hosted compilation
- Verify correctness and stability of bootstrapped compiler

## Phase 4: Release & Ecosystem Growth (Months 10+)

- Package stable releases and Docker images
- Write tutorials, sample projects, and documentation
- Encourage community contributions
- Plan advanced features: generics, interfaces/traits, concurrency primitives

---

## Continuous Tasks

- Code reviews and refactoring
- Documentation updates
- Issue tracking and bug fixing

---

*Dates and scope are estimates and may be adjusted based on progress and priorities.*