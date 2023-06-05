the app is: GoMageSync

the files we have decided to generate are: 
1. go_files: 
   - main.go
   - magento_api.go
   - deskera_api.go
   - sync.go
   - models.go
   - utils.go
2. java_files:
   - Main.java
   - MagentoSync.java
   - DeskeraSync.java
   - Scheduler.java
   - Mapper.java
3. other_files:
   - README.md

Shared dependencies:

1. Exported variables:
   - authToken

2. Data schemas:
   - Product
   - Customer
   - Order
   - Payment
   - DeskeraProduct
   - DeskeraContact
   - DeskeraSalesInvoice

3. Function names:
   - getAllProducts
   - getAllCustomers
   - getAllOrders
   - getAllPayments
   - getProductDetails
   - getCustomerDetails
   - getOrderDetails
   - syncMagentoToDeskeraProducts
   - syncMagentoToDeskeraContacts
   - syncMagentoToDeskeraSalesInvoices
   - mapCustomerPayload
   - mapProductPayload
   - mapOrderPayload

4. Message names:
   - productSyncMessage
   - customerSyncMessage
   - orderSyncMessage

5. ID names of DOM elements (if applicable):
   - N/A (no frontend specified)