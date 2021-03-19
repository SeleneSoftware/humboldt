local c = require("controller")
local t = require("template")

c.setheader("jason","marshall")

t.setFile("default")

t.variable("title", "This is a new page")
