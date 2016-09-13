/* Implementation of syack in swift

`swift -sdk $(xcrun --show-sdk-path --sdk macosx) queue.swift`
*/


class queue{

    var list = [Int]()

    func is_empty() -> Bool{
        if list.count == 0{
            return true
        }
        else{
            return false
        }
    }

    func size() -> Int{
        return list.count
    }

    func enqueue(value: Int){
        list.append(value)
    }

    func dequeue() -> Int{
        return list.removeFirst()
    }

    func print_queue(){
        print(list)
    }
}

var s = queue()
print(s.is_empty())
print(s.size())
s.enqueue(2)
s.enqueue(4)
s.print_queue()
s.dequeue()
s.print_queue()
