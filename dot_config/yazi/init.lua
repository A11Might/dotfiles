-- https://github.com/yazi-rs/plugins/tree/main/git.yazi
th.git = th.git or {}

th.git.added = ui.Style():fg("green")
th.git.added_sign = "A"

th.git.deleted = ui.Style():fg("red"):bold()
th.git.deleted_sign = "D"

th.git.ignored = ui.Style():fg("gray")
th.git.ignored_sign = "!!"

th.git.modified = ui.Style():fg("blue")
th.git.modified_sign = "M"

th.git.untracked = ui.Style():fg("yellow")
th.git.untracked_sign = "??"

th.git.updated = ui.Style():fg("green")
th.git.updated_sign = "U"

require("git"):setup()

-- https://github.com/Rolv-Apneseth/starship.yazi
require("starship"):setup()
