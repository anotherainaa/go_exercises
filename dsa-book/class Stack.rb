class Stack
  def initialize
    @data = []
  end
  
  def push(element)
    @data << element
  end
  
  def pop
    @data.pop
  end
  
  def read
    @data.last
  end
end

# Example use case - linter

class Linter
  def initialize
    @stack = Stack.new
  end

  def lint(text)
    text.each_char do |char|
      if is_opening_brace?(char) 
        @stack.push(char)
      elsif is_closing_brace?(char)
        popped_opening_brace = @stack.pop
        
        if !popped_opening_brace
          return "#{char} doesn't have opening brace"
        end
        
        if is_not_a_match(popped_opening_brace, char)
          return "#{char} has mismatched opening brace"
        end
      end
    end
    
    if @stack.read
      return "#{@stack.read} does not have closing brace"
    end
    
    return true
  end
  
  private
  
  def is_opening_brace?(char)
    ["(", "[", "{"].include?(char)
  end
  
  def is_closing_brace?(char)
    [")", "]", "}"].include?(char)
  end
  
  def is_not_a_match(opening_brace, closing_brace)
    closing_brace != {"(" => ")", "[" => "]", "{" => "}"}[opening_brace]
  end
end

# linter = Linter.new
# puts linter.lint("(var x = { y: [1, 2, 3]")


class Queue
  def initialize
    @data = []
  end
  
  def enqueue(element)
    @data << element
  end
  
  def dequeue
    @data.shift
  end
  
  def read
    @data.first
  end
end

# Example use case

class PrintManager
  def initialize
    @queue = Queue.new
  end
  
  def queue_print_job(document)
    @queue.enqueue(document)
  end
  
  def run
    while @queue.read
      print(@queue.dequeue)
    end
  end
  
  private
  
  def print(document)
    puts document
  end
end

# print_manager = PrintManager.new
# print_manager.queue_print_job("First doc")
# print_manager.queue_print_job("Second doc")
# print_manager.queue_print_job("Third doc")
# print_manager.run


def string_reverse(str)
  stack = Stack.new()

  str.chars do |char|
    stack.push(char) 
  end

  reversed = ''
  while stack.read 
    reversed << stack.pop
  end
  reversed
end

def reverse(str)
  return str[0] if str.length == 1

  return reverse(str[1, str.length - 1]) + str[0]
end

def number_of_paths(n)
  # original base case
  # return 0 if n <= 0
  # return 1 if n == 1
  # return 2 if n == 2
  # return 4 if n == 3

  return 0 if n < 0
  return 1 if n == 1 || n == 0

  return number_of_paths(n - 1) + number_of_paths(n - 2) + number_of_paths(n - 3)
end


def anagrams_of(string)
  # base case: if the string is only one character
  # return an array containing just a single-character string

  return [string[0]] if string.length == 1
  
  # create an array to hold all the anagrams
  collection = []

  # Find all anagrams of the substring from the second character until the end. 
  # For example, if the string is "abcd", the substring is "bcd"
  # so we'll find all the anagrams of "bcd"
  substring_anagrams = anagrams_of(string[1, string.length - 1])

  # Iterate over each substring
  substring_anagrams.each do |substring_anagram|
  
    # iterate over each index of the substring, from 0 until
    # one index past the end of the string
    (0..substring_anagram.length).each do |index|

      # Create a copy of the substring anagram
      copy = String.new(substring_anagram)
      # Insert the first character of our string into the substring anagram copy. 
      # Where it will go depends on the index we're up to within this loop. 
      # Then, take this new string and add it to our collection of anagrams:
      collection << copy.insert(index, string[0])
    end

    # Return the entire collection of anagrams
    return collection
  end
end

def total_string(array) 
  return array[0].length if array.length == 1

  return array[0].length + total_string(array[1, array.length - 1])
end

def even_numbers(array)
  # if array.length == 1 
  #   if array[0] % 2 == 0 
  #     return [array[0]]
  #   else
  #     return []
  #   end
  # end

  return [] if array.length == 0

  if array[0] % 2 == 0 
    return [array[0]] + even_numbers(array[1, array.length - 1])
  else
    return even_numbers(array[1, array.length - 1])
  end
end


def triangle_number(num)
  return 1 if num == 1 

  return triangle_number(num - 1) + num
end

def first_x(string)
  return 0 if string.length == 1

  if string[0] == 'x'
    return 0
  else
    return 1 + first_x(string[1, string.length - 1])
  end
end

def unique_paths(rows, columns)
  return 1 if rows == 1 || columns == 1

  return unique_paths(rows - 1, columns) + unique_paths(rows, columns - 1)
end

# p unique_paths(3, 7)

def sum(low, high)
  return low if high == low

  return high + sum(low, high - 1)
end

# p sum(1, 10)

def print_num(array)
  return if array.length == 0

  array.each do |elem|
    if elem.is_a? Array
      print_num(elem)
    else 
      puts elem
    end
  end
end

def fib(n, memo)
  if n == 0 || n == 1
    return n
  end

  if !memo[n]
    memo[n] = fib(n - 2, memo) + fib(n - 1, memo)
  end

  return memo[n]
end


p fib(10, {})

def fib2(n)
  if n == 0
    return 0
  end

  a = 0
  b = 1

  for i in 1...n
    temp = a
    a  = b
    b = temp + a
  end

  return b
end

p fib2(10)

# def add_until_100(array)
#   return 0 if array.length == 0

#   if array[0] + add_until_100(array[1, array.length - 1]) > 100
#     return add_until_100(array[1, array.length - 1])
#   else
#     return array[0] + add_until_100(array[1, array.length - 1])
#   end
# end
# print_num([1, 2, 3, [4, 6, 7], 5, [8, 9, [10, 11]]])

# p first_x("abx") 4
# p first_x("x") 
# p first_x("abcdx") 
# p first_x("abxd") 
# p first_x("abcxdef") 
# p first_x(("a".."z").to_a.join("")) 
# p triangle_number(8)
# p even_numbers([1, 2, 3, 4])
# puts total_string(["ab", "c", "def", "ghij"])
# puts total_string(["ab", "c", "def", "ghij", "abc"])
# puts reverse('abcde')
# puts number_of_paths(8)
# puts anagrams_of("abc")

class SortableArray
  attr_reader :array

  def initialize(array)
    @array = array
  end

  def partition!(left_pointer, right_pointer)
    # We always choose the right-most element as the pivot
    # We keep the index of the pivot for later use 
    pivot_index = right_pointer

    # We grab the pivot value itself:
    pivot = @array[pivot_index]

    # We start the right pointer immediately to the left 
    right_pointer -= 1

    while true
      # move the left pointer to the right as long as it 
      # points to a value that is greater than the pivot: 
      while @array[left_pointer] < pivot do
        left_pointer += 1
      end

      # Move to the right to the left as long as it
      # points to a value that is greater than the pivot
      while @array[right_pointer] > pivot do
        right_pointer -= 1
      end

      # We've now reached the point where we've stopped
      # moving both the left and right pointers
      
      # We check whether the left pointer has reached 
      # or gone beyond the right pointer. If it has, 
      # we break out of the loop so we can swap the pivot later on in our code

      if left_pointer >= right_pointer
        break

      # if the pointer is still to the left of the right pointer,
      # we swap the values of the left and right pointers:
      else 
        @array[left_pointer], @array[right_pointer] = @array[right_pointer], @array[left_pointer]
      end
    end

    # as the final step of the partition, we swap the value
    # of the left pointer with the pivot:
    @array[left_pointer], @array[pivot_index] = @array[pivot_index], @array[left_pointer]

    # we return the left_pointer for the sake of the quicksort method which will appear later in this chapter
    return left_pointer
  end

  def quicksort!(left_index, right_index)
    # Base case: the subarray has 0 or 1 elements:
    if right_index - left_index <= 0
      return
    end

    # Partition the range of elements and grab the index of the pivot:
    pivot_index = partition!(left_index, right_index)

    # Recursively call this quicksort! ethod on wathever is to the left of the pivot:
    quicksort!(left_index, pivot_index - 1)

    # Recursively call this quicksort! method on whatever is to the right of the pivot:
    quicksort!(pivot_index + 1, right_index)
  end

  def quickselect!(kth_lowest_value, left_index, right_index)
    # If we reach a base case - that is, that the subarray has one cell, 
    # we know we've found the value we're looking for:
    if right_index - left_index <= 0
      return @array[left_index]
    end

    # partition the array and grab the index of the pivot:
    pivot_index = partition!(left_index, right_index)

    # if what we're looking for is to the left of the pivot:
    if kth_lowest_value < pivot_index
      # recursively perform quickselect on the subaray to the left of the pivot
      quickselect!(kth_lowest_value, left_index, pivot_index - 1)
    
    # If what we're looking for is to the right of the pivot:
    elsif kth_lowest_value > pivot_index
      # Recursively perform quickselect on the subarray to the right of the pivot
      quickselect!(kth_lowest_value, pivot_index + 1, right_index)
    else # if kth lowest_value == pivot_index
      # if after the partition, the pivot position is in the same spot
      # as the kth lowest value, we've found the value we're looking for
      return @array[pivot_index]
    end
  end
end


array = [0, 5, 2, 1, 6, 3]

sortable_array = SortableArray.new(array)
sortable_array.quicksort!(0, array.length - 1)
p sortable_array.array

array2 = [0, 50, 20, 10, 60, 30]
sortable_array2 = SortableArray.new(array2)
p sortable_array2.quickselect!(1, 0, array.length - 1)
p sortable_array2.array


def has_duplicate_val(array)
  array.sort!

  array.each_with_index do |elem, index|
    if elem == array[index + 1]
      return true
    end
  end

  return false
end

def largest_sum_of_3(array)
  array.sort! do |a, b|
    b <=> a
  end

  return array[0] + array[1] + array[2]
end

def find_missing_num(array)
  array.sort!

  for i in 0..array.length
    current_elem = array[i]
    next_elem = array[i + 1]

    if next_elem - current_elem > 1 
      return current_elem + 1
    end
  end
  return nil
end

array3 = [0, 5, 2, 1, 6, 3, 5]
p has_duplicate_val(array3)

p largest_sum_of_3(array3)

array4 = [5, 2, 4, 1, 0]
array5 = [9, 3, 2, 5, 6, 7, 1, 0, 4]

p find_missing_num(array4)
p find_missing_num(array5)

# O(N2) implementation
def greatest_number1(array)
  # nested loops? 
  array.each do |elem1|
    elem_is_greatest_number = true

    array.each do |elem2|
      if elem2 > elem1
        elem_is_greatest_number = false
      end
    end

    if elem_is_greatest_number
      return elem1
    end
  end
end

# O(N log N) implementation
def greatest_number2(array)
  # sort the array, return the last or first number depending on how it was sort
  array.sort!

  return array[-1]
end

# O(N) implementation
def greatest_number3(array)
  # temp max number, iterate through the array once only
  max = array.shift

  array.each do |elem|
    if elem > max
      max = elem
    end
  end
  return max
end

p greatest_number1(array5)
p greatest_number2(array5)
p greatest_number3(array5)

class Node
  attr_accessor :data, :next_node

  def initialize(data)
    @data = data
  end
end

node_1 = Node.new("once")
node_2 = Node.new("upon")
node_3 = Node.new("a")
node_4 = Node.new("time")

node_1.next_node = node_2
node_2.next_node = node_3
node_3.next_node = node_4

class LinkedList
  attr_accessor :first_node

  def initialize(first_node)
    @first_node = first_node
  end

  def read(index)
    # We begin at the first node of the list:
    current_node = first_node
    current_index = 0

    while current_index < index
      # We keep following the links of each node until we get to the index we're looking for
      current_node = current_node.next_node
      current_index += 1

      # If we're past the end of the list, that means the value cannot be found in the list, so return nil
      return nil unless current_node
    end

    return current_node.data
  end

  def index_of(value)
    # We begin at the first node of the list:
    current_node = first_node
    current_index = 0

    begin
      # If we find the data we're looking for, we return it
      if current_node.data == value
        return current_index
      end

      # Otherwise, we move on to the next node 
      current_node = current_node.next_node
      current_index += 1
    end while current_node

    # If we get through the entire list without finding the data, we return nil
    return nil
  end

  def insert_at_index(index, value)
    # We create the new node with the provided value
    new_node = Node.new(value)

    # If we are inserting at the beginning of the list
    if index == 0
      # Have our new node link to what was the first node 
      new_node.next_node = first_node

      # Establish that our new node is now the list's first node
      self.first_node = new_node
      return
    end

    # If we are inserting anywhere other than the beginning
    current_node = first_node
    current_index = 0

    # First, we access the node immediately before where the new node will go
    while current_index < (index -  1) do
      current_node = current_node.next_node
      current_index += 1
    end

    # Have the new node link to the next node
    new_node.next_node = current_node.next_node

    # Mofidy the link of the previous node to point to our new node
    current_node.next_node = new_node
  end

  def delete_at_index(index)
    # If we are deleting the first node
    if index == 0
      # Simply set the first node to be what is currently the second node:
      self.first_node = first_node.next_node
      return
    end

    current_node = first_node
    current_index = 0

    # First, we find the node immediately before the one we
    # want to delete and call it current_node:
    while current_index < (index - 1) do
      current_node = current_node.next_node
      current_index += 1
    end

    # We find the node that comes after the one we're deleting: 
    node_after_deleted_node = current_node.next_node.next_node

    # We change the link of the current node to point to the node_after_deleted_node, leaving the node we want to delete out of the list:
    current_node.next_node = node_after_deleted_node
  end

  def print_all
    # starting at first node, print everything until we reach the end of the node

    current_node = first_node
    while current_node
      puts current_node.data
      current_node = current_node.next_node
    end
  end

  def last
    # starting from first list
    current_node = first_node

    return if !current_node

    while current_node
      current_node = current_node.next_node

      return current_node.data if !current_node.next_node
    end
  end

  def reverse_list
    # starting from the first node

    current_node = first_node
    previous_node = nil

    while current_node
      # save the next node for use later
      next_node = current_node.next_node

      current_node.next_node = previous_node

      previous_node = current_node
      current_node = next_node
    end

    # We break out of the loop after reaching the end
    self.first_node = previous_node
  end
end

class DoublyLinkedList
  attr_accessor :data, :next_node, :previous_node

  def initialize(first_node=nil, last_node=nil)
    @first_node = first_node
    @last_node = last_node
  end

  def insert_at_end(value)
    new_node = Node.new(value)

    # if there are no element yet in the linked list
    if !first_node
      @first_node = new_node
      @last_node = new_node
    else
      new_node.previous_node = @last_node
      @last_node.next_node = new_node
      @last_node = new_node
    end
  end

  def remove_from_front
    removed_node = @first_node
    @first_node = @first_node.next_node
    return removed_node
  end

  def print_all_reverse
    current_node = @last_node

    return if !current_node

    while current_node
      puts current_node.data
      current_node = current_node.previous_node
    end
  end
end

class QueueV2
  attr_accessor :queue

  def initialize
    @data = DoublyLinkedList.new
  end

  def enqueue(element)
    @data.insert_at_end(element)
  end

  def dequeue
    removed_node = @data.remove_from_front
    return removed_node.data
  end

  def read
    return nil unless @data.first_node
    return @data.first_node.data 
  end
end
 
list = LinkedList.new(node_1)
# p list.read(3)
# p list.index_of("time")
# p list.insert_at_index(3, "purple")
# p list.read(2)
# p list.read(3)
# p list.read(4)
# p list.delete_at_index(3)
# p list.read(3)
list.print_all

list.reverse_list

list.print_all

# list.print_all

# def delete_middle_note(node)
#   node.data = node.next_node.data
#   node.next_node = node.next_node.next_node
# end


