class TreeNode
  attr_accessor :value, :left_child, :right_child

  def initialize(value, left = nil, right = nil)
    @value = value
    @left_child = left
    @right_child = right
  end
end

node1 = TreeNode.new(25)
node2 = TreeNode.new(75)
root = TreeNode.new(50, node1, node2)


# Big notation of search in a binary search tree is O(log N)
# A binary search tree will have Log(N) levels / rows
def search(search_value, node)
  if node == nil || node.value == search_value
    return node
  elsif search_value < node.value
    
    return search(search_value, node.left_child)
  else
    return search(search_value, node.right_child)
  end
end

# Big O Notation is search big O + 1
# Therefore the answer is O(log N) + 1
def insert(value, node)
  if value < node.value

    # if the left child does not exist, we want to insert the value as the left child
    if node.left_child == nil
      node.left_child = TreeNode.new(value)
    else
      insert(value, node.left_child)
    end

  elsif value > node.value
    # If the right child does not exist, we want to inset the value as the right chile
    if node.right_child == nil
      node.right_child = TreeNode.new(value)
    else
      insert(value, node.right_child)
    end
  end
end

def delete(value_to_delete, node)
  # The base case is when we've hit the bottom of the tree, and the parent node has no children
  if node == nil
    return nil
  
  # if the value we're deleting is less or greater than the current node, 
  # we set the left or right child respectively to be 
  # the return value of a recurisve call of this very method on the current
  # node's left or right subtree
  elsif value_to_delete < node.value
    node.left_child = delete(value_to_delete, node.left_child)

    # We return the current node (and its subtree if existent)
    # to be used as the new value of it's parent's left or right child 
    return node
  elsif value_to_delete > node.value
    node.right_child = delete(value_to_delete, node.right_child)
    return node

  # if the current node is the on we want to delete 
  elsif value_to_delete ==  node.value

    # if the current node has no left child, we delete it by returning its right child (and its subtree if existent)
    # to be its parent's new subtree
    if node.left_child == nil 
      return node.right_child

      # if the current node has no left or right child, this ends up being none as per the first line of code in this function
    elsif node.right_child == nil
      return node.left_child

      # if the current node has two children, we delete the current node
      # by calling the lift function 
      # which changes the current node's value to the value of its successor node 
    else
      # if the current node has two children, we delete the current node
      # by caling the lift function (below)
      # which changes the current node's
      # value to the value of its successor node 
      node.right_child = lift(node.right_child, node)
      return node
    end
  end
end

def lift(node, node_to_delete)
  # if the current node of this function has a left child
  # we recursively call this function to continue down 
  # the left subtree to find the successor node 
  if node.left_child
    node.left_child = lift(node.left_child, node_to_delete)
    return node

  # if the current node has no left child, that means the current node
  # of this function is the successor node, and we take it's value
  # and make it the new value of the node we're deleting
  else 
    node_to_delete.value = node.value
    return  node.right_child
  end
end

def traverse_and_print(node)
  return if node == nil

  traverse_and_print(node.left_child)
  puts node.value
  traverse_and_print(node.right_child)
end

def max_value(node)
  current_node = node
  if current_node.right_child
    next_node = current_node.right_child
  end

  loop do
    current_node = current_node.right_child
    next_node = current_node.right_child
    break if next_node == nil
  end

  return current_node.value
end

def max_recursion(node)
  if node.right_child == nil
    return node.value
  end

  return max_recursion(node.right_child)
end

# Pre-order traversal
def traverse_and_print2(node)
  return if node == nil
  p node.value

  traverse_and_print2(node.left_child)
  traverse_and_print2(node.right_child)
end

# post-order traversal
def traverse_and_print3(node)
  return if node == nil

  traverse_and_print3(node.left_child)
  traverse_and_print3(node.right_child)
  p node.value
end

insert(100, root)
insert(88, root)
insert(23, root)
insert(49, root)
insert(66, root)

# traverse_and_print(root)
p max_value(root)
p max_recursion(root)
# p search(23, root)