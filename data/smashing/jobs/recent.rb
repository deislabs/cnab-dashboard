require 'rest-client'
require 'json'

claimsapi = ENV['CLAIMS_API']

SCHEDULER.every '2s' do
  response = RestClient::Request.new(
     :method => :get,
     :url => "http://#{claimsapi}:3031/recent"
  ).execute
  results = JSON.parse(response.to_str)

  items = Array.new
  results.each do |i|
    items << { 'label': i['bundle'], 'value': i['action'] }
  end

  send_event('recent', { items: items })
end