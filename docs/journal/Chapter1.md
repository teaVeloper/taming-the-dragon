# Chapter 1

Now we start with the proper journey!

We want to start building a very simple toolset

## Requirements

### Compiler driver

My compiler driver is `hcc` which stands for 'Haskel C Compiler'
Basically it will be a python CLI and delegating subprocesses depending on the input.
I could use a simple bash-script, but I have a bit more fine grained controll when using python.

The driver should run
`gcc -E -P INPUT_FILE -o PREPROCESSED_FILE`
to run the preprocessor
then on the preprocessed file run my compiler, which for now will be just nothing
when compiling it creates an assembler file `output.s` and deletes `PREPROCESSED_FILE`
after it fill run
`agcc ASSEMBLY_FILE -o OUTPUT_FILE` and delete `ASSEMBLY_FILE`

The CLI shoud be
`hcc path/to/input.c` and create `path/to/input` as binary files and terminate with exit code `0`.
In case of an error, it should terminate with non-zero exit code.

It should expose the following options:
- `--lex` just run the lexer, stop before parsing
- `--parse` run lexer and parser, stop before assembly generation
- `--codegen` will run lexer, parser and code generation, but stop before code emission

These options should generate no output files (just dump to stdout?) and exit 0 on success.

- `-S` to generate an assembly `.s` file.

## Lexer

`int` keyword
`main` identifier
`(` open praenthesis
`void` keyword
`)` closing prenthesis
`{` open brace
`return` keyword
`2` constant
`;` semicolon
`}` clse brace


|   token    |       regex      |
| ---------- | ---------------- |
| identifier | `[a-zA-Z_]\w*\b` |
| constant   | `[0-9]+\b`       |
| int kw     | `int\b`          |
| void kw    | `void\b`         |
| return kw  | `return\b`       |
| (          | `\(`             |
| )          | `\)`             |
| {          | `{`              |
| }          | `}`              |
| ;          | `;`              |


pseudocode:

```
while input isn't empty:
    if input starts with whitespace:
        trim whitespace from start of input
    else:
        find longest match at start of input for any regex
        if no match if found, raise an error
        convert matching substring into a token
        remove matching substring from start of input
```
I will create a flex lexer and one in Haskell.

### Parser

Parser can be created by bison, antlr4 or hand (Haskell for me). I want to utilize all (and tree-sitter as well).

our 'ASDL'
```
program = Program(function_definition)
function_definition = Function(identifier name, statemend body)
statement = Return(exp)
exp = Constant(int)
```

Formal Grammar in 'EBNF'
```
<program> ::= <function>
<function> ::= "int" <identifier> "(" "void" ")" "{" <statement> "}"
<statement> ::= "return" <exp> ";"
<exp> ::= <int>
<identifier> ::= ? An identifier token ?
<int> ::= ? A constant token ?
```

Recursice Descent Parsing pseudocode:
```
parse_statement(tokens):
    expect("return", tokens)
    return_val = parse_exp(tokens)
    expect(";", tokens)
    return Return(return_val)

expect(expected, tokens):
    actual = take_token(tokens)
    if actual != expected:
        fail("Syntax error")
```

Pretty printed AST could look like
```
Program(
    Function(
        name="main",
        body=Return(
          Constant(2)
        )
    )
)
```

Add informative Error Messages!

### Assembly

ASDL for assembly
```
program = Program(function_definition)
function_definition = Function(identifier name, instruction* instructions)
instruction = Mov(operand src, operand dst) | Ret
operand = Imm(int) | Register
```


| AST Node | Assembly |
| -------- | -------- |
| Program(function_definition) |   Program(function_definition) |
| function(name, body) |   Function(name, instructions) |
|  Return(exp) |   Mov(exp, Register) ; Ret |
| Constant(int) | Imm(int) |


### Code emission

on Linux always add this to end of file
`.section .note.GNU-stack,"",@progbits`
This indicates the code does not need an [executable stack](https://www.airs.com/blog/archives/518)

Assembly output:

`Program(function_defintion)` ->
Print out the function definition (and add the line with non-executable stack at the end)

`Function(name, instruction)` ->
```
    .global <name>
<name>:
    <instructions>
```

Formatting
`Mov(src, dst)` -> `movl <src>, <dst>`
`Ret` -> `ret`
`Register` -> `%eax`
`Imm(int)` -> `$<int>`


