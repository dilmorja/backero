#include <iostream>
#include "../backero.hpp"

using namespace std;
using namespace back;

string test_uti() {
  UTI unique("test", "v0", "test");

  if (unique.toString() != "test_v0_test") {
    return "test_uti: toString() is not equal to \"test_v0_test\"";
  }

  return "test_uti: Done!";
}

int main() {

  cout << test_uti();

  return 0;
}