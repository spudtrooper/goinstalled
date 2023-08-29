# goinstalled

Shows packages usd to install binaries in ~/go/bin

## Usage

### With `go install`

Install with

```bash
go install github.com/spudtrooper/goinstalled@latest
```

and then

```bash
goinstalled
```

### From source


```bash
go run main.go
```

## Example

Example output:

```
6 with NO matches
  dlv
  genopts
  gitversion
  goimports
  staticcheck
  uselocalgomodules

9 with ONE match
  anew -> github.com/tomnomnom/anew
  go-outline -> github.com/ramya-rao-a/go-outline
  gomodifytags -> github.com/fatih/gomodifytags
  goplay -> github.com/haya14busa/goplay
  gopls -> golang.org/x/tools/gopls
  gotests -> github.com/cweill/gotests
  goutil -> github.com/spudtrooper/goutil
  impl -> github.com/josharian/impl
  minimalcli -> github.com/spudtrooper/minimalcli

  INSTALL: go install github.com/tomnomnom/anew github.com/ramya-rao-a/go-outline github.com/fatih/gomodifytags github.com/haya14busa/goplay golang.org/x/tools/gopls github.com/cweill/gotests github.com/spudtrooper/goutil github.com/josharian/impl github.com/spudtrooper/minimalcli

0 with MULTIPLE matches
```