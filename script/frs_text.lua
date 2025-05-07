local num = CustomTool1.NumberIn1
if num == 0 then
    self.StyledText = "-" -- pau
    self:SetData("prevText", self.StyledText)
    return
elseif num == 1 then
    self.StyledText = "-" -- N
    self:SetData("prevText", self.StyledText)
    return
elseif num == 2 then
    self.StyledText = "o" -- a
    self:SetData("prevText", self.StyledText)
    return
elseif num == 3 then
    self.StyledText = "-" -- cl
    self:SetData("prevText", self.StyledText)
    return
elseif num == 4 then
    self.StyledText = "=" -- e
    self:SetData("prevText", self.StyledText)
    return
elseif num == 5 then
    self.StyledText = "--" -- i
    self:SetData("prevText", self.StyledText)
    return
elseif num == 6 then
    self.StyledText = "O" -- o
    self:SetData("prevText", self.StyledText)
    return
elseif num == 7 then
    self.StyledText = "•" -- u
    self:SetData("prevText", self.StyledText)
    return
elseif num == 8 then
    self.StyledText = "o" -- A
    self:SetData("prevText", self.StyledText)
    return
elseif num == 9 then
    self.StyledText = "=" -- E
    self:SetData("prevText", self.StyledText)
    return
elseif num == 10 then
    self.StyledText = "--" -- I
    self:SetData("prevText", self.StyledText)
    return
elseif num == 11 then
    self.StyledText = "O" -- O
    self:SetData("prevText", self.StyledText)
    return
elseif num == 12 then
    self.StyledText = "•" -- U
    self:SetData("prevText", self.StyledText)
    return
end
print(num)
