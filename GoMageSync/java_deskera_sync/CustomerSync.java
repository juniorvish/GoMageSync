import java.time.LocalDateTime;
import java.util.List;
import java.util.Timer;
import java.util.TimerTask;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class CustomerSync {

    private static final String MAGENTO_API_URL = "https://your-magento-url.com/rest/V1";
    private static final String DESKERA_API_URL = "https://api.deskera.com";
    private static final String authToken = "your-auth-token";
    private static final long syncInterval = 1800000; // 30 minutes in milliseconds

    public static void main(String[] args) {
        Timer timer = new Timer();
        timer.schedule(new SyncCustomersTask(), 0, syncInterval);
    }

    static class SyncCustomersTask extends TimerTask {
        @Override
        public void run() {
            try {
                OkHttpClient client = new OkHttpClient();
                Gson gson = new Gson();

                // Get customers from Magento
                Request magentoRequest = new Request.Builder()
                        .url(MAGENTO_API_URL + "/customers/search?searchCriteria[filter_groups][0][filters][0][field]=created_at&searchCriteria[filter_groups][0][filters][0][value]=" + LocalDateTime.now().minusMinutes(30).toString() + "&searchCriteria[filter_groups][0][filters][0][condition_type]=gt")
                        .addHeader("Authorization", "Bearer " + authToken)
                        .build();

                Response magentoResponse = client.newCall(magentoRequest).execute();
                List<Customer> magentoCustomers = gson.fromJson(magentoResponse.body().string(), new TypeToken<List<Customer>>() {}.getType());

                // Sync customers to Deskera
                for (Customer customer : magentoCustomers) {
                    Request deskeraRequest = new Request.Builder()
                            .url(DESKERA_API_URL + "/contacts")
                            .addHeader("Authorization", "Bearer " + authToken)
                            .addHeader("Content-Type", "application/json")
                            .post(gson.toJson(customerMapping(customer)))
                            .build();

                    client.newCall(deskeraRequest).execute();
                }

            } catch (Exception e) {
                e.printStackTrace();
            }
        }

        private CustomerMapping customerMapping(Customer customer) {
            CustomerMapping mapping = new CustomerMapping();
            mapping.setName(customer.getFirstName() + " " + customer.getLastName());
            mapping.setEmail(customer.getEmail());
            mapping.setPhone(customer.getPhone());
            mapping.setAddress(customer.getAddress());
            return mapping;
        }
    }

    static class Customer {
        private String firstName;
        private String lastName;
        private String email;
        private String phone;
        private String address;

        // Getters and setters
    }

    static class CustomerMapping {
        private String name;
        private String email;
        private String phone;
        private String address;

        // Getters and setters
    }
}