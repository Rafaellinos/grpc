import grpc
import stock_market_pb2
import stock_market_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:8081')
    stub = stock_market_pb2_grpc.StockPriceStub(channel)
    response = stub.GetStockPrice(stock_market_pb2.StockRequest(symbol='GRPC'))
    print("Greeter client received:", str(response))

if __name__ == '__main__':
    run()
