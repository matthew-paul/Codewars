import java.util.List;
import java.util.ArrayList;

public class Kata {
  
  public static List<Object> filterList(final List<Object> list) {
    ArrayList<Object> filteredList = new ArrayList<Object>();
    for (int i = 0; i < list.size(); i++) {
      if (list.get(i).getClass() != String.class) {
        filteredList.add(list.get(i));
      }
    }
    
    return filteredList;
  }
}
