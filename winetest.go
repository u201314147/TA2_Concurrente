
package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {


	rawData, err := base.ParseCSVToInstances("datasets/wine.csv", false)

	//dataset de api
	//rawData, err := base.ParseCSVToInstances("datasets/wineapi.csv", false)


	if err != nil {
		panic(err)
	}

	// Imprimimos la lista de data del csv y el endpoint
	fmt.Println(rawData)

	//Inicializamos el clasificador KNN
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	//Entrenamos la data
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	//calculamos la distancia Euclidiana y retorna la mas cercana
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	// Imprimimos los resultados en la matriz de confusion
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("No se pudo tenre la matriz de confusion: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}
