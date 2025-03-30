#!/usr/bin/env ruby
require 'faker'
require 'securerandom'
require 'optparse'
require 'date'

# Color codes for terminal output
class String
  def red; "\e[31m#{self}\e[0m" end
  def green; "\e[32m#{self}\e[0m" end
  def yellow; "\e[33m#{self}\e[0m" end
  def blue; "\e[34m#{self}\e[0m" end
  def magenta; "\e[35m#{self}\e[0m" end
  def cyan; "\e[36m#{self}\e[0m" end
  def bold; "\e[1m#{self}\e[22m" end
end

# Configuration options
options = {
  size: :small,
  output: nil, # Will be set based on size
  verbose: false
}

# Parse command line arguments
OptionParser.new do |opts|
  opts.banner = "Usage: generate_bank_data.rb [options]"

  opts.on("-s", "--size SIZE", [:small, :medium, :large], 
          "Data size (small, medium, large)") do |s|
    options[:size] = s
  end

  opts.on("-o", "--output FILE", "Output SQL file") do |file|
    options[:output] = file
  end

  opts.on("-v", "--verbose", "Show verbose output") do
    options[:verbose] = true
  end

  opts.on("-h", "--help", "Prints this help") do
    puts opts
    exit
  end
end.parse!

# Set default output filename if not specified
options[:output] ||= "./sql/bank_data_#{options[:size]}.sql"

# Define data sizes
DATA_SIZES = {
  small: {
    branches: 5,
    customers: 20,
    accounts_per_customer: (1..2),
    transactions_per_account: (3..10),
    loan_percentage: 0.1
  },
  medium: {
    branches: 20,
    customers: 100,
    accounts_per_customer: (1..3),
    transactions_per_account: (5..20),
    loan_percentage: 0.2
  },
  large: {
    branches: 50,
    customers: 500,
    accounts_per_customer: (1..4),
    transactions_per_account: (10..30),
    loan_percentage: 0.25
  }
}

size_config = DATA_SIZES[options[:size]]

puts "Generating #{options[:size].to_s.upcase} test dataset...".bold.cyan
puts "Configuration:".bold
puts "  Branches: #{size_config[:branches]}".green
puts "  Customers: #{size_config[:customers]}".green
puts "  Output file: #{options[:output]}".green

# Initialize SQL output
sql_script = []
sql_script << "-- Bank Database Test Data"
sql_script << "-- Generated: #{Time.now}"
sql_script << "-- Size: #{options[:size].to_s.upcase}"
sql_script << ""
sql_script << "SET FOREIGN_KEY_CHECKS = 0;"
sql_script << ""

# Helper methods
def log(message, level = :info, options)
  return unless options[:verbose] || level == :important
  
  prefix = case level
           when :error then "[ERROR]".red
           when :warn then "[WARN]".yellow
           when :info then "[INFO]".blue
           when :important then "[IMPORTANT]".bold.green
           end
  
  puts "#{prefix} #{message}"
end

def sql_escape(str)
  str.to_s.gsub("'", "''")
end

# Generate branches (multi-row INSERT)
log("Generating branches...", :important, options)
branch_ids = []
branch_values = []
size_config[:branches].times do |i|
  branch_name = Faker::Bank.unique.name
  branch_location = "#{Faker::Address.street_address}, #{Faker::Address.city}"
  branch_values << "('#{sql_escape(branch_name)}', '#{sql_escape(branch_location)}')"
  branch_ids << i + 1
end
sql_script << "INSERT INTO branches (branch_name, branch_location) VALUES #{branch_values.join(", ")};"
sql_script << ""

# Generate customers (multi-row INSERT)
log("Generating customers...", :important, options)
customer_ids = []
customer_values = []
size_config[:customers].times do |i|
  first_name = Faker::Name.first_name
  last_name = Faker::Name.last_name
  city = Faker::Address.city
  mobile_no = Faker::PhoneNumber.cell_phone.gsub(/\D/, '')[0...20]
  idcard_no = SecureRandom.random_number(10**18).to_s.rjust(18, '0')
  dob = Faker::Date.birthday(min_age: 18, max_age: 80)
  customer_values << "('#{sql_escape(first_name)}', '#{sql_escape(last_name)}', '#{sql_escape(city)}', '#{mobile_no}', '#{idcard_no}', '#{dob}')"
  customer_ids << i + 1
end
sql_script << "INSERT INTO customers (first_name, last_name, city, mobile_no, IDcard_no, dob) VALUES #{customer_values.join(", ")};"
sql_script << ""

# Generate accounts (multi-row INSERT)
log("Generating accounts...", :important, options)
account_ids = []
account_values = []
account_counter = 1
customer_ids.each do |customer_id|
  rand(size_config[:accounts_per_customer]).times do
    account_type = ['savings account', 'current account', 'fixed deposit'].sample
    currency = ['cny', 'usd', 'eur'].sample
    status = ['Inactive', 'Active', 'Closed'].sample
    balance = Faker::Number.decimal(l_digits: rand(1..4), r_digits: 2)
    account_values << "(#{customer_id}, #{balance}, '#{status}', '#{account_type}', '#{currency}')"
    account_ids << account_counter
    account_counter += 1
  end
end
sql_script << "INSERT INTO accounts (customer_id, balance, account_status, account_type, currency) VALUES #{account_values.join(", ")};"
sql_script << ""

# Generate transactions (multi-row INSERT)
log("Generating transactions...", :important, options)
transaction_values = []
account_ids.each do |account_id|
  rand(size_config[:transactions_per_account]).times do
    transaction_type = ['deposit', 'withdrawal', 'transfer'].sample
    medium = ['ATM', 'mobile', 'computer', 'manual service', 'other'].sample
    status = ['pending', 'failed', 'success'].sample
    amount = Faker::Number.decimal(l_digits: rand(1..4), r_digits: 2)
    date_issued = Faker::Date.between(from: Date.today - 365, to: Date.today)
    reference = SecureRandom.random_number(10**10)
    to_account = transaction_type == 'transfer' ? account_ids.sample : 'NULL'
    transaction_values << "('#{transaction_type}', #{account_id}, #{to_account}, '#{date_issued}', #{amount}, '#{medium}', '#{status}', #{reference})"
  end
end
# Split into chunks to avoid too large SQL statements
transaction_values.each_slice(1000) do |chunk|
  sql_script << "INSERT INTO transactions (transaction_type, from_account_id, to_account_id, date_issued, amount, transaction_medium, status, reference) VALUES #{chunk.join(", ")};"
end
sql_script << ""

# Generate loans (multi-row INSERT)
log("Generating loans...", :important, options)
loan_values = []
customer_ids.sample((customer_ids.size * size_config[:loan_percentage]).to_i).each do |customer_id|
  loan_amount = Faker::Number.decimal(l_digits: rand(3..5), r_digits: 2)
  date_issued = Faker::Date.between(from: Date.today - (3*365), to: Date.today)
  branch_id = branch_ids.sample
  loan_values << "(#{customer_id}, #{branch_id}, #{loan_amount}, '#{date_issued}')"
end
sql_script << "INSERT INTO loans (customer_id, branch_id, loan_amount, date_issued) VALUES #{loan_values.join(", ")};"

# Finalize SQL script
sql_script << ""
sql_script << "SET FOREIGN_KEY_CHECKS = 1;"
sql_script << ""
sql_script << "-- Data generation complete"

# Write to file
File.open(options[:output], 'w') { |f| f.write(sql_script.join("\n")) }

# Summary
puts "\nData generation complete!".bold.green
puts "Summary:".bold
puts "  Branches generated: #{size_config[:branches]}".green
puts "  Customers generated: #{size_config[:customers]}".green
puts "  Accounts generated: #{account_ids.size}".green
puts "  Transactions generated: #{transaction_values.size}".green
puts "  Loans generated: #{loan_values.size}".green
puts "\nSQL script written to: #{options[:output].bold}"