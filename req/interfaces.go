package req

type IRequest interface {
}

type IError interface {
    Code() int
    Message() string
}
