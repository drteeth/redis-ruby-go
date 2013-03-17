require 'rubygems'
require 'redis'
require 'json'

$redis = Redis.new(:timeout => 0)

$redis.subscribe('chanlol') do |on|
  on.message do |channel, msg|
    data = JSON.parse(msg)
    puts "##{channel} - [#{data['user']}]: #{data['msg']}"
  end
end
