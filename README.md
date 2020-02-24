# labelmaker
Command-Line Interface tool for managing Github issue labels for repositories.

This project primarily came about from seeing other utilities out there but none conforming particularly well to *nix style conventions for CLI tools and also a desire for me to take a crack at using GraphQL. The Github API v4 nicely provides a GraphQL layer to play with for this tool.

## Sample Output

`$ GITHUB_TOKEN=NOTAREALTOKEN ./labelmaker list github/semantic`

```
Label			Color
bug			d73a4a
duplicate		cfd3d7
enhancement		a2eeef
help wanted		008672
good first issue	7057ff
invalid			e4e669
question		7faad8
wontfix			ffffff
SECURITY		e11d21
security-critical	eb6420
security-high		e11d21
security-medium		fbca04
security-low		fef2c0
security-informational	e6fca6
language: python	d497f4
infrastructure		83d8ef
diffing			bfdadc
performance: time	eb6420
libraries		fef2c0
maintenance		c2e0c6
performance: space	eb6420
documentation		ff33bb
assignment		b9ed36
language: ruby		d497f4
scope graph		27cea1
analysis		6adffc
language: go		d497f4
language: markdown	d497f4
language: typescript	d497f4
language: php		d497f4
language: java		d497f4
tests			ff9196
could not reproduce	bfdadc
blocked			ff3d97
backcompat		c5def5
language-support	d497f4
build			395dd3
compiler		fef2c0
blocker			eac641
ast:codegen		e99695
ast:marshal		5319e7
```
