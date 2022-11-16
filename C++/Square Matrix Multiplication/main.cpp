#include <vector>

using namespace std;

std::vector<std::vector<int>> matrix_multiplication(std::vector<std::vector<int>> &a, std::vector<std::vector<int>> &b, size_t n)
{
  // TODO: Return the result of A × B in the form of an n × n matrix
  // NOTE: Please allocate and return matrix C
  vector<vector<int>> result(n, vector<int>(n, 0));

  for (int r = 0; r < n; r++)
  {
    for (int c = 0; c < n; c++)
    {
      int total = 0;
      for (int i = 0; i < n; i++)
      {
        total += a[r][i] * b[i][c];
      }
      result[r][c] = total;
    }
  }

  return result;
}