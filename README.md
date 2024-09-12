# Taming the Dragon
or build a simple Compiler for studying.

## Why

For a while I want to dig deeper into the topic of Computer Languages and how they work. 
I thought building a simple Compiler and exploring different tools and solutions is a great project to understand more and expand my knowledge.

I will follow the book [Writing a C Compiler](https://nostarch.com/writing-c-compiler) as my main Guide and add on and dig deeper with aditional resources on the go.


## What

I will work on a subset of C, which I will call **`dragon-C`**. This subset will focus on key features of the C programming language, simplifying the full C language to make the compiler more manageable while still covering important concepts like variables, control structures, and functions.
The idea is to follow the book building a C-Compiler, and while at this explore different tools for lexing and parsing to also learn this toolsets, I also want to build a LLVM Frontend, so I have a excurse into another toolchain opening up possibilities.

## How


The project is split into different parts.

`docs` will keep general documentation of the project as well as a 'journal' in which I will keep the progress and write down my experience and aditional thoughts or explorations.

`hcc` is the Compiler driver. It is a Python-based CLI tool that orchestrates different stages of the compilation process. Initially, it will use existing tools such as `gcc` or `clang` for preprocessing and assembling, but over time, it will increasingly rely on my own implementation of `dragoncc` for compilation.

`src` keeps all the source code I intend to write:
 - `dragoncc` - dragon c compiler, will be the compiler, implemented in haskell
 - `llvm-frontend` - a LLVM frontend implemented in Rust for the same subset of C
 - `experiments` - here I will explore tools like `flex`, `bison`, `antlr4` and `tree-sitter` with the same langauge, to explore their usage.


 Maybe I will add:
 - Language Server - I will decide on the go.
 - A C repl - thus an interpreter - I will decide on the go.
 - Garbage Collection into Dragon-C - I will decide once the main project is done.


 I choose 'Haskell' for the Compiler, as as a functional language with good pattern matchin capabilities it is well suited for Compilers and I really want to dig into that language, so this seems like a great project to deepen my understanding.

For the LLVM frontend I want to use another language, that I want to explore, 'Rust'.

## Roadmap

- [x] Setup project scaffold
- [ ] Implement basic lexer (dragoncc)
- [ ] Implement parser (hand-written or with tools)
- [ ] Start working on LLVM frontend in Rust
- [ ] Add support for functions and control structures
- [ ] Investigate adding garbage collection to Dragon-C

## Credits & Resources

Here I will add Books I use, as well as articles, repos or people that will turn out to be helpful or inspiring on the progress.

- [Writing a C Compiler](Writing a C Compiler) - as my main driver
- [The Dragon Book](https://en.wikipedia.org/wiki/Compilers:_Principles,_Techniques,_and_Tools) - To look up theory on the way
- [Language Implementation Patterns](https://pragprog.com/titles/tpdsl/language-implementation-patterns/) - As a secondary resource on my way
- [Flex & Bison](https://www.oreilly.com/library/view/flex-bison/9780596805418/) - To learn these tools
- [Crafting Interpreters](https://www.youtube.com/c/tjdevries) - as tertiary resource on my way


- [LLVM Documentation](https://llvm.org/docs/) - For understanding LLVM's intermediate representation (IR) and toolchain.
- [LLVM Kaleidoscope Tutorial](https://llvm.org/docs/tutorial/MyFirstLanguageFrontend/index.html) - A hands-on tutorial for creating an LLVM-based language frontend.
- [TinyCC](https://bellard.org/tcc/) - A similar project I can get a lot of inspiration from.


[TJDevries](https://www.youtube.com/c/tjdevries) His channel on youtube and work on Neovim brought me to `LSP` and `tree-sitter`
