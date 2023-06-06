import java.util.Timer;
import java.util.TimerTask;
import java.util.List;
import java.util.Map;
import java.util.HashMap;
import java.io.IOException;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.MediaType;
import okhttp3.RequestBody;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

public class ProductSync {

    private static final String MAGENTO_API_URL = "https://your-magento-url.com/rest/V1";
    private static final String DESKERA_API_URL = "https://api.deskera.com";
    private static final String authToken = "your-auth-token";
    private static final long syncInterval = 1800000; // 30 minutes in milliseconds

    public static void main(String[] args) {
        Timer timer = new Timer();
        timer.schedule(new SyncProductsTask(), 0, syncInterval);
    }

    static class SyncProductsTask extends TimerTask {
        OkHttpClient client = new OkHttpClient();
        Gson gson = new Gson();

        @Override
        public void run() {
            try {
                List<Map<String, Object>> products = fetchMagentoProducts();
                syncProductsToDeskera(products);
            } catch (IOException e) {
                e.printStackTrace();
            }
        }

        private List<Map<String, Object>> fetchMagentoProducts() throws IOException {
            Request request = new Request.Builder()
                    .url(MAGENTO_API_URL + "/products?searchCriteria[filter_groups][0][filters][0][field]=created_at&searchCriteria[filter_groups][0][filters][0][value]=" + getLastSyncTime() + "&searchCriteria[filter_groups][0][filters][0][condition_type]=gt")
                    .addHeader("Authorization", "Bearer " + authToken)
                    .build();

            try (Response response = client.newCall(request).execute()) {
                String responseBody = response.body().string();
                Map<String, Object> resultMap = gson.fromJson(responseBody, new TypeToken<Map<String, Object>>() {}.getType());
                return (List<Map<String, Object>>) resultMap.get("items");
            }
        }

        private void syncProductsToDeskera(List<Map<String, Object>> products) throws IOException {
            for (Map<String, Object> product : products) {
                Map<String, Object> deskeraProduct = productMapping(product);
                String json = gson.toJson(deskeraProduct);

                RequestBody requestBody = RequestBody.create(json, MediaType.parse("application/json; charset=utf-8"));
                Request request = new Request.Builder()
                        .url(DESKERA_API_URL + "/product")
                        .post(requestBody)
                        .addHeader("Authorization", "Bearer " + authToken)
                        .build();

                try (Response response = client.newCall(request).execute()) {
                    if (!response.isSuccessful()) {
                        System.out.println("Failed to sync product: " + product.get("sku"));
                    }
                }
            }
        }

        private Map<String, Object> productMapping(Map<String, Object> product) {
            Map<String, Object> deskeraProduct = new HashMap<>();
            deskeraProduct.put("code", product.get("sku"));
            deskeraProduct.put("name", product.get("name"));
            deskeraProduct.put("description", product.get("custom_attributes"));
            deskeraProduct.put("price", product.get("price"));
            return deskeraProduct;
        }

        private String getLastSyncTime() {
            LocalDateTime lastSyncTime = LocalDateTime.now().minusMinutes(30);
            DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
            return lastSyncTime.format(formatter);
        }
    }
}