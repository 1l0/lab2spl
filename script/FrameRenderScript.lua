local num = CustomTool1.NumberIn1
if num == 0 then
    self.StyledText = "" -- pau
    self:setData("prevText", self.StyledText)
elseif num == 1 then
    self.StyledText = "" -- N
    self:setData("prevText", self.StyledText)
elseif num == 2 then
    self.StyledText = "o" -- a
    self:setData("prevText", self.StyledText)
elseif num == 3 then
    self.StyledText = "" -- cl
    self:setData("prevText", self.StyledText)
elseif num == 4 then
    self.StyledText = "=" -- e
    self:setData("prevText", self.StyledText)
elseif num == 5 then
    self.StyledText = "-" -- i
    self:setData("prevText", self.StyledText)
elseif num == 6 then
    self.StyledText = "O" -- o
    self:setData("prevText", self.StyledText)
elseif num == 7 then
    self.StyledText = "•" -- u
    self:setData("prevText", self.StyledText)
elseif num == 8 then
    self.StyledText = "o" -- A
    self:setData("prevText", self.StyledText)
elseif num == 9 then
    self.StyledText = "=" -- E
    self:setData("prevText", self.StyledText)
elseif num == 10 then
    self.StyledText = "-" -- I
    self:setData("prevText", self.StyledText)
elseif num == 11 then
    self.StyledText = "O" -- O
    self:setData("prevText", self.StyledText)
elseif num == 12 then
    self.StyledText = "•" -- U
    self:setData("prevText", self.StyledText)
else
    self.StyledText = self:getData("prevText")
end
