package service

import (
	"frete-rapido/src/domain/repository"
	"strconv"
	"strings"
)

// Check - validates the request
func Check(request repository.RequestQuote) (args []string) {
	args = make([]string, 0)

	// contains zipcode?
	if strings.EqualFold(request.Recipient.Address.Zipcode, "") {
		args = append(args, "Zipcode is required")
	}

	// contains volumes?
	if len(request.Volumes) == 0 {
		args = append(args, "Volumes is required")
	}

	// contains specific variables?
	for i, volume := range request.Volumes {
		if strings.EqualFold(volume.Category, "") {
			args = append(args, "Category is required for Volume "+strconv.Itoa(i))
		}
		if volume.Amount == 0 {
			args = append(args, "Amount is required for Volume "+strconv.Itoa(i))
		}
		if volume.UnitaryWeight == 0 {
			args = append(args, "UnitaryWeight is required for Volume "+strconv.Itoa(i))
		}
		if volume.Price == 0 {
			args = append(args, "Price is required for Volume "+strconv.Itoa(i))
		}
		if strings.EqualFold(volume.Sku, "") {
			args = append(args, "Sku is required for Volume "+strconv.Itoa(i))
		}
		if volume.Height == 0 {
			args = append(args, "Height is required for Volume "+strconv.Itoa(i))
		}
		if volume.Width == 0 {
			args = append(args, "Width is required for Volume "+strconv.Itoa(i))
		}
		if volume.Length == 0 {
			args = append(args, "Length is required for Volume "+strconv.Itoa(i))
		}
	}

	return args
}
