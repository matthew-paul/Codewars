#include <string>
#include <sstream>

std::string createPhoneNumber(const int arr [10]){
  std::ostringstream oss;
  oss << "(" << arr[0] << arr[1] << arr[2] << ") " << arr[3] << arr[4] << arr[5] << "-" << arr[6] << arr[7] << arr[8] << arr[9];
  return oss.str();
}