
# Keymaps

## Tmux

#### general

| key        | command     | description          |
|------------|-------------|----------------------|
| `<prefix>s`  |             | toggle status bar    |

#### session

| key        | command     | description          |
|------------|-------------|----------------------|
|            | tmux        | start tmux           |
| `<prefix>d`  |             | detach session       |
|            | tmux ls     | list session         |
|            | tmux a -t 0 | attach session 0     |

#### window

| key        | command     | description          |
|------------|-------------|----------------------|
| `<prefix>c`  |             | create new window    |
| `<prefix>,`  |             | rename window        |
| `<prefix>1`  |             | switch window 1      |
| `<prefix>w`  |             | open window list     |

#### pane

| key        | command     | description          |
|------------|-------------|----------------------|
| `<prefix>-`  |             | split pane below     |
| `<prefix>\|` |             | split pane right     |
| `<C-h>`      |             | go to left pane      |
| `<C-j>`      |             | go to lower pane     |
| `<C-k>`      |             | go to upper pane     |
| `<C-l>`      |             | go to right pane     |
| `<prefix>z`  |             | toggle pane maximize |
| `<prefix>x`  |             | close pane           |

#### copy mode

| key        | command     | description               |
|------------|-------------|---------------------------|
| `<prefix>[`  |             | enter copy mode           |
| `v`          |             | begin selection           |
| `<C-v>`      |             | toggle rectangle          |
| `y`          |             | copy selection and cancel |
| `<prefix>q`  |             | cancel                    |
| `<prefix>]`  |             | paste                     |

## LazyVim

### General

| key        | command | description        |
|------------|---------|--------------------|
| `<leader>uh` |         | toggle inlay hints |

### Debug

| key        | command | description          |
|------------|---------|----------------------|
| `<leader>da` |         | run with args        |
| `<leader>dl` |         | run last             |
| `<leader>db` |         | toggle breakpoint    |
| `<leader>dB` |         | breakpoint condition |
| `<leader>dc` |         | continue             |
| `<f8>`       |         | step over            |
| `<f7>`       |         | setp into            |
| `<leader>dt` |         | terminate            |

### Markdown

| key        | command | description            |
|------------|---------|------------------------|
| `<leader>um` |         | toggle markdown render |
| `<leader>cp` |         | markdown preview       |
