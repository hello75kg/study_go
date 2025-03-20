package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.17.0"
)

// **初始化 Jaeger 追踪**
func initTracer() func() {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("order-service"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	return func() { _ = tp.Shutdown(context.Background()) }
}

// **第 3 级**：检查用户信用
func checkUserCredit(ctx context.Context) {
	tracer := otel.Tracer("order-tracer")
	ctx, span := tracer.Start(ctx, "CheckUserCredit")
	defer span.End()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("用户信用检查完成")
}

// **第 3 级**：检查库存可用性
func verifyStockAvailability(ctx context.Context) {
	tracer := otel.Tracer("order-tracer")
	ctx, span := tracer.Start(ctx, "VerifyStockAvailability")
	defer span.End()

	time.Sleep(600 * time.Millisecond)
	fmt.Println("库存可用性检查完成")
}

// **第 2 级**：验证订单
func validateOrder(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	tracer := otel.Tracer("order-tracer")

	ctx, span := tracer.Start(ctx, "ValidateOrder")
	defer span.End()

	time.Sleep(1 * time.Second)
	fmt.Println("订单校验完成")

	// 并行执行 `CheckUserCredit` 和 `VerifyStockAvailability`
	var childWg sync.WaitGroup
	childWg.Add(2)

	go func() {
		defer childWg.Done()
		checkUserCredit(ctx)
	}()
	go func() {
		defer childWg.Done()
		verifyStockAvailability(ctx)
	}()

	childWg.Wait()
}

// **第 3 级**：更新库存数据库
func updateInventoryDatabase(ctx context.Context) {
	tracer := otel.Tracer("order-tracer")
	ctx, span := tracer.Start(ctx, "UpdateInventoryDatabase")
	defer span.End()

	time.Sleep(600 * time.Millisecond)
	fmt.Println("库存数据库更新完成")
}

// **第 3 级**：通知仓库发货
func notifyWarehouse(ctx context.Context) {
	tracer := otel.Tracer("order-tracer")
	ctx, span := tracer.Start(ctx, "NotifyWarehouse")
	defer span.End()

	time.Sleep(700 * time.Millisecond)
	fmt.Println("仓库通知完成")
}

// **第 2 级**：扣减库存
func deductInventory(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	tracer := otel.Tracer("order-tracer")

	ctx, span := tracer.Start(ctx, "DeductInventory")
	defer span.End()

	time.Sleep(2 * time.Second)
	fmt.Println("库存扣减完成")

	// 并行执行 `UpdateInventoryDatabase` 和 `NotifyWarehouse`
	var childWg sync.WaitGroup
	childWg.Add(2)

	go func() {
		defer childWg.Done()
		updateInventoryDatabase(ctx)
	}()
	go func() {
		defer childWg.Done()
		notifyWarehouse(ctx)
	}()

	childWg.Wait()
}

// **第 1 级**：处理订单
func processOrder(ctx context.Context) {
	tracer := otel.Tracer("order-tracer")

	ctx, span := tracer.Start(ctx, "ProcessOrder")
	defer span.End()

	// 并行执行 `ValidateOrder` 和 `DeductInventory`
	var wg sync.WaitGroup
	wg.Add(2)

	go validateOrder(ctx, &wg)
	go deductInventory(ctx, &wg)

	wg.Wait()

	fmt.Println("订单处理完成")
}

// jaeger 收到的链路：
// Trace
// └── ProcessOrder (span)
//    ├── ValidateOrder (span)
//    │   ├── CheckUserCredit (span)
//    │   └── VerifyStockAvailability (span)
//    ├── DeductInventory (span)
//    │   ├── UpdateInventoryDatabase (span)
//    │   └── NotifyWarehouse (span)

func main() {
	cleanup := initTracer()
	defer cleanup()

	// 创建 Context 并执行追踪
	ctx := context.Background()
	processOrder(ctx)

	time.Sleep(5 * time.Second) // 等待数据上报
}
