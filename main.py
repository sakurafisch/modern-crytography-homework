import matplotlib.pyplot as plt

def birthday(n: int = 365, possibility: float = 0.5) -> float:
    m = n;
    poss = 1.0;
    possibility = 1 - possibility
    for i in range(0,m):
        poss = poss * (n - i) / n
        if poss <= possibility:
            return i
    raise Exception("Exception while calling func birthday")

if __name__ == '__main__':
    x = []
    y = []
    for i in range(1, 99):
        result = birthday(2**32, possibility=i/100)
        x.append(result)
        y.append(i/100)
    plt.plot(x, y)
    plt.show()
    plt.plot(y, x)
    plt.show()