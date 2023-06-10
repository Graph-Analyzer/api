package web

import (
	_ "graph-analyzer/api/docs"
	graphPropertyActions "graph-analyzer/api/graph-property/action"
	graphPropertyServices "graph-analyzer/api/graph-property/domain"
	infrastructureActions "graph-analyzer/api/infrastructure/action"
	"graph-analyzer/api/repository"
	uploadActions "graph-analyzer/api/upload/action"
	uploadServices "graph-analyzer/api/upload/domain"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Serve() {
	driver := InitDBConnection()
	defer func(driver neo4j.Driver) {
		err := driver.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(driver)

	// Repository
	graphRepository := repository.NewGraphRepository(driver)
	// Service
	averageClusteringCoefficientService := graphPropertyServices.NewAverageClusteringCoefficientService(graphRepository)
	connectedService := graphPropertyServices.NewConnectedService(graphRepository)
	degreeAssortativityCoefficientService := graphPropertyServices.NewDegreeAssortativityCoefficientService(graphRepository)
	degreeCorrelationService := graphPropertyServices.NewDegreeCorrelationService(graphRepository)
	degreeDistributionService := graphPropertyServices.NewDegreeDistributionService(graphRepository)
	densityService := graphPropertyServices.NewDensityService(graphRepository)
	diameterService := graphPropertyServices.NewDiameterService(graphRepository)
	fullGraphService := graphPropertyServices.NewFullGraphService(graphRepository)
	generalInformationService := graphPropertyServices.NewGeneralInformationService(graphRepository)
	hasCutEdgeService := graphPropertyServices.NewHasCutEdgeService(graphRepository)
	hasCutVertexService := graphPropertyServices.NewHasCutVertexService(graphRepository)
	queryService := graphPropertyServices.NewQueryService(graphRepository)
	robustnessService := graphPropertyServices.NewRobustnessService(graphRepository)
	uploadService := uploadServices.NewUploadService()
	uploadStatusService := uploadServices.NewUploadStatusService()
	// Action
	averageClusteringCoefficientAction := graphPropertyActions.NewAverageClusteringCoefficientAction(averageClusteringCoefficientService)
	connectedAction := graphPropertyActions.NewConnectedAction(connectedService)
	degreeAssortativityCoefficientAction := graphPropertyActions.NewDegreeAssortativityCoefficientAction(degreeAssortativityCoefficientService)
	degreeCorrelationAction := graphPropertyActions.NewDegreeCorrelationAction(degreeCorrelationService)
	degreeDistributionAction := graphPropertyActions.NewDegreeDistributionAction(degreeDistributionService)
	densityAction := graphPropertyActions.NewDensityAction(densityService)
	diameterAction := graphPropertyActions.NewDiameterAction(diameterService)
	fullGraphAction := graphPropertyActions.NewFullGraphAction(fullGraphService)
	generalInformationAction := graphPropertyActions.NewGeneralInformationAction(generalInformationService)
	hasCutEdgeAction := graphPropertyActions.NewHasCutEdgeAction(hasCutEdgeService)
	hasCutVertexAction := graphPropertyActions.NewHasCutVertexAction(hasCutVertexService)
	queryAction := graphPropertyActions.NewQueryAction(queryService)
	robustnessAction := graphPropertyActions.NewRobustnessAction(robustnessService)
	livenessAction := infrastructureActions.NewLivenessAction()
	readinessAction := infrastructureActions.NewReadinessAction()
	uploadAction := uploadActions.NewUploadAction(uploadService)
	uploadStatusAction := uploadActions.NewUploadStatusAction(uploadStatusService)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	publicEndpoints := router.Group("/api")
	{
		publicEndpoints.GET("/doc/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		graphPropertiesEndpoints := publicEndpoints.Group("/graph-property")
		{
			graphPropertiesEndpoints.GET("/average-clustering-coefficient", averageClusteringCoefficientAction.Invoke)
			graphPropertiesEndpoints.GET("/connected", connectedAction.Invoke)
			graphPropertiesEndpoints.GET("/degree-assortativity-coefficient", degreeAssortativityCoefficientAction.Invoke)
			graphPropertiesEndpoints.GET("/degree-correlation", degreeCorrelationAction.Invoke)
			graphPropertiesEndpoints.GET("/degree-distribution", degreeDistributionAction.Invoke)
			graphPropertiesEndpoints.GET("/density", densityAction.Invoke)
			graphPropertiesEndpoints.GET("/diameter", diameterAction.Invoke)
			graphPropertiesEndpoints.GET("/full-graph", fullGraphAction.Invoke)
			graphPropertiesEndpoints.GET("/general-information", generalInformationAction.Invoke)
			graphPropertiesEndpoints.GET("/has-cut-edge", hasCutEdgeAction.Invoke)
			graphPropertiesEndpoints.GET("/has-cut-vertex", hasCutVertexAction.Invoke)
			graphPropertiesEndpoints.GET("/robustness", robustnessAction.Invoke)
			graphPropertiesEndpoints.POST("/query", queryAction.Invoke)
		}
		publicEndpoints.POST("/upload", uploadAction.Invoke)
		publicEndpoints.GET("/upload-status", uploadStatusAction.Invoke)
	}

	router.GET("/live", livenessAction.Invoke)
	router.GET("/ready", readinessAction.Invoke)

	// Expose port 8080 for the api
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
