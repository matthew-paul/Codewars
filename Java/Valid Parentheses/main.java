public class Solution{
  
  public static boolean validParentheses(String parens)
  {
    int openParens = 0;
    for (int i = 0; i < parens.length(); i++) {
      if (parens.charAt(i) == '(') openParens++;
      else if (parens.charAt(i) == ')') {
        if (openParens == 0) {
          return false;
        }
        else openParens--;
      }
    }
    return (openParens == 0);
  }
}