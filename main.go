package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"scrape/ginger/controllers"
	"scrape/ginger/models"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
)

func initTracer() func(context.Context) error {

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithEndpoint(collectorURL),
			otlptracehttp.WithURLPath("/v1/traces"),
		),
	)

	if err != nil {
		log.Fatalf("Failed to create exporter: %v", err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Fatalf("Could not set resources: %v", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}

func init() {
	// Create the database instance
	models.ConnectDatabase()
}

func main() {
	// Initialize the tracing exporter
	cleanup := initTracer()
	defer cleanup(context.Background())

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))

	// CORS middleware to allow all origins
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":   "Get /random quote",
			"source": "https://github.com/h4ck4life/ginger",
			"author": "@h4ck4life",
		})
	})

	r.GET("/random", controllers.FindQuotes)

	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
