package models;


import lombok.AllArgsConstructor;
import lombok.Data;

@AllArgsConstructor
@Data
public class User {
    private String userId;
    private String name;
    private String email;
    private String mobileNumber;
}
