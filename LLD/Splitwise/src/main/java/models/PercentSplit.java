package models;

import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
public class PercentSplit extends Split{
    double percent;

    public PercentSplit(User user, double percent) {
        super(user);
        this.percent = percent;
    }
}
