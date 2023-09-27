# SDP_Assignment_1
# Observer Pattern Implementation in Go

This is a simple implementation of the Observer design pattern in the Go programming language. The Observer pattern is used to define a one-to-many dependency between objects, so that when one object changes state, all its dependents are notified and updated automatically.

This implementation includes the following key components:

- `Observer` Interface: Defines the method(s) that concrete observers must implement to receive updates.
- `ConcreteObserver` Implementation: A concrete observer that implements the `Observer` interface. It receives and processes updates.
- `Subject` Interface: Defines the methods that concrete subjects (observable objects) must implement.
- `ConcreteSubject` Implementation: A concrete subject that maintains a list of observers and notifies them when its state changes.

# Payment Strategy Implementation in Go

This is a simple implementation of the Payment Strategy pattern in the Go programming language. The Payment Strategy pattern is used to define a family of algorithms, encapsulate each one, and make them interchangeable. In this example, we have implemented two payment strategies: credit card payment and PayPal payment, and a ShoppingCart context that uses these strategies for payment.

This implementation consists of the following key components:

- `PaymentStrategy` Interface: Defines the method for executing a payment strategy.
- `CreditCardPayment` Strategy: A concrete payment strategy for credit card payments.
- `PayPalPayment` Strategy: A concrete payment strategy for PayPal payments.
- `ShoppingCart` Context: The context that utilizes the payment strategy for processing payments.

## Prerequisites

To run this implementation, you need to have Go (Golang) installed on your system. You can download and install Go from the official [Go website](https://golang.org/dl/).

## Usage

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Qanysh/SDP_Assignment_1

## Prerequisites

To run this implementation, you need to have Go (Golang) installed on your system. You can download and install Go from the official [Go website](https://golang.org/dl/).
