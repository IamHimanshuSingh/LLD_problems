package models;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@AllArgsConstructor
@Setter
@Getter
public class ExpenseMetadata {
    private String name;
    private String imgUrl;
    private String notes;
}
