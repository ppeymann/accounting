package server

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPathSize(ctx *gin.Context) int {
	s, ok := ctx.Params.Get("size")
	if !ok {
		return 0
	}

	sn, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return sn
}

func GetPathOffset(ctx *gin.Context) int {
	s, ok := ctx.Params.Get("offset")
	if !ok {
		return 0
	}

	sn, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return sn
}

func GetPathID(ctx *gin.Context) (int64, error) {
	value, ok := ctx.Params.Get("id")
	if !ok {
		return 0, errors.New("no path param with :id name found")
	}

	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.New("failed to parse :id value to int64 type")
	}

	return id, nil
}

func GetPathUint64(ctx *gin.Context) (uint64, error) {
	value, ok := ctx.Params.Get("id")
	if !ok {
		return 0, errors.New("no path param with :id name found")
	}

	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, errors.New("failed to parse :id value to uint64 type")
	}

	return id, nil

}

func GetInt64Path(name string, ctx *gin.Context) (int64, error) {
	value, ok := ctx.Params.Get(name)
	if !ok {
		return 0, errors.New("no path param with " + name + " name found")
	}

	res, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.New("failed to parse " + name + " value to int64 type")
	}

	return res, nil
}

func GetStringPath(name string, ctx *gin.Context) (string, error) {
	value, ok := ctx.Params.Get(name)
	if !ok {
		return "", errors.New("no path param with " + name + " name found")
	}

	return value, nil
}
