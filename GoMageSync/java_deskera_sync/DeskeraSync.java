import java.util.Timer;
import java.util.TimerTask;

public class DeskeraSync {

    private static final long syncInterval = 1800000; // 30 minutes in milliseconds

    public static void main(String[] args) {
        Timer timer = new Timer();

        // Sync products
        timer.schedule(new ProductSync(), 0, syncInterval);

        // Sync customers
        timer.schedule(new CustomerSync(), 0, syncInterval);

        // Sync orders
        timer.schedule(new OrderSync(), 0, syncInterval);
    }
}