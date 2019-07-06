class Dashing.Health extends Dashing.Widget

  ready: ->
    if not @get('status')
      @set 'error', 'No data'
      @set 'status', 'error'

  onData: (data) ->

    status = switch
      when not (data.hasOwnProperty('pending') and data.hasOwnProperty('failed')) then 'error'
      when data.hasOwnProperty('error') then 'error'
      when @get('failed') > 0 then 'red'
      when @get('pending') > 0 then 'yellow'
      else 'green'

    @set 'status', status

    if status is 'error'
      if not data.hasOwnProperty('error')
        # Error condition because of a missing field
        @set 'error', 'Data provided without "failed" and "pending" fields.'
