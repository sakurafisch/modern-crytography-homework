def birthday(n=365):
    k = n;
    tmp=1.0;
    for i in range(0,k):
        tmp = tmp*(n-i)/n
        if tmp <= 0.5:
            print(tmp)
            print(i)
            break

if __name__ == '__main__':
    birthday(2**32)