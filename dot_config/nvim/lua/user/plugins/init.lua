return {
  ["goolord/alpha-nvim"] = { disable = true },
  ["folke/which-key.nvim"] = { disable = true },

  -- markdown preview
  {
    "iamcco/markdown-preview.nvim",
    run = function()
      vim.fn["mkdp#util#install"]()
    end,
    setup = function()
      vim.g.mkdp_filetypes = { "markdown" }
    end,
    ft = { "markdown" },
  },

  -- operate surroundings
  { "tpope/vim-surround" },

  -- add text-object 'a' (argument)
  { "vim-scripts/argtextobj.vim" }
}
