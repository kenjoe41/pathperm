# Pathperm
Pathperm is a tool for generating all possible permutations of a URL's path segments, with the option to add additional words from a supplied wordlist. This can be useful for security testing and identifying potential vulnerabilities in web applications. It can take input urls from standard input or from a file and a wordlist from a file. The permutations can then be written to a file or standard output.

## Installation
To install pathperm, use the following command:
```bash
go install github.com/kenjoe41/pathperm@latest
```

## Usage
```bash
pathperm [-d domain] [-w wordlist] [-l level] [-s] [-o output] [urls...]
```

## Options
```bash
Usage of pathperm:
  -d string
        The target domain if wordlist is supplied or if supplied paths have different domains.
  -i string
        Path to the input URLs file
  -l int
        level of permutation (shorthand) (default 3)
  -level int
        level of permutation (default 3)
  -o string
        Path to the output file
  -s    Silent mode
  -w string
        Path to the wordlist file
```

## Examples
```bash
pathperm -i urls.txt -w wordlist.txt -o output.txt

```
This command reads input urls from urls.txt, generates permutations using the words in wordlist.txt and writes the permuted paths to output.txt.
```bash
cat urls.txt | pathperm
```
This command reads input urls from standard input, generates permutations using the words from the paths of the provided urls only, and writes the permuted paths to standard output.
```bash
pathperm -i input.txt -w wordlist.txt -d example.com -o output.txt
```
This command reads input urls from urls.txt, generates permutations using the words in wordlist.txt, appends the domain example.com to the permuted paths and writes the permuted paths to output.txt

## Dependencies
* Golang 1.11 or later

## Contributing

1. Fork the repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## License

Pathperm is released under the MIT License. See [LICENSE](https://github.com/kenjoe41/pathperm/blob/main/LICENSE) for more details.