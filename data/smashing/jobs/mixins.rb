mixins = ['exec', 'helm', 'terraform', 'azure', 'aws', 'kubernetes'] 
mixin_counts = Hash.new({ value: 0 })

SCHEDULER.every '2s' do
  random_mixin = mixins.sample
  mixin_counts[random_mixin] = { label: random_mixin, value: (mixin_counts[random_mixin][:value] + 1) % 30 }
  
  send_event('mixins', { items: mixin_counts.values })
end