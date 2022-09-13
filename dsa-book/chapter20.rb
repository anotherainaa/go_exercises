authors = [
  {"author_id" => 1, "name" => "Virginia Woolf"}, 
  {"author_id" => 2, "name" => "Leo Tolstoy"}, 
  {"author_id" => 3, "name" => "Dr. Seuss"}, 
  {"author_id" => 4, "name" => "J. K. Rowling"}, 
  {"author_id" => 5, "name" => "Mark Twain"}
]

books = [
  {"author_id" => 3, "title" => "Hop on Pop"},
  {"author_id" => 1, "title" => "Mrs. Dalloway"},
  {"author_id" => 4, "title" => "Harry Potter and the Sorcerer's Stone"}, {"author_id" => 1, "title" => "To the Lighthouse"},
  {"author_id" => 2, "title" => "Anna Karenina"},
  {"author_id" => 5, "title" => "The Adventures of Tom Sawyer"}, {"author_id" => 3, "title" => "The Cat in the Hat"},
  {"author_id" => 2, "title" => "War and Peace"},
  {"author_id" => 3, "title" => "Green Eggs and Ham"},
  {"author_id" => 5, "title" => "The Adventures of Huckleberry Finn"}
]

# approach # 1
def connect_books_with_authors(books, authors)
  books_with_authors = []
  books.each do |book|
    authors.each do |author|
      if book["author_id"] == author["author_id"]
        books_with_authors << {
          title: book["title"],
          author: author["name"]
        }
      end
    end 
  end
  
  return books_with_authors 
end

# approach 2
def connect_books_with_authors(books, authors) 
  books_with_authors = []
  author_hash_table = {}
  
  # Convert author data into author hash table:
  authors.each do |author| 
    author_hash_table[author["author_id"]] = author["name"]
  end

  books.each do |book| 
    books_with_authors << {
      "title" => book["title"],
      "author" => author_hash_table[book["author_id"]]
    }
  end
  return books_with_authors
end

basketball_players = [
	{first_name: "Jill", last_name: "Huang", team: "Gators"},
	{first_name: "Janko", last_name: "Barton", team: "Sharks"},
	{first_name: "Wanda", last_name: "Vakulskas", team: "Sharks"},
	{first_name: "Jill", last_name: "Moloney", team: "Gators"},
	{first_name: "Luuk", last_name: "Watkins", team: "Gators"}
]

football_players = [
	{first_name: "Hanzla", last_name: "Radosti", team: "32ers"},
	{first_name: "Tina", last_name: "Watkins", team: "Barleycorns"},
	{first_name: "Alex", last_name: "Patel", team: "32ers"},
	{first_name: "Jill", last_name: "Huang", team: "Barleycorns"},
	{first_name: "Wanda", last_name: "Vakulskas", team: "Barleycorns"}
]

# Time complexity O(N + M)
# Space complexity O(N) for hash + O(N) for result?
def two_sports_players(basketball_players, football_players)
  players = {}
  result = []

  # add into the hash the full player names from basketball players
  # iterate through football players and if the player exists in the players hash, add the player full name to the result array 
  basketball_players.each do |player|
    full_name = player[:first_name] + " " + player[:last_name]
    players[full_name] = true
  end

  football_players.each do |player|
    full_name = player[:first_name] + " " + player[:last_name]
    if players[full_name]
      result << full_name
    end
  end

  return result
end

def greatest_product(array)
  greatest_num = array[0]
  lowest_num = array[0]
  greatest_product_sofar = 0

  array.each do |elem|
    p elem
    # p lowest_num
    # p greatest_num
    # check if the current number is higher than the greatest num
    if elem > greatest_num
      greatest_num = elem
      next
    end

    # check if the current number is lowest than the lowest num
    if elem < lowest_num 
      lowest_num = elem
      next
    end

    # check if the current product is greatest 
    if greatest_num * elem > greatest_product_sofar 
      greatest_product_sofar = greatest_num * elem
    end
    
    # check if the current product is greatest
    if lowest_num * elem > greatest_product_sofar
      greatest_product_sofar = lowest_num * elem
    end
  end
  return greatest_product_sofar
end

# [98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1]
def sort_body_temperature(array)
  hash = Hash.new(0)

  array.each do |elem|
    hash[elem] += 1
  end

  sorted_array = []
  
  # 97
  for i in 0..9
    current = "97.#{i}".to_f 
    if hash[current]
      hash[current].times do 
        sorted_array << current
      end
    end
  end

  for i in 0..9
    current = "98.#{i}".to_f 
    if hash[current]
      hash[current].times do 
        sorted_array << current
      end
    end
  end

  for i in 0..9
    current = "99.#{i}".to_f 
    if hash[current]
      hash[current].times do 
        sorted_array << current
      end
    end
  end
  sorted_array
end

p sort_body_temperature([98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1])

def longest_sequence(array)
  hash = {}
  count_hash = Hash.new(0)

  array.each do |elem|
    hash[elem] = true
  end

  # array_copy = array.dup 
  longest_sequence_so_far = 0
  temp_sequence = 0

  array.each do |elem|
    if hash[elem + 1] && hash[elem - 1]
      next
    elsif hash[elem + 1]
      next
    elsif hash[elem - 1]
      "only below"
      current = elem
      temp_sequence = 0
      while hash[current]
        temp_sequence += 1
        current -= 1
      end
    else
      next
    end

    if temp_sequence > longest_sequence_so_far
      longest_sequence_so_far = temp_sequence
    end
    temp_sequence = 0
  end
  return longest_sequence_so_far
end

# p two_sports_players(basketball_players, football_players)
# p greatest_product([5, 10, -6, 9, 4])
# p greatest_product([5, 5, 2, 3, 4, -4])

p longest_sequence([10, 5, 12, 3, 55, 30, 4, 11, 2])
p longest_sequence([19, 13, 15, 12, 18, 14, 17, 11])