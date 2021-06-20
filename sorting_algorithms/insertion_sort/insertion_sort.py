def insertion_sort(A):
    for i in range(1,len(A)):
        hole = i
        value = A[i]

        while (hole > 0 and A[hole -1] > value):
            # shifting greater values to the right
            A[hole] = A[hole - 1]
            hole = hole - 1
        # putting the value in its sorted place
        A[hole] = value

t = [7,2,5,3,1]
insertion_sort(t)
print(t)