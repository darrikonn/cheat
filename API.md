# cheat
A fast and flexible cheatsheet manager built with
Go. Complete documentation is available at
https://github.com/darrikonn/cheat/API.md

```
Usage:
  cheat [regex] [command]

Available Commands:
  add         Add a new cheat
  delete      Delete a cheat
  edit        Edit a cheat
  get         Get cheat info
  *           Search cheats from regex (default)

Flags:
  -h, --help      help for cheat
      --verbose   verbose output
      --version   version for cheat

Use "cheat [command] --help" for more information about a command.
```

<hr />

## Search cheats from regex (default)
Search your cheats from a regex

```
Usage:
  cheat [regex] [flags]

Flags:
  -h, --help          help for search
  -i, --ignore-case   Case insensitive search

Global Flags:
      --verbose   verbose output
```

#### Examples
1. `cheat "git.*"`
2. `cheat`
3. `cheat GIT -i`

<hr />

## Add a new cheat
Add a new cheat to your cheatsheet. You'll be prompted for
the cheat's "description" in your preferred editor.

```
Usage:
  cheat [regex] add [flags]

Aliases:
  add, a

Flags:
  -h, --help         help for add
  -w, --weight int   weight of the cheat; used for sorting query results

Global Flags:
      --verbose   verbose output
```

#### Examples
1. `cheat git add`
2. `cheat grb a`
3. `cheat "git add -u" a --weight 2`

<hr />

## Delete a cheat
Delete a cheat from your cheatsheet.

```
Usage:
  cheat [regex] delete [flags]

Aliases:
  delete, d

Flags:
  -h, --help          help for delete
  -i, --ignore-case   Case insensitive search
  -y, --yes           Skip prompt

Global Flags:
      --verbose   verbose output
```

#### Examples
1. `cheat git delete`
2. `cheat grb d`
3. `cheat "GIT ADD -u" d --yes -i`

<hr />

## Edit a cheat
Edit a cheat's "name" and/or "weight". You'll also be
prompted for the cheat's "description" in your preferred editor.

```
Usage:
  cheat [regex] edit [flags]

Aliases:
  edit, e

Flags:
  -h, --help          help for edit
  -i, --ignore-case   Case insensitive search
  -n, --name string   Rename the cheat
  -w, --weight int    Weight of the cheat; used for sorting query results

Global Flags:
      --verbose   verbose output
```

#### Examples
1. `cheat git edit`
2. `cheat grb e --weight 2 --name "new name"`
3. `cheat "git ADD -u" e -i`

<hr />

## Get cheat info
Get a cheat info by name

```
Usage:
  cheat [regex] get [flags]

Aliases:
  get, g

Flags:
  -h, --help          help for get
  -i, --ignore-case   Case insensitive search

Global Flags:
      --verbose   verbose output
```

#### Examples
1. `cheat git get`
2. `cheat grb g`
3. `cheat "GIT ADD -u" g -i`
