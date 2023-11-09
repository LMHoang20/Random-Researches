def fibonacciGenerator():
    a, b = 0, 1
    while True:
        a, b = b, a + b
        yield b


if __name__ == '__main__':
    n = 10
    for i in fibonacciGenerator():
        print(i)
        n -= 1
        if n < 0:
            break