# inaneshell

The silliest shell in the universe.

## Installation

The shell can be installed by using the command below.
```
go install github.com/kuwairo/inaneshell/cmd/inaneshell@latest
```

## Configuration

To configure the shell, create a file `/etc/inaneshell/config.json` with content
similar to the following.
```json
{
  "cd": true,
  "exit": "quit",
  "prompt": "% ",
  "provide": [
    "/usr/bin/dir",
    "/usr/bin/id",
    "/usr/bin/ping",
    "/usr/bin/pwd"
  ]
}
```

Then it is possible to set the shell as the default for the current user.
```
chsh -s /path/to/inaneshell
```
