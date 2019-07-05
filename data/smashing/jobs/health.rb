fails = 0
pendings = 0

SCHEDULER.every '2s' do
  fails = rand(2)
  pendings  = rand(2)

  send_event('health', { fails: fails, pendings: pendings })
end