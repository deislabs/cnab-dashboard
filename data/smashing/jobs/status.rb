require 'rest-client'
require 'json'

claimsapi = ENV['CLAIMS_API']

SCHEDULER.every '2s' do
  response = RestClient::Request.new(
     :method => :get,
     :url => "http://#{claimsapi}:3031/status"
  ).execute
  results = JSON.parse(response.to_str)
  failed = results['failed']
  pending = results['pending']

  send_event('status', { failed: failed, pending: pending })
end
