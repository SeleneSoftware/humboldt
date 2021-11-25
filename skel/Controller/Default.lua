local c = require("controller")
local t = require("template")
local f = require("forms")

c.setheader("jason","marshall")

t.setFile("default")

t.variable("title", "Humboldt Web Server Framework")

print(c.getheader("Accept"))

-- p = form.new("Steeve")
-- print(p:name()) -- "Steeve"
-- element = form.element("Jason")
-- g = form.new("Alice")
-- print(g) -- "Alice"

emailField = {
    Name = "email",
    Type = "text",
    Value = "",
}
form = {
    name = "Jason",
    elements = {
        emailField,
        {
            Name = "first",
            Type = "text",
            Value = "",
        }
    }
}

t.variable("form",form)

-- t.variable("form",f.render(form))

-- getrequestheader("Jason")
