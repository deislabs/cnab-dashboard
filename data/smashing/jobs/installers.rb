installers = ['porter', 'duffle', 'docker app'] 
installer_counts = Hash.new({ value: 0 })

SCHEDULER.every '2s' do
  random_installer = installers.sample
  installer_counts[random_installer] = { label: random_installer, value: (installer_counts[random_installer][:value] + 1) % 30 }
  
  send_event('installers', { items: installer_counts.values })
end