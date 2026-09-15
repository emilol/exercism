package lasagnamaster

func PreparationTime(layers []string, averageTime int) int {
	if averageTime == 0 {
		averageTime = 2
	}
	return len(layers) * averageTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	scale := float64(portions) / 2
	scaledQuantities = []float64{}

	for _, quantity := range quantities {
		scaledQuantity := quantity * scale
		scaledQuantities = append(scaledQuantities, scaledQuantity)
	}
	return scaledQuantities
}
