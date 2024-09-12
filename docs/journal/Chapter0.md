# Chapter 0

## Intro

For quite a while I am interested in digging a bit under the hood of Computer Languages work. I have been reading books
on different things and explored some Computer Architecture. As I have some Hardware Background this was very natural
for me to explore.
How Compilers work and how the tooling around languges work is still quite some magic for me. Instead of only digging
into theory I want to get my hands dirty, because in the process I will be more interested and also receptive for theory
and concepts.

As I found the promising book [Writing a C Compiler](https://nostarch.com/writing-c-compiler) that seems to be my ideal
driver I wanted to start this journey.

I kind of made it into an ambitious project, because I not only want to build a C Compiler, but also an LLVM Frontend. I
want to implement both more or less in parallel to explore both approaches.
Also I want to iteratively build up other toolings around the langauge, as parsers, syntax-highlighting, pretty-printer
and language server.
Maybe I will dig a bit more into some of these, maybe I'll just glance over or ignore some of my ideas, but I think this
'holistic' approach will be a great way to explore and study more things.

As for implementation the book recommends a functional language, and as I started learning Haskell and am very
interested in exploring functional programming, this is my choice of implementation language.

For LLVM I want to use a different language and I thought of Rust to be a good candidate, this will be new for me, but
also something I want to explore for a while.

The language server will be written in Go, here I will follow the [video](https://www.youtube.com/watch?v=YsdlcQoHqPY)
from tjdevries and expand it on the go with my language features.

As I started already exploring some language tools - I am always curious about toolings and different approaches to
doing things, because they will open me either ideas that I can use in my existing tools, or a new tool to use for my
tasks.
So I was looking into 'flex' and also found a great use-case to utilize 'flex' for some project in my work.
I want to explore parsers and parser generators, here it will be 'bison' more for historic reasons and in combination
with 'flex'.
To get a more modern tool I want to look at 'antlr4'.
For aditional tooling and explore something I use daily, I also want to build 'tree-sitter' grammaer for my language.

I think this is a lot to pack into the project, but also I am very excited about all of it and really want to get this
going.

## Learnings and Take aways so far

I want to summarize the learnings I had so far.

I started exploring 'flex' with the [oreilly book](https://www.oreilly.com/library/view/flex-bison/9780596805418/) and
found the example using flex for recreating 'wc' quite fascinating.

In my work I want to fetch different patterns I define in regexes from html files, to extract metrics about features.
This is done for 'security research' in order to compare if some features in html sites are more likely to be available
in 'benign' or 'malicious' sites, these patterns are strongly inspired by 'indicators of compromise' and a great
resource is [ioc searcher](https://github.com/malicialab/iocsearcher). Instead of running the regexes with python (or
spark on bigger scale) I thought that I could build a 'lexer' utilizing `flex` to extract and count (with the help of a
hashmap) the indicators and compile the 'lexer' with python bindings, so I can send in my data and return 'numpy-arrays'
with my metrics.
My first draft version looked promising, but so far I did not create a full version, and I need to create some
benchmarks to compare and see what performance gain I get.

I think along the journey there will likely be a lot similar learnings and synergies showing up and allowing me to
integrate some techniques and tools into my skill-set.

## Outlooks

What do I hope to be able to do with these skills.

I am not sure if I ever will work on Compilers - you never know what your interest and focus will become at some later
point.
But I think understanding parsers, and tools you use in your daily work opens up many possibilties. You can easily
contribute to the tools you use if you need enhancements or find bugs.
You can build your own parsers, if there lacks a good possibility for your current need. You have the abilities to build
small DSLs in case you see the need for them.
You understand better how computer languages work and will become a better engineer.

And last but not least I think many of the concepts you learn along the journey will become helpful in many cases and
situations you encounter in your later career.
