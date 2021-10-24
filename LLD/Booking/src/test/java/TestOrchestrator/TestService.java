package TestOrchestrator;


import Models.Booking;
import Models.BookingStatus;
import Models.Driver;
import Orchestrator.Service;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;


public class TestService {

    @Mock
    Service service = new Service();

    @Test
    public void TestRegisterDriver(){
        Driver driver = new Driver(1,"Driver1",0, BookingStatus.UNBOOKED,0);
        Driver testDriver = service.registerDriver(1);
        Assertions.assertEquals(driver.getId(), testDriver.getId());
        Assertions.assertEquals(driver.getBookingCount(), testDriver.getBookingCount());
        Assertions.assertEquals(driver.getBookingStatus(), testDriver.getBookingStatus());
        Assertions.assertEquals(driver.getDistCovered(), testDriver.getDistCovered());
    }

    @Test
    public void TestDispatchDriverForBooking(){
        service.getDriverList().add(new Driver(1,"Driver1",0, BookingStatus.UNBOOKED,0));
        service.getDriverList().add(new Driver(2,"Driver2",0, BookingStatus.UNBOOKED,0));
        service.getDriverList().add(new Driver(3,"Driver3",0, BookingStatus.UNBOOKED,0));
        Assertions.assertEquals(1,service.dispatchDriverForBooking(10));
        Assertions.assertEquals(2,service.dispatchDriverForBooking(10));
        Assertions.assertEquals(3,service.dispatchDriverForBooking(10));
        Assertions.assertEquals(-1,service.dispatchDriverForBooking(10));
        TestCompleteBooking();
    }

    @Test
    @Disabled
    public void TestCompleteBooking(){
        Assertions.assertTrue(service.completeBooking(3));
        Assertions.assertTrue(service.completeBooking(2));
        Assertions.assertTrue(service.completeBooking(1));
        Assertions.assertFalse(service.completeBooking(3));
    }

}
