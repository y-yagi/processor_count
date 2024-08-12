require "concurrent-ruby"
require "etc"

puts "Etc.nprocessors: #{Etc.nprocessors}"
puts "Concurrent.available_processor_count: #{Concurrent.available_processor_count}"
