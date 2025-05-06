local num = CustomTool1.NumberIn1
if num == 0 then
    self.StyledText = "-" -- pau
    setData("prevText", self.StyledText)
elseif num == 1 then
    self.StyledText = "-" -- N
    setData("prevText", self.StyledText)
elseif num == 2 then
    self.StyledText = "□" -- a
    setData("prevText", self.StyledText)
elseif num == 3 then
    self.StyledText = "-" -- cl
    setData("prevText", self.StyledText)
elseif num == 4 then
    self.StyledText = "ㅂ" -- e
    setData("prevText", self.StyledText)
elseif num == 5 then
    self.StyledText = "ㅂ" -- i
    setData("prevText", self.StyledText)
elseif num == 6 then
    self.StyledText = "ﾛ" -- o
    setData("prevText", self.StyledText)
elseif num == 7 then
    self.StyledText = "ε" -- u
    setData("prevText", self.StyledText)
elseif num == 8 then
    self.StyledText = "□" -- A
    setData("prevText", self.StyledText)
elseif num == 9 then
    self.StyledText = "ㅂ" -- E
    setData("prevText", self.StyledText)
elseif num == 10 then
    self.StyledText = "ㅂ" -- I
    setData("prevText", self.StyledText)
elseif num == 11 then
    self.StyledText = "ﾛ" -- O
    setData("prevText", self.StyledText)
elseif num == 12 then
    self.StyledText = "ε" -- U
    setData("prevText", self.StyledText)
else
    self.StyledText = getData("prevText")
end
