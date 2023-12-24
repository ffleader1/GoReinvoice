package customerr

import "errors"

var ErrInvalidCellToMerge = errors.New("invalid cell to merge")
var ErrInvalidTableSize = errors.New("invalid table size")
var ErrUnsupportedImageFormat = errors.New("unsupported image format")
var ErrInvalidScaleData = errors.New("invalid scale data")
var ErrInvalidPointsLength = errors.New("invalid points length")
var ErrMissMatchObjectAndType = errors.New("miss match object and type")
var ErrDuplicatedObjectID = errors.New("object with the same id existed")
var ErrInvalidObjectID = errors.New("object with the id not found")
