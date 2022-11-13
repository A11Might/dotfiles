return {
	n = {
		["<leader>bb"] = { "<cmd>tabnew<cr>", desc = "New tab" },
		["<leader>bc"] = { "<cmd>BufferLinePickClose<cr>", desc = "Pick to close" },
		["<leader>bj"] = { "<cmd>BufferLinePick<cr>", desc = "Pick to jump" },
		-- conflict
		-- ["K"] = { "5k", desc = "Move up faster" },
		-- ["J"] = { "5j", desc = "Move down faster" },
		-- ["H"] = { "5h", desc = "Move left faster" },
		-- ["L"] = { "5l", desc = "Move right faster" },
		-- split
		-- ["s"] = { "<nop>", desc = "Disable s" },
		-- ["sk"] = { "<cmd>set nosplitbelow<cr><cmd>split<cr>", desc = "Open above horizontal split" },
		-- ["sj"] = { "<cmd>set splitbelow<cr><cmd>split<cr>", desc = "Open below horizontal split" },
		-- ["sh"] = { "<cmd>set nosplitright<cr><cmd>vsplit<cr>", desc = "Open left vertical split" },
		-- ["sl"] = { "<cmd>set splitright<cr><cmd>vsplit<cr>", desc = "Open right vertical split" },
		-- ["sV"] = { "<C-w>t<C-w>H", desc = "Switch horizontal split to viertical" },
		-- ["sH"] = { "<C-w>t<C-w>K", desc = "Switch vertical split to horizontal" },
		-- ["<leader>k"] = { "<C-w>k", desc = "Move cursor to above" },
		-- ["<leader>j"] = { "<C-w>j", desc = "Move cursor to below" },
		-- ["<leader>h"] = { "<C-w>h", desc = "Move cursor to left" },
		-- ["<leader>l"] = { "<C-w>l", desc = "Move cursor to right" },
		-- ["<up>"] = { "<cmd>:resize+5<cr>", desc = "Increase window height" },
		-- ["<down>"] = { "<cmd>:resize-5<cr>", desc = "Decrease window height" },
		-- ["<left>"] = { "<cmd>:vertical resize+5<cr>", desc = "Increase window width" },
		-- ["<right>"] = { "<cmd>:vertical resize-5<cr>", desc = "Decrease window width" },
		-- -- tab
		-- ["tn"] = { "<cmd>tabe<cr>", desc = "Open new tab" },
		-- ["th"] = { "<cmd>-tabnext<cr>", desc = "Move to left tab" },
		-- ["tl"] = { "<cmd>+tabnext<cr>", desc = "Move to right tab" },
		-- file
		["S"] = { "<cmd>w<cr>", desc = "Save file" },
		["Q"] = { "<cmd>q<cr>", desc = "Quit file" },
		-- search
		["<leader>sl"] = { "<cmd>nohlsearch<cr>", desc = "Search highlight clear" },
		-- spell
		-- [s/]s: preview/next spelling mistake
		-- z=/<C-x>s: normal/insert mode correction list
		["<leader>sk"] = { "<cmd>set spell!<cr>", desc = "Toggle spell check" },
		-- NeoTree
		-- ["ff"] = { "<cmd>Neotree toggle<cr>", desc = "Toggle Explorer" },
		-- ["ff"] = { "<cmd>Neotree focus<cr>", desc = "Focus Explorer" },
		-- call figlet
		["tx"] = { ":r !figlet ", desc = "Focus Explorer" },
	},
	i = {
		-- other useful stuff
		["<leader><leader>"] = { "<esc>/<++><cr><cmd>nohlsearch<cr>c4l", desc = "Jump to next '<++>' and edit it" },
	}
}
