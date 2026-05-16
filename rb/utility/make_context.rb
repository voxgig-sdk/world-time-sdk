# WorldTime SDK utility: make_context
require_relative '../core/context'
module WorldTimeUtilities
  MakeContext = ->(ctxmap, basectx) {
    WorldTimeContext.new(ctxmap, basectx)
  }
end
