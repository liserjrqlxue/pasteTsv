# pasteTsv
`paste` tsv but omit header columns

The go implement of `GNU coreutils paste` with feature:
1. omit first `-omit` columns (split by `-omitSep`) in case that paste files with same header columns
2. `-sep` implement `-d`/`--delimiters` of `paste` 
3. work line by line and low memory usage