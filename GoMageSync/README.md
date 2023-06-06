# GoMageSync

GoMageSync is a project that provides APIs to interact with Magento community latest edition in Golang and syncs data with Deskera using Java cron jobs.

## Getting Started

To get started, clone the repository:

```
git clone https://github.com/juniorvish/GoMageSync.git
```

## Golang APIs

The Golang APIs include:

1. Get all products with filters and pagination
2. Get all customers with filters and pagination
3. Get all orders with filters and pagination
4. Get all payments against orders with filters and pagination
5. Get details on a product by id
6. Get details on a customer by id
7. Get details on an order by id

### Running the Golang APIs

To run the Golang APIs, navigate to the `GoMageSync` directory and execute:

```
go run golang_api.go
```

## Java Cron Jobs

The Java cron jobs sync data between Magento and Deskera:

1. Sync newly created products from Magento to Deskera Products every 30 minutes
2. Sync newly created customers from Magento to Deskera Contacts every 30 minutes
3. Sync newly created orders from Magento to Deskera Sales Invoice every 30 minutes

### Running the Java Cron Jobs

To run the Java cron jobs, navigate to the `GoMageSync/java_deskera_sync` directory and execute:

```
javac DeskeraSync.java
java DeskeraSync
```

## Tailwind CSS (if required)

If the project requires a web-app with CSS, Tailwind CSS is used for the UI. Ensure the UI is neat and clean.

## License

This project is licensed under the MIT License.