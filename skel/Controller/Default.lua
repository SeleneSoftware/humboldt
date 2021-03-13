local c = require("controller")
local t = require("template")

c.setheader("jason","marshall")

t.setFile("default")

t.variable("title", "Humboldt Web Server Framework")

print(c.getheader("Accept"))
