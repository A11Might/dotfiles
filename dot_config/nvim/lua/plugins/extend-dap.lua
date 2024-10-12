return {
  "mfussenegger/nvim-dap",

  -- stylua: ignore
  keys = {
    { "<f7>", function() require("dap").step_into() end, desc = "Step Into" },
    { "<f8>", function() require("dap").step_over() end, desc = "Step Over" },
  },
}
