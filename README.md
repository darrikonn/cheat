<p align="center">
    <img src="https://user-images.githubusercontent.com/5694851/70360689-1fafa100-1877-11ea-8c2a-92386220ae63.png" alt="Icon"/>
  <br />
  <br />
  <a href="https://pypi.org/project/cheat/">
    <img src="https://img.shields.io/pypi/v/cheat.svg?style=flat-square"/>
  </a>
  <a href="https://pypi.org/project/cheat/">
    <img src="https://img.shields.io/pypi/dm/cheat?style=flat-square"/>
  </a>
  <a href="https://github.com/darrikonn/cheat/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/Licence-MIT-yellow.svg?longCache=true&style=flat-square"/>
  </a>
  <a href="https://www.python.org/">
    <img src="https://img.shields.io/badge/Made With-Go-red.svg?longCache=true&style=flat-square"/>
  </a>
</h3>

<pre>
  <p align="center"><a href="https://pypi.org/project/cheat/"><strong>cheat</strong></a> is a command line cheat manager, <br/>where you can create and manage your personal cheatsheet</p>
  <p align="center"><img class="img-responsive" width="500" src="https://raw.githubusercontent.com/darrikonn/cheat/master/img/cheat.gif" alt="gif"/></p>
  <p align="center"><a href="https://circleci.com/gh/darrikonn/cheat"><img src="https://circleci.com/gh/darrikonn/cheat.svg?style=svg" /></a></p>
</pre>


## Installation
[**cheat**](https://pypi.org/project/cheat/) works on all OSs, so it needs to be installed with the package manager that suits your OS.
```bash
brew install cheat
```

## Getting started
Run `cheat --help` to see possible commands.

Here are some to get you started:
- Run `cheat` to list all your cheats.

- Run `cheat some.*regex` to cheats matching your regex.

- Run `cheat some.*regex add` to add a new cheat.


## API
Check out the [`api`](https://github.com/darrikonn/cheat/blob/master/API.md).

## Configuring
### Database name
Your database instance will be located in your home directory (`~/`).
By default it'll be named `.cheetsheet.db`.

You can change your database name by specifying `database` in your `~/.cheat.yaml` file:
```yaml
database: ~/.custom-database-name.db
```
This results in a database instance at `~/.custom-database-name.db`

### Editor
When adding/editing a cheat, you'll be prompted to edit the cheat's `description` in your preferred editor. You can set you preferred editor in the `~/.cheat.yaml` config file:
```yaml
editor: nvim
```
If no *editor* config is specified, the editor will fallback to your `EDITOR` environment variable. If that can't be found, the default selected editor will be `vi`.

## Tags
A need way to search your cheats, is by describing them with tags.
```txt
my summary
tags: [awesome, golang]

my description
```
That way, you can simply search your cheats by tags, resulting in group like option for your cheats.
```bash
cheat 'tags: \[.*golang.*\]'
```
