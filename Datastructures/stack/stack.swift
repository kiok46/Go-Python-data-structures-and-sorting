/* Implementation of syack in swift

`swift -sdk $(xcrun --show-sdk-path --sdk macosx) stack.swift`
*/


class stack{

    var list = [Int]()

    func is_empty() -> Bool{
        if list.count == 0{
            return true
        }
        else{
            return false
        }
    }

    func peek() -> Int{
        if !is_empty(){
            return list[list.count-1]
        }
        else{
            return 0
        }
    }

    func size() -> Int{
        return list.count
    }

    func push(value: Int){
        list.append(value)
    }

    func print_stack(){
        print(list)
    }

    func pop() -> Int{
        return list.removeLast()
    }

}

var s = stack()
print(s.is_empty())
print(s.peek())
print(s.size())
s.push(2)
s.print_stack()
print(s.pop())
s.print_stack()
