def merge(array, ):
    pass

def mergesort(array, initial_value, final_value):
    if initial_value < final_value:

        mid_value = (initial_value + final_value - 1)/2
        mergesort(array, initial_value, mid_value)
        mergesort(array, mid_value+1, final_value)
        merge(array)


input_array = [1, 4, 2, 5, 3]
print "Input array is {}".format(input_array)

mergesort(input_array, 0, len(input_array)-1)

