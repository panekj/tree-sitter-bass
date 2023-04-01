# tree-sitter-bass

This crate provides a Bass grammar for the [tree-sitter][] parsing library. To
use this crate, add it to the `[dependencies]` section of your `Cargo.toml`
file. (Note that you will probably also need to depend on the
[`tree-sitter`][tree-sitter crate] crate to use the parsed result in any useful
way.)

```toml
[dependencies]
tree-sitter = "~0.20.3"
tree-sitter-bass = "0.0.2"
```

Typically, you will use the [language][language func] function to add this
grammar to a tree-sitter [Parser][], and then use the parser to parse some code:

```rust
let code = r#"
(provide [git-submodules]
  (defn git-submodules [src]
    (map (fn [[_ path]] (string->dir path))
      (-> ($ git config --file src/.gitmodules --get-regexp path)
          (with-image (linux/alpine/git))
          (read :unix-table)
          take-all))))
"#;
let mut parser = Parser::new();
parser.set_language(tree_sitter_bass::language()).expect("Error loading Bass grammar");
let parsed = parser.parse(code, None);
```

If you have any questions, please reach out to us in the [tree-sitter
discussions] page.

[language func]: https://docs.rs/tree-sitter-bass/*/tree_sitter_bass/fn.language.html
[parser]: https://docs.rs/tree-sitter/*/tree_sitter/struct.Parser.html
[tree-sitter]: https://tree-sitter.github.io/
[tree-sitter crate]: https://crates.io/crates/tree-sitter
[tree-sitter discussions]: https://github.com/tree-sitter/tree-sitter/discussions
