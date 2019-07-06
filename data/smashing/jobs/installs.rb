require 'rest-client'
require 'json'

claimsapi = ENV['CLAIMS_API']

SCHEDULER.every '2s' do
  response = RestClient::Request.new(
     :method => :get,
     :url => "http://#{claimsapi}:3031/installs"
  ).execute
  results = JSON.parse(response.to_str)
  installedBundles = results['installedBundles']
  points = results['installHistory']

  send_event('installs', points: points, displayedValue: installedBundles)
end