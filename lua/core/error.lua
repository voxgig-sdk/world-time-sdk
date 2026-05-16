-- WorldTime SDK error

local WorldTimeError = {}
WorldTimeError.__index = WorldTimeError


function WorldTimeError.new(code, msg, ctx)
  local self = setmetatable({}, WorldTimeError)
  self.is_sdk_error = true
  self.sdk = "WorldTime"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function WorldTimeError:error()
  return self.msg
end


function WorldTimeError:__tostring()
  return self.msg
end


return WorldTimeError
