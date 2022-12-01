public class Solution{
  
  public static boolean validParentheses(String parens)
  {
    if (parens.length() == 0) return true;
    else if (parens.charAt(0) == ')') return false;
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