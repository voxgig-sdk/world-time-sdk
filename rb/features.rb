# WorldTime SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module WorldTimeFeatures
  def self.make_feature(name)
    case name
    when "base"
      WorldTimeBaseFeature.new
    when "test"
      WorldTimeTestFeature.new
    else
      WorldTimeBaseFeature.new
    end
  end
end
