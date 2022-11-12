#include <string>

std::string incrementString(const std::string &str)
{
  std::string result = str;
  if (result.length() == 0)
    return "1";
  if (!isdigit(result[result.length() - 1]))
  {
    return result + "1";
  }
  else
  {
    if (result[result.length() - 1] != '9')
    {
      result[result.length() - 1]++;
      return result;
    }
    else
    {
      int i = result.length() - 1;
      for (; i >= 0 && result[i] == '9'; i--)
      {
        result[i] = '0';
      }
      if (i == 0)
        return "1" + result;
      else if (isdigit(result[i]))
      {
        result[i]++;
        return result;
      }
      else
      {
        return result.substr(0, i + 1) + "1" + result.substr(i + 1);
      }
    }
  }
}