the app is: GoMageSync

the files we have decided to generate are: 
1. golang_api.go
2. java_api_cronjobs.java
3. README.md

Shared Dependencies:
1. Auth Token: authToken
2. Filter Parameters:
   - createdDate
   - updatedDate
   - productName
   - productCode
   - customerName
   - customerCode
   - orderId
3. API Endpoints:
   - getAllProducts
   - getAllCustomers
   - getAllOrders
   - getAllPayments
   - getProductDetails
   - getCustomerDetails
   - getOrderDetails
4. Deskera API Endpoints:
   - syncProducts
   - syncCustomers
   - syncOrders
5. Cron Job Intervals:
   - syncInterval (30 minutes)
6. Payload Mappings:
   - customerMapping
   - productMapping
   - orderMapping
7. Project Repository URL: https://github.com/juniorvish/GoMageSync
8. Tailwind CSS (if required)