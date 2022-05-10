package misc

func T[T any](b bool, x, y T) T {
    if b {
        return x
    }
    return y
}
