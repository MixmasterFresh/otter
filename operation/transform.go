package operation

import()

func (first Operation) transform(second *simpleOperation) []Operation {
    if second.format == INSERT {
        return first.transformWithInsert(second)
    }else if second.format == DELETE{
        return first.transformWithDelete(second)
    }else{
        return nil
    }
}

func (first Operation) transformWithInsert(second *simpleOperation) []Operation {
    if second.index < first.Index {
        copy := first
        copy.Index += second.length
        return []Operation{copy}
    }else if second.index > first.Index {
        if first.Format == DELETE {
            return first.transformDeleteWithInsert(second)
        }else{
            return []Operation{first}
        }
    }else{
        if first.Id < second.id {
            first.Index += second.length
            return []Operation{first}
        }else if  first.Id > second.id {
            return []Operation{first}
        }else if first.Author > second.author {
            copy := first
            copy.Index += second.length
            return []Operation{copy}
        }else{
           return []Operation{first}
        }
    }
}

func (first Operation) transformDeleteWithInsert(second *simpleOperation) []Operation {
    if second.index > first.Index && second.index - first.Index < first.Length {
        pivot := second.index - first.Index
        half1 := first
        half1.Length = pivot
        
        half2 := first
        half2.Index = second.index + second.length
        half2.Length = first.Length - 1
        return []Operation{half2,half1}
    }else{
        return []Operation{first}
    }
}

func (first Operation) transformWithDelete(second *simpleOperation) []Operation {
    if first.Format == DELETE {
        return first.transformDeleteWithDelete(second)
    }else{
        if second.index < first.Index {
            if first.Index - second.index <= second.length {
                first.Index = second.index
            }else{
                first.Index = first.Index - second.length
            }
            return []Operation{first}
        }else{
            return []Operation{first}
        }
    }
}

func (first Operation) transformDeleteWithDelete(second *simpleOperation) []Operation {
    if first.Index > second.index && first.Index < second.index + second.length {
        if first.Length > second.length - (first.Index - second.index) {
            first.Length -= second.length - (first.Index - second.index)
            first.Index = second.index 
            return []Operation{first}
        }else{
            return nil
        }
    }else if second.index > first.Index && second.index < first.Index + first.Length {
        if first.Index + first.Length > second.index + second.length {
            first.Length -= second.length
            return []Operation{first}
        }else{
            first.Length = second.index - first.Index
            return []Operation{first}
        }
    }else{
        return []Operation{first}
    }
}