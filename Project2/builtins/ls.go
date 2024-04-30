package builtins

import (
    "fmt"
    "io"
    "io/ioutil"
)

func ListDirectory(w io.Writer, args ...string) error {
    var path string
    switch len(args) {
    case 0: // no directory path provided, use current directory
        path = "."
    case 1: // use provided directory path
        path = args[0]
    default: // more than one argument provided
        return fmt.Errorf("%w: expected zero or one argument (directory)", ErrInvalidArgCount)
    }

    files, err := ioutil.ReadDir(path)
    if err != nil {
        return err
    }

    for _, file := range files {
        fmt.Fprintln(w, file.Name())
    }

    return nil
}