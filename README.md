# Cryptography Lab

This repository contains the implementation of various cryptographic algorithms and number theory concepts. 🛡️🔐

## Table of Contents

- [Introduction](#introduction)
- [Algorithms Implemented](#algorithms-implemented)
- [Getting Started](#getting-started)
- [Tasks](#tasks)
- [Key Features](#key-features)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This repository serves as a comprehensive resource for understanding and implementing fundamental cryptographic algorithms and number theory concepts. It provides clear and concise code examples for each algorithm, making it a valuable reference for anyone interested in cryptography and related mathematical topics.
The algorithms were implemented using math/big package for handling large integers.

## Algorithms Implemented

- **Euler's Totient Function (φ)**: Calculate Euler's totient function for a given integer.
- **Möbius Function (μ)**: Compute the Möbius function for a given integer.
- **Least Common Multiple (LCM)**: Find the least common multiple of a set of numbers.
- **Chinese Remainder Theorem (CRT)**: Solve systems of linear congruences using the Chinese Remainder Theorem.
- **Legendre and Jacobi Symbols**: Calculate Legendre and Jacobi symbols for given integers.
- **Pollard's Rho Algorithm**: Factorize large integers using Pollard's Rho algorithm.
- **Baby Step Giant Step Algorithm**: Solve the discrete logarithm problem using the Baby Step Giant Step algorithm.
- **Cipolla's Algorithm**: Find square roots modulo a prime.
- **Miller-Rabin Primality Test**: Check the primality of an integer using the Miller-Rabin primality test.
- **RSA Cryptosystem**: Implement the RSA encryption and decryption scheme.
- **ElGamal Cryptosystem**: Demonstrate the ElGamal cryptosystem over elliptic curves.

## Getting Started
### Prerequisites

- Go (Golang) should be installed on your machine. If not, you can download it from [the official Go website](https://golang.org/dl/).

### Clone the Repository

You can clone this repository to your local machine using the following command:

```bash
git clone https://github.com/YuriiCodes/go-cryptography-lab-1.git
```
### Install Go Dependencies

Navigate to the project directory and install the required Go dependencies. You can do this using the following command:

```bash
go get -u ./...

This command will download and install all the necessary dependencies for the project.
```
### Build the Project

After installing the dependencies, you can build the Cryptography Lab 1 project using the following command:

```bash
go build crypto-lab-1
```

This will compile the project, and you will find the executable file named `crypto-lab-1` in the project directory.

### Run the Project

To run the Cryptography Lab 1 application, use the following command:

```bash
./crypto-lab-1
```


## Tasks

This repository is organized into tasks corresponding to the implementation of different algorithms and concepts. Each task is a self-contained module that can be independently executed. Here's a quick overview of the tasks:

1. **Task 1**: Euler's Totient Function (φ), Möbius Function (μ), and Least Common Multiple (LCM).
2. **Task 2**: Chinese Remainder Theorem (CRT) for solving linear congruences.
3. **Task 3**: Legendre and Jacobi Symbols for quadratic residues.
4. **Task 4**: Pollard's Rho Algorithm for integer factorization.
5. **Task 5**: Baby Step Giant Step Algorithm for solving discrete logarithms.
6. **Task 6**: Cipolla's Algorithm for finding square roots modulo a prime.
7. **Task 7**: Miller-Rabin Primality Test for checking primality.
8. **Task 8**: Implementation of the RSA Cryptosystem.
9. **Task 9**: Demonstration of the ElGamal Cryptosystem over elliptic curves.

