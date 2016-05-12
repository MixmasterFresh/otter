package rope

//import{}

type Branch struct{
  parent *Branch
  left *Branch
  right *Branch
  contents string //Speed up may be possible by following the 4 byte multiplicity rule and replacing string with array
  split uint //refers to the dividing index between the left and right node relative to the subtree where this branch is the root 
  length uint
}

type Rope struct{
  head *Branch
  length uint
}

const{
  MAX_CONTENT_LENGTH = 8
}

//Rope functions

func (target *Rope)insert(index uint, content string){
  if(index <= target.length){
    target.head.insert(index, content)
    target.length += len(content)
  }
}

func (target *Rope)delete(index uint, length uint){
  if(index + length <= length){
    target.head.delete(index, length)
    target.length -= length
  }
}

func (target *Rope)to_string() string{
  return target.head.to_string()
}

//Branch Functions

func (b *Branch)has_left() bool{
  return !(left == nil)
}

func (b *Branch)has_right() bool{
  return !(right == nil)
}

func (b *Branch)is_empty() bool{
  return (left == nil && right == nil)
}

func (b *Branch)is_leaf() bool{
  return (!b.has_left() && !b.has_right())
}

func (b *Branch)is_left() bool{
  if(b.parent != nil){
    return b.parent.left == b
  }
  return false
}

func (b *Branch)is_root(){
  return b.parent == nil
}

func (b *Branch)to_string() string{
  if(b.has_left() && b.has_right()){
    return b.left.to_string() + b.right.to_string()
  }else if(b.has_left()) {
    return b.left.to_string()
  }else{
    return b.contents
  }
}

func (b *Branch)shatter_to(index uint) *Branch {
  if(b.length > MAX_CONTENT_LENGTH && b.is_leaf()){
    var left_length uint
    left_length = b.length / 2
    
    b.split_at(left_length)
    
    //Go deeper
    if(index >= left_length){
      b.right.shatter_to(index - left_length)
    } else{
      b.left.shatter_to(index)
    }
    
  } else if(!b.is_leaf()){
    if(index >= b.split){
      b.right.shatter_to(index - b.split)
    } else{
      b.left.shatter_to(index)
    }
  }
}

func (b *Branch)insert(index uint, content string, length uint){
  //Split up branch if necessary
  if(b.length > MAX_CONTENT_LENGTH && b.is_leaf()){
    b.shatter_to(index)
  }
  
  b.length += length
  
  //Perform the insert
  if(!b.is_leaf()){
    if(index > b.split){
      b.right.insert(index - b.split, content)
      
    }else if(index < b.split){
      b.split += length
      b.left.insert(index, content)
      
    }else if(b.left.length > b.right.length){
      b.push_out_right()
      new_node = Branch{parent: b.right, length: length, is_leaf: true, b.contents: content}
      b.right.left = new_node
      
      b.right.split += length
      b.right.length += length
      
    }else{
      b.push_out_left()
      new_node = Branch{parent: b.left, length: length, is_leaf: true, b.contents: content}
      b.left.right = new_node
      
      b.split += length
      b.left.length += length
    }
  }else{
    if(length + b.length > MAX_CONTENT_LENGTH){
      b.split_at(index)
      b.insert(index, content, length)
    }else{
      b.content = b.content[0:index] + content + b.content[index:length]
    }
  }
}

func (b *Branch)delete(index uint, length uint){//TODO: finish this
  if(length == 0){
    return
  }
  
  if b.is_leaf(){
    b.remove_content(index, length)
  }else{
    var left_length uint = b.split - index
    var right_length uint = length - left_length 
    
    if(left_length > 0){
      if(left_length == b.split){
        b.left.destroy()
      }else{
        b.left.delete(index,left_length)
      }
      b.length -= left_length
      b.split -= left_length
    }
    
    if(right_length > 0){
      if(right_length == b.length - b.split){
        b.right.destroy()
      }else{
        b.right.delete(0,right_length)
      }
      b.length -= right_length
    }
  }
  
  if !(b.has_left() && b.has_right()) {
    b.truncate()
  }
}

func (b *Branch)remove_content(index uint, length uint) {
  if(index == 0){
    b.contents = b.contents[length:b.length - 1]
  }else if(index + length >= b.length){
    b.contents = b.contents[0:index]
  }else{
    b.contents = b.contents[0:index] + b.contents[index+length:b.length-1]
  }
  b.length -= length
}

func (b *Branch)push_out_left(){
  new_left := Branch{parent: b, length: b.left.length, is_leaf: false, split: b.left.length, left: b.left}
  old_left := b.left
  old_left.parent = new_left
  b.left = new_left
}

func (b *Branch)push_out_right(){
  new_right := Branch{parent: b, length: b.right.length, is_leaf: false, split: 0, right: b.right}
  old_right := b.right
  old_right.parent = new_right
  b.right = new_right
}

func (b *Branch)split_at(index uint){
  //Create new branches
  new_left := Branch{parent: b, length: index, is_leaf: true}
  new_right := Branch{parent: b, length: b.length - index, is_leaf: true}
  
  new_left.contents = b.contents[0:index]
  new_right.contents = b.contents[index:b.length]
  
  //Clean up old branch
  b.contents = "" // may be alright to change to nil, but I think i'll play it safe for now.
  b.left = new_left
  b.right = new_right
  b.split = index
  b.is_leaf = false
}

func (b *Branch)truncate(){
  if(b.is_leaf()){
    b.destroy()
  }else{
    if( (b.has_left() && b.has_right()) || b.is_root() ){
      return
    }
    parent := b.parent
    var child *Branch
    if(b.has_left()){
      child = b.left
    }else{
      child = b.right
    }
    
    if(b.is_left()){
      parent.left = child
    }else{
      parent.right = child
    }
    
    child.parent = parent
    
    //Clean up afterwards
    b.left = nil
    b.right = nil
    b.parent = nil
  }
}

func (b *Branch)destroy(){
  if(!b.is_leaf()){
    if(b.has_left()){
      b.left.destroy()
      b.left = nil
    }
    if(b.has_right()){
      b.right.destroy()
      b.right = nil
    }    
  }
  b.parent = nil
}


