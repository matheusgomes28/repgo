-- Assumption: the input is a stripped string, args
-- are separated by a space

-- Assumption: each argument is a number that is convertible
-- to an integer

local function split_spaces(arguments)
    local a = {}
    for arg, what_is_this in string.gmatch(arguments, "%S+") do
        table.insert(a, arg)
    end

    return a
end

-- add all elements of array `a'
function Main(a)
    local split = split_spaces(a)

    -- ensure that they have size 2
    -- do this on golang size
    local sum = split[1] - split[2]
    return tostring(sum)
end
