import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Service {
    private List<Driver> driverList;
    private List<Booking> bookingsList;
    private Map<Integer, Driver> bookingDriverMap;

    public Service(){
        this.driverList = new ArrayList<>();
        this.bookingsList = new ArrayList<>();
        this.bookingDriverMap =  new HashMap<>();
    }

    private static final int BOOKING_NOT_POSSIBLE = -1;

    private static int bookingIdRandom = 1;

    public int dispatchDriverForBooking(int bookingDistance){
        boolean bookingPossible = false;
        Driver driverRequired = null;
        for(Driver d : driverList){
            if(d.getBookingStatus() == BookingStatus.UNBOOKED){
                d.setBookingStatus(BookingStatus.BOOKED); // Driver booked
                driverRequired = d;
                bookingPossible = true;
                break;
            }
        }
        if(!bookingPossible){
            System.out.println("Sorry, Driver is not available!!!");
            return BOOKING_NOT_POSSIBLE;
        }
        Booking booking = new Booking();
        booking.setId(bookingIdRandom++);
        booking.setDistance(bookingDistance);
        bookingsList.add(booking);
        bookingDriverMap.put(booking.getId(), driverRequired);
        System.out.println("Driver " + driverRequired.getId() + " Assigned to Booking " + booking.getId());
        return booking.getId();
    }

    public void registerDriver(int driverId) {

        for(Driver d: driverList){
            if(d.getId() == driverId){
                System.out.println("INVALID REGISTRATION. DRIVER ALREADY EXISTS!!!");
                return;
            }
        }

        Driver driver = new Driver();
        driver.setId(driverId);
        driverList.add(driver);
        driver.setBookingStatus(BookingStatus.UNBOOKED);
        driver.setBookingCount(0);
        driver.setDistCovered(0);
        System.out.println("Driver driver-" + driver.getId() + " Registered");
    }

    public void completeBooking(int bookingId){
        if(!bookingDriverMap.containsKey(bookingId)){
            System.out.println("BOOKING ID DOES NOT EXIST");
            return;
        }
        Driver driverWhoCompletedBooking = bookingDriverMap.get(bookingId);
        driverWhoCompletedBooking.setBookingCount(driverWhoCompletedBooking.getBookingCount()+1);
        int distanceCoveredInCurrentBooking = 0;
        for(Booking b: bookingsList){
            if(b.getId() == bookingId){
                distanceCoveredInCurrentBooking = b.getDistance();
                break;
            }
        }
        driverWhoCompletedBooking.setDistCovered(driverWhoCompletedBooking.getDistCovered()+distanceCoveredInCurrentBooking);
        bookingDriverMap.remove(bookingId);
        driverWhoCompletedBooking.setBookingStatus(BookingStatus.UNBOOKED);
        System.out.println("Driver driver-" + driverWhoCompletedBooking.getId() + " is released to the pool");
    }

    public void driverBookingsGT(int bookings){
        for(Driver d: driverList){
            if(d.getBookingCount() >= bookings){
                System.out.println("Driver driver-" + d.getId() + "has completed " + bookings + " bookings");
            }
        }
    }

    public void driverDistanceGT(int distance){
        for(Driver d: driverList){
            if(d.getDistCovered() >= distance){
                System.out.println("Driver driver-" + d.getId() + "has covered" + distance + " distance");
            }
        }
    }

}
