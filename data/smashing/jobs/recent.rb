bundles = [
  {'label': 'azure-mysql', 'value': 'upgrade'},
  {'label': 'cnab-dashboard', 'value': 'install'},
  {'label': 'brigade', 'value': 'upgrade'},
  {'label': 'jenkins', 'value': 'upgrade'}, 
  {'label': 'jenkins-aci-connector', 'value': 'install'}, 
  {'label': 'quickstart', 'value': 'uninstall'}
] 

SCHEDULER.every '2s' do
  send_event('recent', { items: bundles })
end