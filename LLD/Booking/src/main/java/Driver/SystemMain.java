package Driver;

import Orchestrator.Service;

import java.util.Scanner;

public class SystemMain {



    public static void main(String[] args) {
        Service service = new Service();
        final String registerCommand = "register_driver";
        final String dispatchDriverCommand = "dispatch_driver_for_a_booking";
        final String completeBookingCommand = "complete_booking";
        final String driversBookingsGTCommand = "drivers_completed_booking_gt";
        final String driversDistanceGTCommand = "drivers_completed_distance_gt";



        Scanner sc= new Scanner(System.in);
        while(true){

            System.out.print("Enter a command: ");
            String str= sc.nextLine();
            String strArray[] = str.split("\\s");
            for(String s: strArray){
                System.out.println(s + " ");
            }
            if(strArray.length < 2){
                System.out.println("Invalid input");
                System.exit(-1);
            }
            switch(strArray[0]){
                case registerCommand: service.registerDriver(Integer.parseInt(strArray[1]));
                    break;
                case dispatchDriverCommand:service.dispatchDriverForBooking(Integer.parseInt(strArray[1]));
                    break;
                case completeBookingCommand: service.completeBooking(Integer.parseInt(strArray[1]));
                    break;
                case driversBookingsGTCommand: service.driverBookingsGT(Integer.parseInt(strArray[1]));
                    break;
                case driversDistanceGTCommand: service.driverDistanceGT(Integer.parseInt(strArray[1]));
                    break;
                default:
                    System.out.println("Invalid input.");
                    System.exit(-1);
            }

        }

    }
}
