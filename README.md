## blog

   

used_memory_human:447.19M
used_memory_rss_human:1.03G
mem_fragmentation_ratio:2.41

vim.highlight.create is deprecated, use vim.api.nvim_set_hl instead. See :h deprecated
This function will be removed in Nvim version 0.9
stack traceback:
        ...ar/neovim/0.8.1/share/nvim/runtime/lua/vim/highlight.lua:14: in function 'create'
        ...k/packer/start/toggleterm.nvim/lua/toggleterm/colors.lua:99: in function 'set_highlights'
        ...ite/pack/packer/start/toggleterm.nvim/lua/toggleterm.lua:359: in function 'setup'
        /Users/xsh/.config/nvim/lua/plugin-config/toggleterm.lua:8: in main chunk
        [C]: in function 'require'
        /Users/xsh/.config/nvim/init.lua:20: in main chunk
cmp_nvim_lsp.update_capabilities is deprecated, use cmp_nvim_lsp.default_capabilities instead. See :h
 deprecated
This function will be removed in cmp-nvim-lsp version 1.0.0
stack traceback:
        ...pack/packer/start/cmp-nvim-lsp/lua/cmp_nvim_lsp/init.lua:89: in function 'update_capabilit
ies'
        /Users/xsh/.config/nvim/lua/lsp/config/ts.lua:7: in main chunk
        [C]: in function 'require'
        /Users/xsh/.config/nvim/lua/lsp/setup.lua:20: in main chunk
        [C]: in function 'require'
        /Users/xsh/.config/nvim/init.lua:26: in main chunk
[nvim-lsp-installer] (automatic installation) Installing LSP server: emmet_ls
[nvim-lsp-installer] (automatic installation) Installing LSP server: rust_analyzer
sumneko_lua is deprecated, use lua_ls instead. See :h deprecated
This function will be removed in lspconfig version 0.2.0
stack traceback:
        .../site/pack/packer/start/nvim-lspconfig/lua/lspconfig.lua:36: in function '__index'
        /Users/xsh/.config/nvim/lua/lsp/setup.lua:29: in main chunk
        [C]: in function 'require'
        /Users/xsh/.config/nvim/init.lua:26: in main chunk
'lua-dev' was renamed to 'neodev'. Please update your config.
[nvim-lsp-installer] (automatic installation) Installing LSP server: yamlls
[nvim-lsp-installer] (automatic installation) Installing LSP server: cssls
You're using the old way of setting up neodev (previously lua-dev).
Please check the docs at https://github.com/folke/neodev.nvim#-setup
Spawning language server with cmd: `lua-language-server` failed. The language server is either not in
stalled, missing from PATH, or not executable.

executing vim.schedule lua callback: ...w/Cellar/neovim/0.8.1/share/nvim/runtime/lua/vim/lsp.lu
a:1506: attempt to call field 'is_closing' (a nil value)
stack traceback:
        ...w/Cellar/neovim/0.8.1/share/nvim/runtime/lua/vim/lsp.lua:1506: in function 'is_stopped'
        ...w/Cellar/neovim/0.8.1/share/nvim/runtime/lua/vim/lsp.lua:598: in function 'send_changes'
        ...w/Cellar/neovim/0.8.1/share/nvim/runtime/lua/vim/lsp.lua:655: in function ''