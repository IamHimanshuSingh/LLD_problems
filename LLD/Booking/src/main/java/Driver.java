import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class Driver {
    private int id;
    private String name;
    private double distCovered;
    BookingStatus bookingStatus;
    private int bookingCount;
}
