# dotfiles

1. Install [chezmoi](https://github.com/twpayne/chezmoi/)

```bash
brew install chezmoi
```

2. Update source directory

   2.1 Clone dofiles from GitHub into the source directory

   ```bash
   chezmoi init A11Might
   ```

   2.2 Or pull the lastest changes from remote repo

   ```bash
   chezmoi update
   ```

3. Check the changes that `chezmoi apply` would make to you home directory

```bash
chezmoi diff
```

4. Update local dotfiles

```bash
chezmoi apply
```

5. Enjoying it, my [keymaps](/keymaps.md)
