# Tea Ceremony with Dragons
or build a simple Compiler for studying.

I always wanted to drink tea with dragons, so I thought I will invite them for a ceremony!

## Why

For a while I want to dig deeper into the topic of Computer Languages and how they work.
I thought building a simple Compiler and exploring different tooling around is a great project to understand more and expand my knowledge.

As the main course I will follow the book [Writing a C Compiler](https://nostarch.com/writing-c-compiler) as my main Guide and add on and dig deeper with aditional resources on the go.
I dont want to follow a straight track, and make little side trips exploring different other aspects and toolings. I think it is quite ambitious, but the iterative approach on building the Compiler and taking the same steps for
the other tooling I think is a worthwhile effort.


## What

I will work on a subset of C, which I will call **`dragon-C`**. This subset will focus on key features of the C programming language, simplifying the full C language to make the compiler more manageable while still covering important concepts like variables, control structures, and functions.
The idea is to follow the book building a C-Compiler, and while at this explore different tools for lexing and parsing to also learn  nd explore a few more toolsets.
I also want to build a LLVM Frontend, so I have a excurse into another toolchain opening up possibilities.
I want to explore parsers a little and will look at `flex`, `bison`, `antlr4` and `tree-sitter`. With `tree-sitter` I want to explore building tooling for my editor (`neovim` but I guess its). And I want to build a 'language server'.

## How


The project is split into different parts.

`docs` will keep general documentation of the project as well as a 'journal' in which I will keep the progress and write down my experience and aditional thoughts or explorations.

`hcc` is the Compiler driver. It is a Python-based CLI tool that orchestrates different stages of the compilation process. Initially, it will use existing tools such as `gcc` or `clang` for preprocessing and assembling, but over time, it will increasingly rely on my own implementation of `dragoncc` for compilation.

`src` keeps all the source code I intend to write:
 - `dragoncc` - dragon c compiler, will be the compiler, implemented in haskell
 - `dragonclang` - a LLVM frontend implemented in Rust for the same subset of C
 - `experiments` - here I will explore tools like `flex`, `bison`, `antlr4` and `tree-sitter` and a 'language server' with the same langauge, to explore their usage.


 Maybe I will add:
 - A C repl - thus an interpreter - I will decide on the go.
 - Garbage Collection into Dragon-C - I will decide once the main project is done.


I choose 'Haskell' for the Compiler, as as a functional language with good pattern matchin capabilities it is well suited for Compilers and I really want to dig into that language, so this seems like a great project to deepen my understanding.

For the LLVM frontend I want to use another language, that I want to explore, 'Rust'.

The 'language server' will be written in 'Go'.


Along the way and following the Book as well as Tutorial and additional Resources I will explore the theory and concepts I think are interesting and summarize all in the Journal.
I hope this way my project might even be helpful for other people than me.

## Roadmap

- [x] Setup project scaffold & first docs and journal Chapter 0
- [ ] Write the first version of the `hcc` driver.
- [ ] Chapter 1: Minimal Compiler that handles integer constants
- [ ] Chapter 2: Add unary operations
- [ ] Chapter 3: Add binary operations
- [ ] Chapter 4: Add logical and relational operations
- [ ] Chapter 5: Local variables
- [ ] Chapter 6: conditionals
- [ ] Chapter 7: compound statements
- [ ] Chapter 8: loops
- [ ] Chapter 9: functions
- [ ] Chapter 10: file scope variable and storage-class specifiers
- [ ] Chapter 11: TBD
- [ ] Chapter ..

- [ ] Investigate adding garbage collection to Dragon-C
- [ ] Investigate adding a debugger

## Credits & Resources

Here I will add Books I use, as well as articles, repos or people that will turn out to be helpful or inspiring on the progress.

- [Writing a C Compiler](https://nostarch.com/writing-c-compiler) - as my main driver
- [The Dragon Book](https://en.wikipedia.org/wiki/Compilers:_Principles,_Techniques,_and_Tools) - To look up theory on the way
- [Language Implementation Patterns](https://pragprog.com/titles/tpdsl/language-implementation-patterns/) - As a secondary resource on my way
- [Flex & Bison](https://www.oreilly.com/library/view/flex-bison/9780596805418/) - To learn these tools
- [Crafting Interpreters](https://craftinginterpreters.com/) - as tertiary resource on my way
- [LLVM Documentation](https://llvm.org/docs/) - For understanding LLVM's intermediate representation (IR) and toolchain.
- [LLVM Kaleidoscope Tutorial](https://llvm.org/docs/tutorial/MyFirstLanguageFrontend/index.html) - A hands-on tutorial for creating an LLVM-based language frontend.
- [TinyCC](https://bellard.org/tcc/) - A similar project I can get a lot of inspiration from.
- [TJDevries](https://www.youtube.com/c/tjdevries) His channel on youtube and work on Neovim brought me to `LSP` and `tree-sitter`
- [TJDevries LSP Tutorial](https://www.youtube.com/watch?v=YsdlcQoHqPY) A Tutorial on implementing a LSP
- [Compiler Explorer](https://godbolt.org/) Interactive Website to explore Compilers and compare
- [Intel 64 Developer Manual](https://www.intel.com/content/www/us/en/developer/articles/technical/intel-sdm.html) - Official Intel docs
- [Intel Instruction Reference](https://www.felixcloutier.com/x86/) - Inofficial Reference for Intel Instruction
