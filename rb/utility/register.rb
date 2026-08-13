# WorldTime SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

WorldTimeUtility.registrar = ->(u) {
  u.clean = WorldTimeUtilities::Clean
  u.done = WorldTimeUtilities::Done
  u.make_error = WorldTimeUtilities::MakeError
  u.feature_add = WorldTimeUtilities::FeatureAdd
  u.feature_hook = WorldTimeUtilities::FeatureHook
  u.feature_init = WorldTimeUtilities::FeatureInit
  u.fetcher = WorldTimeUtilities::Fetcher
  u.make_fetch_def = WorldTimeUtilities::MakeFetchDef
  u.make_context = WorldTimeUtilities::MakeContext
  u.make_options = WorldTimeUtilities::MakeOptions
  u.make_request = WorldTimeUtilities::MakeRequest
  u.make_response = WorldTimeUtilities::MakeResponse
  u.make_result = WorldTimeUtilities::MakeResult
  u.make_point = WorldTimeUtilities::MakePoint
  u.make_spec = WorldTimeUtilities::MakeSpec
  u.make_url = WorldTimeUtilities::MakeUrl
  u.param = WorldTimeUtilities::Param
  u.prepare_auth = WorldTimeUtilities::PrepareAuth
  u.prepare_body = WorldTimeUtilities::PrepareBody
  u.prepare_headers = WorldTimeUtilities::PrepareHeaders
  u.prepare_method = WorldTimeUtilities::PrepareMethod
  u.prepare_params = WorldTimeUtilities::PrepareParams
  u.prepare_path = WorldTimeUtilities::PreparePath
  u.prepare_query = WorldTimeUtilities::PrepareQuery
  u.graphql_body = WorldTimeUtilities::GraphqlBody
  u.graphql_errors = WorldTimeUtilities::GraphqlErrors
  u.result_basic = WorldTimeUtilities::ResultBasic
  u.result_body = WorldTimeUtilities::ResultBody
  u.result_headers = WorldTimeUtilities::ResultHeaders
  u.transform_request = WorldTimeUtilities::TransformRequest
  u.transform_response = WorldTimeUtilities::TransformResponse
}
