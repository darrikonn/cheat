<p align="center">
    <img src="https://user-images.githubusercontent.com/5694851/70360689-1fafa100-1877-11ea-8c2a-92386220ae63.png" alt="Icon"/>
  <br />
  <br />
  <a href="https://github.com/darrikonn/cheat/releases/latest">
    <img src="https://img.shields.io/github/release/darrikonn/cheat.svg?style=flat-square"/>
  </a>
  <a href="https://github.com/goreleaser">
    <img src="https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square"/>
  </a>
  <a href="https://github.com/darrikonn/cheat/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/Licence-MIT-yellow.svg?longCache=true&style=flat-square"/>
  </a>
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Made With-Go-9cf.svg?longCache=true&style=flat-square"/>
  </a>
</h3>

<pre>
  <p align="center"><a href="https://pypi.org/project/cheat/"><strong>cheat</strong></a> is a command line cheat manager, <br/>where you can create and manage your personal cheatsheet</p>
  <p align="center"><img class="img-responsive" width="500" src="https://user-images.githubusercontent.com/5694851/70467469-2f202b80-1abd-11ea-9f29-0d52abfd09e9.gif" alt="gif"/></p>
  <p align="center"><a href="https://circleci.com/gh/darrikonn/cheat"><img src="https://circleci.com/gh/darrikonn/cheat.svg?style=svg" /></a></p>
</pre>


## Installation
Pre-built packages for Windows, macOS, and Linux are found on the [releases](https://github.com/darrikonn/cheat/releases) page.

Managed packages are in:
* **Homebrew** (*MacOs*)
  ```bash
  brew tap darrikonn/formulae
  brew install darrikonn/formulae/cheat
  ```
* **Scoop** (*Windows*)
  ```powerline
  scoop bucket add app https://github.com/darrikonn/cheat.git
  scoop install cheat
  ```
* **Other** (*Linux distros*)
  ```bash
  curl -s https://github.com/darrikonn/cheat/blob/master/install.sh | bash -s -- -b /usr/local/bin
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
### Database
By default the database will be located at `~/.cheetsheet.db`.
You can change your database path, and name, by specifying `database` in your `~/.cheat.yaml` file:
```yaml
database: ~/.custom-database-name.db
```
This results in a database instance at `~/.custom-database-name.db`

### Editor
When adding/editing a cheat, you'll be prompted to edit the cheat's `description` in your preferred editor. You can set your desired editor in the `~/.cheat.yaml` config file:
```yaml
editor: nvim
```
If no *editor* config is specified, the editor will fallback to your `EDITOR` environment variable. If that can't be found, the default selected editor will be `vi`.

## Tags
A neat way to search your cheats, is by describing them with tags.
```txt
my summary
tags: [awesome, golang]

my description
```
That way, you can simply search your cheats by tags, resulting in group like option for your cheats.
```bash
cheat 'tags: \[.*golang.*\]'
```
